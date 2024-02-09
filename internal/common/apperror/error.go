package apperror

import "errors"

const HandlerErrFmt = "failed to establish handler: %v"

var (
	NoTemplateError = errors.New("missing template map")
	NilConfigError  = errors.New("nil config")
)
