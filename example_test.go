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
	errors := multierror.MultiError{}

	// Let's say we're doing a multi-phase set of operations, where errors can
	// happen independently and we only care about the state between phases.

	// First phase
	errors.Append(okay("a")) // == nil
	errors.Append(okay("b")) // == nil
	errors.Append(okay("c")) // == nil

	// Handle possible errors
	if errors.Err() != nil {
		fmt.Println("errors in first phase:", errors)
		// No output, because errors == nil
	}

	// Second phase
	errors.Append(oops("d")) // != nil
	errors.Append(oops("e")) // != nil
	errors.Append(okay("f")) // != nil, because an error has already occurred.

	// Handle pssible errors
	if errors.Err() != nil {
		// err.Error() == "2 errors: failed to d; failed to e"

		fmt.Println("errors in second phase:", errors)
		// Output: errors in second phase: 2 errors: failed to d; failed to e
	}
}
