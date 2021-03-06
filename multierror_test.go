package multierror

import (
	"errors"
	"fmt"
	"testing"
)

func ExampleNew() {
	err := MultiError{}
	// err.Err() == nil

	err.Append(nil)
	// err.Append(nil) == nil
	// err.Err() == nil

	anErr := errors.New("an error")
	err.Append(anErr)
	// err.Append(anErr) != nil
	// err.Err() != nil

	fmt.Println(err)
	// Output: an error
}

func ExampleNew_multiple() {
	err := MultiError{}
	// err.Err() == nil

	anErr := errors.New("an error")
	err.Append(anErr)
	err.Append(anErr)

	fmt.Println(err)
	// Output: 2 errors: an error; an error
}

func ExampleNew_convert() {
	// Convert existing error slice
	myErrors := []error{
		errors.New("a"),
		errors.New("b"),
		errors.New("c"),
	}

	err := MultiError(myErrors)

	fmt.Println(err)
	// Output: 3 errors: a; b; c
}

func ExampleNew_uncast() {
	// Convert a MultiError back to a normal error slice for processing.
	multiErr := MultiError{errors.New("a"), errors.New("b"), errors.New("c")}

	// We can convert it back to a slice of errors
	errs := []error(multiErr)
	for _, err := range errs {
		// Deal with errors like you normally would
		fmt.Println("err:", err)
	}
	// Output: err: a
	// err: b
	// err: c
}

func TestErrorCompat(t *testing.T) {
	err := MultiError{}

	noError := func() error {
		return nil
	}
	noErr := noError()

	// Both should be == nil
	if (err.Err() == nil) != (noErr == nil) {
		t.Errorf("New() is not the same as a new error: %q != %q", err, noErr)
	}

	// Single error should produce the same .Error()
	anErr := errors.New("some error")
	err.Append(anErr)
	if got, want := err.Error(), anErr.Error(); want != got {
		t.Errorf("got %q; want %q", got, want)
	}
}

// Test the primary intended scenario
func TestErrors(t *testing.T) {
	err := MultiError{}
	if err.Err() != nil {
		t.Error("new MultiError is not nil")
	}

	if e := err.Append(nil); e != nil {
		t.Error("err.Append(nil) returned non-nil")
	}
	if err.Err() != nil {
		t.Error("err.Append(nil) is not nil", err)
	}

	anErr := errors.New("an error")
	if e := err.Append(anErr); e == nil {
		t.Error("err.Append(Error) returned nil")
	}
	if err == nil {
		t.Error("err.Append(Error) is nil")
	}
	if got, want := err.Error(), "an error"; got != want {
		t.Errorf("got %q; want %q", got, want)
	}

	if e := err.Append(anErr); e == nil {
		t.Error("err.Append(Error) returned nil")
	}
	if got, want := err.Error(), "2 errors: an error; an error"; got != want {
		t.Errorf("got %q; want %q", got, want)
	}
}

// Test the alternate API, using slice lengths.
func TestAlternate(t *testing.T) {
	err := MultiError{}
	if err == nil {
		t.Error("manual MultiError is nil")
	}

	err.Append(errors.New("an error"))
	err.Append(errors.New("another error"))

	if len(err) != 2 {
		t.Error("err length is not 2")
	}

	if got, want := err.Error(), "2 errors: an error; another error"; got != want {
		t.Errorf("got %q; want %q", got, want)
	}
}

// Test a convenience function of converting many errors to MultiError.
func TestManyNew(t *testing.T) {
	err := MultiError{errors.New("an error"), errors.New("another error")}

	if err == nil {
		t.Error("new MultiError with starting errors is nil")
	}

	if got, want := err.Error(), "2 errors: an error; another error"; got != want {
		t.Errorf("got %q; want %q", got, want)
	}

	myErrors := []error{errors.New("a"), errors.New("b"), errors.New("c")}
	err = MultiError(myErrors)

	if got, want := err.Error(), "3 errors: a; b; c"; got != want {
		t.Errorf("got %q; want %q", got, want)
	}
}
