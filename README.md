[![GoDoc](https://godoc.org/github.com/shazow/go-multierror?status.svg)](https://godoc.org/github.com/shazow/go-multierror)

# go-multierror

There are [many multierror packages](https://godoc.org/?q=multierror) out there.

This go-multierror is compact yet has a very usable and versatile interface.

```
$ go get github.com/shazow/go-multierror
```

```go
import "github.com/shazow/go-multierror"
```

MultiError behaves like a normal Error whenever possible.

```go
errors := multierror.New()

errors == nil  // Just like a normal error, it can be nil by default.

errors.Add(nil)
errors == nil  // It's still nil, even if we add a nil error

anErr := errors.New("some error")
errors.Add(anErr)

errors != nil  // Once a non-nil error is added, it's non-nil just like normal errors.

// When there is only one error, the aggregated .Error() string is the same.
errors.Error() == anErr.Error()
```

MultiError is convenient to use in different scenarios.

```go
// Do a bunch of things that might fail independently and check for errors once in the end:
errors := multierror.New()

errors.Add(someFunc(...))
errors.Add(anotherFunc(...))
errors.Add(maybeThisOne(...))

if errors != nil {
	// Oh noes, we had a failure.
	log.Fatal(errors.Error())
	// .Error() will aggregate all the errors into one string, like:
	// "2 errors: this is one error; this is another error"
}
```

```go
errors := multierror.New()

// Can be used similar to how we do errors today:

// Instead of:
if err := maybeFail(); err != nil { ... }

// We can do:
if err := errors.Add(maybeFail()); err != nil { ... }

// Because errors.Add(nil) == nil, just like normal errors
```

[Check the godocs](https://godoc.org/github.com/shazow/go-multierror) for more 
details.


## License

MIT
