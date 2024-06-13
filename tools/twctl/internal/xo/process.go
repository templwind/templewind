package xo

type Process struct {
	Filename                string
	XoOnly                  bool
	IgnoredTypes            map[string]bool
	BaseImportPath          string
	InputPath               string
	OutputPath              string
	AdditionalIgnoreTypes   []string
	DuplicateFunctionMarker string
}

// ProcessOptFn is a function that modifies a Process struct
type ProcessOptFn func(*Process)

// New creates a new Process struct with the provided options
func New(options ...ProcessOptFn) *Process {
	// Create the options struct using the default options
	process := defaultProcessOpts()
	// Apply the options
	for _, opt := range options {
		opt(process)
	}

	return process
}

func defaultProcessOpts() *Process {
	return &Process{
		IgnoredTypes: map[string]bool{
			"ErrInsertFailed": true,
			"Error":           true,
			"ErrUpdateFailed": true,
			"ErrUpsertFailed": true,
			// Add other types to ignore as needed
		},
	}
}

func WithFilename(filename string) ProcessOptFn {
	return func(opts *Process) {
		opts.Filename = filename
	}
}

func WithInputPath(inputPath string) ProcessOptFn {
	return func(opts *Process) {
		opts.InputPath = inputPath
	}
}

func WithOutputPath(outputPath string) ProcessOptFn {
	return func(opts *Process) {
		opts.OutputPath = outputPath
	}
}

func WithIgnoredTypes(ignoredTypes map[string]bool) ProcessOptFn {
	return func(opts *Process) {
		opts.IgnoredTypes = ignoredTypes
	}
}

func WithBaseImportPath(baseImportPath string) ProcessOptFn {
	return func(opts *Process) {
		opts.BaseImportPath = baseImportPath
	}
}
