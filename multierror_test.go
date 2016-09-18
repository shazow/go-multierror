package multierror

import (
	"errors"
	"testing"
)

func ExampleMultiError() {
	err := New()
	// err == nil

	err.Add(nil)
	// err.Add(nil) == nil
	// err == nil

	anErr := errors.New("an error")
	err.Add(anErr)
	// err.Add(anErr) != nil
	// err != nil
	// err.Error() == "an error"

	err.Add(anErr)
	// err.Error() == "2 errors: an error; an error"
}

func TestErrorCompat(t *testing.T) {
	err := New()

	noError := func() error {
		return nil
	}
	noErr := noError()

	// Both should be == nil
	if (err == nil) != (noErr == nil) {
		t.Errorf("New() is not the same as a new error: %q != %q", err, noErr)
	}

	// Single error should produce the same .Error()
	anErr := errors.New("some error")
	err.Add(anErr)
	if got, want := err.Error(), anErr.Error(); want != got {
		t.Errorf("got %q; want %q", got, want)
	}
}

// Test the primary intended, using New and Add.
func TestErrors(t *testing.T) {
	err := New()
	if err != nil {
		t.Error("new MultiError is not nil")
	}

	if e := err.Add(nil); e != nil {
		t.Error("err.Add(nil) returned non-nil")
	}
	if err != nil {
		t.Error("err.Add(nil) is not nil")
	}

	anErr := errors.New("an error")
	if e := err.Add(anErr); e == nil {
		t.Error("err.Add(Error) returned nil")
	}
	if err == nil {
		t.Error("err.Add(Error) is nil")
	}
	if got, want := err.Error(), "an error"; got != want {
		t.Errorf("got %q; want %q", got, want)
	}

	if e := err.Add(anErr); e == nil {
		t.Error("err.Add(Error) returned nil")
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

	err.Add(errors.New("an error"))
	err.Add(errors.New("another error"))

	if len(err) != 2 {
		t.Error("err length is not 2")
	}

	if got, want := err.Error(), "2 errors: an error; another error"; got != want {
		t.Errorf("got %q; want %q", got, want)
	}
}

// Test a convenience function of converting many errors to MultiError.
func TestManyNew(t *testing.T) {
	err := New(errors.New("an error"), errors.New("another error"))

	if err == nil {
		t.Error("new MultiError with starting errors is nil")
	}

	if got, want := err.Error(), "2 errors: an error; another error"; got != want {
		t.Errorf("got %q; want %q", got, want)
	}
}
