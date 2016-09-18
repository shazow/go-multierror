package multierror

import (
	"fmt"
	"strings"
)

// New nil MultiErrors can be created with New() == nil, or it can be
// pre-populated with existing errors like New(err1, err2) != nil.
func New(errors ...error) MultiError {
	if len(errors) != 0 {
		return MultiError(errors)
	}
	var err MultiError = nil
	return err
}

// MultiError implements the Error interface, can be checked as nil just like normal errors.
// It can be instantiated as an Error-compliant nil using New, or a slice of
// errors can be converted directly using err := MultiError(errors)
type MultiError []error

// Add an error to MultiError, return itself. nils are ignored.
func (e *MultiError) Add(err error) error {
	if err == nil {
		if len(*e) == 0 {
			// nil MultiError.Add(nil) is still nil
			return nil
		}
		// non-nil MultiError.Add(nil) is not nil
		return e
	}
	if e == nil {
		// nil MultiError Add will instantiate itself
		*e = []error{err}
		return e
	}
	// Append to existing non-nil MultiError
	*e = append(*e, err)
	return e
}

// Convert MultiError into an aggregated error string, implements the Error interface.
func (e MultiError) Error() string {
	if len(e) == 0 {
		// This behavior is different from normal nil errors which would panic
		// in this condition. We can replicate the original behavior by
		// removing this if-statement block.
		return ""
	}
	if len(e) == 1 {
		// err.Error() == New(err).Error()
		return e[0].Error()
	}
	s := make([]string, 0, len(e))
	for _, err := range e {
		s = append(s, err.Error())
	}
	return fmt.Sprintf("%d errors: ", len(e)) + strings.Join(s, "; ")
}
