# Dependency Injection

Dependency injection in Go helps testing functions or methods with
side effects.

The side-effect causing functions are injected into the function-to-be-tested
as an argument, so that it can be replaced with something that can be watched
in a test situation.

An example is printing something to either stdOut or the web. Expose the
'printing device' as an argument of type `io.Writer`, and inject a `bytes.Buffer`
in a test situation so you can check with `buffer.String()` what was written
to it.
