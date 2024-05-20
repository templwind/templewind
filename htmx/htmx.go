package htmx

import (
	"encoding/json"
	"net/http"
)

// IsHtmxRequest checks if the incoming request is an HTMX request
func IsHtmxRequest(r *http.Request) bool {
	return r.Header.Get("HX-Request") == "true"
}

// Redirect handles HTMX and non-HTMX redirects
func Redirect(w http.ResponseWriter, r *http.Request, path string) {
	if IsHtmxRequest(r) {
		w.Header().Set("HX-Redirect", path)
		w.WriteHeader(http.StatusNoContent) // 204
		return
	}

	http.Redirect(w, r, path, http.StatusFound) // 302
}

// LocationMap is used for JSON encoding the location and target
type LocationMap struct {
	Path   string `json:"path"`
	Target string `json:"target"`
}

// Location handles HTMX and non-HTMX location changes
func Location(w http.ResponseWriter, r *http.Request, target LocationMap) {
	if IsHtmxRequest(r) {
		data, err := json.Marshal(target)
		if err == nil {
			w.Header().Set("HX-Location", string(data))
			w.WriteHeader(http.StatusNoContent) // 204
			return
		}
	}

	http.Redirect(w, r, target.Path, http.StatusFound) // 302
}

// Trigger sets an HTMX trigger
func Trigger(w http.ResponseWriter, trigger string) {
	w.Header().Set("HX-Trigger", trigger)
}
