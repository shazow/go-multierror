package multierror

import (
	"fmt"
	"strings"
)

// MultiError implements the Error interface. Use .Err() to check if it is nil
// or contains errors.
type MultiError []error

// Err returns MultiError as an error type, or nil if empty. Use this to check
// the state of MultiError and whether it contains errors.
func (e MultiError) Err() error {
	if len(e) == 0 {
		return nil
	}
	return e
}

// Append an error to MultiError, return an error representing the multiError state.
// nil errors are ignored.
func (e *MultiError) Append(err error) error {
	if err == nil {
		if len(*e) == 0 {
			// nil e.Append(nil) is still nil
			return nil
		}
		// non-nil e.Append(nil) is not nil
		return e
	}
	if e == nil {
		// nil e.Append(error) will instantiate itself
		*e = []error{err}
		return e
	}
	// Append to existing non-nil e
	*e = append(*e, err)
	return e
}

// Convert MultiError into an aggregated error string, implements the Error interface.
func (e MultiError) Error() string {
	if len(e) == 0 {
		// This behavior is different from normal nil errors which would panic
		// in this condition. We could replicate the original behavior by
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
