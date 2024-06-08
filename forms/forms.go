package forms

import (
	"encoding"
	"encoding/json"
	"encoding/xml"
	"errors"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

// Binder is the interface that wraps the Bind method.
type Binder interface {
	Bind(i interface{}, req *http.Request) error
}

// DefaultBinder is the default implementation of the Binder interface.
type DefaultBinder struct{}

// BindUnmarshaler is the interface used to wrap the UnmarshalParam method.
type BindUnmarshaler interface {
	UnmarshalParam(param string) error
}

// bindMultipleUnmarshaler is used by binder to unmarshal multiple values from request at once to
type bindMultipleUnmarshaler interface {
	UnmarshalParams(params []string) error
}

// BindPathParams binds path params to bindable object
func (b *DefaultBinder) BindPathParams(req *http.Request, i interface{}) error {
	// Path params are not typically available in a standard net/http.Request
	// Custom implementation needed if using a router that supports path params
	return nil
}

// BindQueryParams binds query params to bindable object
func (b *DefaultBinder) BindQueryParams(req *http.Request, i interface{}) error {
	if err := b.bindData(i, req.URL.Query(), "query"); err != nil {
		return &HTTPError{Code: http.StatusBadRequest, Message: err.Error(), Internal: err}
	}
	return nil
}

// BindBody binds request body contents to bindable object
func (b *DefaultBinder) BindBody(req *http.Request, i interface{}) (err error) {
	if req.ContentLength == 0 {
		return
	}

	ctype := req.Header.Get("Content-Type")
	switch {
	case strings.HasPrefix(ctype, "application/json"):
		if err = json.NewDecoder(req.Body).Decode(i); err != nil {
			return &HTTPError{Code: http.StatusBadRequest, Message: err.Error(), Internal: err}
		}
	case strings.HasPrefix(ctype, "application/xml"), strings.HasPrefix(ctype, "text/xml"):
		if err = xml.NewDecoder(req.Body).Decode(i); err != nil {
			return &HTTPError{Code: http.StatusBadRequest, Message: err.Error(), Internal: err}
		}
	case strings.HasPrefix(ctype, "application/x-www-form-urlencoded"), strings.HasPrefix(ctype, "multipart/form-data"):
		if err = req.ParseForm(); err != nil {
			return &HTTPError{Code: http.StatusBadRequest, Message: err.Error(), Internal: err}
		}
		if err = b.bindData(i, req.PostForm, "form"); err != nil {
			return &HTTPError{Code: http.StatusBadRequest, Message: err.Error(), Internal: err}
		}
	default:
		return ErrUnsupportedMediaType
	}
	return nil
}

// BindHeaders binds HTTP headers to a bindable object
func (b *DefaultBinder) BindHeaders(req *http.Request, i interface{}) error {
	if err := b.bindData(i, req.Header, "header"); err != nil {
		return &HTTPError{Code: http.StatusBadRequest, Message: err.Error(), Internal: err}
	}
	return nil
}

// Bind implements the `Binder#Bind` function.
func (b *DefaultBinder) Bind(i interface{}, req *http.Request) (err error) {
	if err := b.BindPathParams(req, i); err != nil {
		return err
	}
	method := req.Method
	if method == http.MethodGet || method == http.MethodDelete || method == http.MethodHead {
		if err = b.BindQueryParams(req, i); err != nil {
			return err
		}
	}
	return b.BindBody(req, i)
}

// bindData will bind data ONLY fields in destination struct that have EXPLICIT tag
func (b *DefaultBinder) bindData(destination interface{}, data map[string][]string, tag string) error {
	if destination == nil || len(data) == 0 {
		return nil
	}
	typ := reflect.TypeOf(destination).Elem()
	val := reflect.ValueOf(destination).Elem()

	if typ.Kind() == reflect.Map && typ.Key().Kind() == reflect.String {
		k := typ.Elem().Kind()
		isElemInterface := k == reflect.Interface
		isElemString := k == reflect.String
		isElemSliceOfStrings := k == reflect.Slice && typ.Elem().Elem().Kind() == reflect.String
		if !(isElemSliceOfStrings || isElemString || isElemInterface) {
			return nil
		}
		if val.IsNil() {
			val.Set(reflect.MakeMap(typ))
		}
		for k, v := range data {
			if isElemString {
				val.SetMapIndex(reflect.ValueOf(k), reflect.ValueOf(v[0]))
			} else {
				val.SetMapIndex(reflect.ValueOf(k), reflect.ValueOf(v))
			}
		}
		return nil
	}

	if typ.Kind() != reflect.Struct {
		if tag == "param" || tag == "query" || tag == "header" {
			return nil
		}
		return errors.New("binding element must be a struct")
	}

	for i := 0; i < typ.NumField(); i++ {
		typeField := typ.Field(i)
		structField := val.Field(i)
		if typeField.Anonymous {
			if structField.Kind() == reflect.Ptr {
				structField = structField.Elem()
			}
		}
		if !structField.CanSet() {
			continue
		}
		structFieldKind := structField.Kind()
		inputFieldName := typeField.Tag.Get(tag)
		if typeField.Anonymous && structFieldKind == reflect.Struct && inputFieldName != "" {
			return errors.New("query/param/form tags are not allowed with anonymous struct field")
		}

		if inputFieldName == "" {
			if _, ok := structField.Addr().Interface().(BindUnmarshaler); !ok && structFieldKind == reflect.Struct {
				if err := b.bindData(structField.Addr().Interface(), data, tag); err != nil {
					return err
				}
			}
			continue
		}

		inputValue, exists := data[inputFieldName]
		if !exists {
			for k, v := range data {
				if strings.EqualFold(k, inputFieldName) {
					inputValue = v
					exists = true
					break
				}
			}
		}

		if !exists {
			continue
		}

		if ok, err := unmarshalInputsToField(typeField.Type.Kind(), inputValue, structField); ok {
			if err != nil {
				return err
			}
			continue
		}

		if ok, err := unmarshalInputToField(typeField.Type.Kind(), inputValue[0], structField); ok {
			if err != nil {
				return err
			}
			continue
		}

		if structFieldKind == reflect.Pointer {
			structFieldKind = structField.Elem().Kind()
			structField = structField.Elem()
		}

		if structFieldKind == reflect.Slice {
			sliceOf := structField.Type().Elem().Kind()
			numElems := len(inputValue)
			slice := reflect.MakeSlice(structField.Type(), numElems, numElems)
			for j := 0; j < numElems; j++ {
				if err := setWithProperType(sliceOf, inputValue[j], slice.Index(j)); err != nil {
					return err
				}
			}
			structField.Set(slice)
			continue
		}

		if err := setWithProperType(structFieldKind, inputValue[0], structField); err != nil {
			return err
		}
	}
	return nil
}

func setWithProperType(valueKind reflect.Kind, val string, structField reflect.Value) error {
	if ok, err := unmarshalInputToField(valueKind, val, structField); ok {
		return err
	}

	switch valueKind {
	case reflect.Ptr:
		return setWithProperType(structField.Elem().Kind(), val, structField.Elem())
	case reflect.Int:
		return setIntField(val, 0, structField)
	case reflect.Int8:
		return setIntField(val, 8, structField)
	case reflect.Int16:
		return setIntField(val, 16, structField)
	case reflect.Int32:
		return setIntField(val, 32, structField)
	case reflect.Int64:
		return setIntField(val, 64, structField)
	case reflect.Uint:
		return setUintField(val, 0, structField)
	case reflect.Uint8:
		return setUintField(val, 8, structField)
	case reflect.Uint16:
		return setUintField(val, 16, structField)
	case reflect.Uint32:
		return setUintField(val, 32, structField)
	case reflect.Uint64:
		return setUintField(val, 64, structField)
	case reflect.Bool:
		return setBoolField(val, structField)
	case reflect.Float32:
		return setFloatField(val, 32, structField)
	case reflect.Float64:
		return setFloatField(val, 64, structField)
	case reflect.String:
		structField.SetString(val)
	default:
		return errors.New("unknown type")
	}
	return nil
}

func unmarshalInputsToField(valueKind reflect.Kind, values []string, field reflect.Value) (bool, error) {
	if valueKind == reflect.Ptr {
		if field.IsNil() {
			field.Set(reflect.New(field.Type().Elem()))
		}
		field = field.Elem()
	}

	fieldIValue := field.Addr().Interface()
	unmarshaler, ok := fieldIValue.(bindMultipleUnmarshaler)
	if !ok {
		return false, nil
	}
	return true, unmarshaler.UnmarshalParams(values)
}

func unmarshalInputToField(valueKind reflect.Kind, val string, field reflect.Value) (bool, error) {
	if valueKind == reflect.Ptr {
		if field.IsNil() {
			field.Set(reflect.New(field.Type().Elem()))
		}
		field = field.Elem()
	}

	fieldIValue := field.Addr().Interface()
	switch unmarshaler := fieldIValue.(type) {
	case BindUnmarshaler:
		return true, unmarshaler.UnmarshalParam(val)
	case encoding.TextUnmarshaler:
		return true, unmarshaler.UnmarshalText([]byte(val))
	}

	return false, nil
}

func setIntField(value string, bitSize int, field reflect.Value) error {
	if value == "" {
		value = "0"
	}
	intVal, err := strconv.ParseInt(value, 10, bitSize)
	if err == nil {
		field.SetInt(intVal)
	}
	return err
}

func setUintField(value string, bitSize int, field reflect.Value) error {
	if value == "" {
		value = "0"
	}
	uintVal, err := strconv.ParseUint(value, 10, bitSize)
	if err == nil {
		field.SetUint(uintVal)
	}
	return err
}

func setBoolField(value string, field reflect.Value) error {
	if value == "" {
		value = "false"
	}
	boolVal, err := strconv.ParseBool(value)
	if err == nil {
		field.SetBool(boolVal)
	}
	return err
}

func setFloatField(value string, bitSize int, field reflect.Value) error {
	if value == "" {
		value = "0.0"
	}
	floatVal, err := strconv.ParseFloat(value, bitSize)
	if err == nil {
		field.SetFloat(floatVal)
	}
	return err
}
