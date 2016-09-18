package multierror

import (
	"fmt"
	"strings"
)

// New nil MultiErrors can be created with New() == nil, or it can be
// pre-populated with existing errors like New(err1, err2) != nil.
// To convert an existing slice of errors, use New(myErrors...)
func New(errors ...error) multiError {
	if len(errors) != 0 {
		return multiError(errors)
	}
	var err multiError = nil
	return err
}

// multiError implements the Error interface, can be checked as nil just like normal errors.
type multiError []error

// Add an error to multiError, return itself. nils are ignored.
func (e *multiError) Add(err error) error {
	if err == nil {
		if len(*e) == 0 {
			// nil e.Add(nil) is still nil
			return nil
		}
		// non-nil e.Add(nil) is not nil
		return e
	}
	if e == nil {
		// nil e.Add(error) will instantiate itself
		*e = []error{err}
		return e
	}
	// Append to existing non-nil e
	*e = append(*e, err)
	return e
}

// Convert multiError into an aggregated error string, implements the Error interface.
func (e multiError) Error() string {
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
