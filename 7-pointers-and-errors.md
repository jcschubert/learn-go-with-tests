# Pointers and Errors

## Go arguments are copied
If you call a function with a struct as argument, that struct is copied. You can
check this by printing the address with `&{variable-name}`.

## Pointers
To have side effects (to manipulate a struct), hand in the pointer. It is automatically dereferenced,
so you can use it normally in a function.

```golang
// a takes a pointer to a Test named b
func a(b *Test) {
    // pointer does not need to be dereferenced.
    b.method();
    // You can check the address of something like so:
    fmt.Printf("%v", &b)
}
```

## Errors

If unexpected cases happen, return errors from your function. Errors can be `nil`, so you need to check for that.
`var` allows you to export errors from your package, so they can be checked against by the clients of your package.

## Error Linting

There is a package to lint for unhandled errors. 