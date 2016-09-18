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

	// First phase
	errors.Add(okay("a")) // == nil
	errors.Add(okay("b")) // == nil
	errors.Add(okay("c")) // == nil

	// Handle possible errors
	if errors != nil {
		fmt.Println("errors in first phase:", errors)
		// No output, because errors == nil
	}

	// Second phase
	errors.Add(oops("d")) // != nil
	errors.Add(oops("e")) // != nil
	errors.Add(okay("f")) // != nil, because an error has already occurred.

	// Handle pssible errors
	if errors != nil {
		// err.Error() == "2 errors: failed to d; failed to e"

		fmt.Println("errors in second phase:", errors)
		// Output: errors in second phase: 2 errors: failed to d; failed to e
	}
}
