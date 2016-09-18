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

## How is this better?

* Other implementations have an additional cast-to-Error function, whereas
  multierror doesn't need to be casted because it's an Error-satisfying type the
  entire time.

* `multiError` is an `[]error` type underneath, so a `multiError` can be casted
  back to a slice of errors and handled as desired. Handy if you want to
  customize the aggregate error formatting.

* `multiError` retains the "no-error is nil" semantics of normal Errors, and can
  be used as such without additional casting.

* It's fairly simple and short, less than 50 lines of actual code and lots of
  tests.


## Keep in mind

* It's not goroutine-safe out of the box, same as any other slice type.

* It uses a pointer receiver for `.Add` and a value receiver for `.Error`
  in order to satisfy the Error semantics. Please open an issue if there are
  specific problems in this case.

* The `multiError` type is not exposed and can only be assigned using `.New()`,
  this is because `multiError{} != nil` which breaks the contract that this
  library is providing.

[Check the godocs](https://godoc.org/github.com/shazow/go-multierror) for more
details.


## License

MIT
