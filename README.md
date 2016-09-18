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
err := multierror.New()

err == nil  // Just like a normal error, it can be nil by default.

err.Add(nil)
err == nil  // It's still nil, even if we add a nil error

anError := errors.New("some error")
err.Add(anError)

err != nil  // Once a non-nil error is added, it's non-nil just like normal errors.

// When there is only one error, the aggregated .Error() string is the same.
err.Error() == anError.Error()
```

MultiError is convenient to use in different scenarios.

```go
// Do a bunch of things that might fail independently and check for errors once in the end:
err := multierror.New()

err.Add(makeNil(...))
err.Add(makeErr(...))
err.Add(makeErr(...))

if err != nil {
	// Oh noes, we had a failure.
    log.Fatal(err.Error())
	// .Error() will aggregate all the errors into one string, like:
	// "2 errors: this is one error; this is another error"
}
```

```go
// Can be used similar to how we do errors today:

// Instead of:
if err := maybeFail(); err != nil { ... }

// We can do:
err := multierror.New()

if err.Add(maybeFail()) != nil { ... }

// err.Add(nil) == nil and err.Add(error) != nil
```

[Check the godocs](https://godoc.org/github.com/shazow/go-multierror) for more 
details.


## License

MIT
