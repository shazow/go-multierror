package multierror_test

import (
	"fmt"

	multierror "github.com/shazow/go-multierror"
)

func oops(input string) error {
	return fmt.Errorf("failed to %s", input)
}

func okay(input string) error {
	return nil
}

func Example() {
	errors := multierror.New()

	// Let's say we're doing a multi-phase set of operations, where errors can
	// happen independently and we only care about the state between phases.

	errors.Add(okay("a")) // == nil
	errors.Add(okay("b")) // == nil
	errors.Add(okay("c")) // == nil

	// Handle errors?
	if errors != nil {
		fmt.Printf("errors in first phase: %q", errors.Error())
		// No output, errors == nil
	}

	errors.Add(oops("d")) // != nil
	errors.Add(oops("e")) // != nil
	errors.Add(okay("f")) // != nil, because an error has already occurred.

	// Handle errors?
	if errors != nil {
		fmt.Printf("errors in second phase: %q", errors.Error())
		// Output: errors in second phase: "2 errors: failed to d; failed to e"
	}
}
