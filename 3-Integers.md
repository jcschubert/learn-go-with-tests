# Integers

## Packages

Go source files can only have one package per directory. See [here](https://dave.cheney.net/2014/12/01/five-suggestions-for-setting-up-a-go-project).

## Documentation

Add comments over functions with `//` and they will appear in `godoc` documentation.

## Examples

Start with `Example` and are part of the `godoc` documentation. It is executed automatically as part of the testsuite. Note: If the output in the comment is wrong, the example fails, when it is executed, and with a proper error message.
