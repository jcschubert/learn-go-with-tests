# Creating a first Program

* make a directory under `$GOPATH/src/github.com/{your-user-id}/hello`.
* you can put `hello` in a subfolder, e.g. `learn-go-with-tests`.
* Note: For each chapter, make a new folder and choose an appropritate name.
* Note: Github seems to be mandatory in the go world.
* The VSCode Go Extension might suggest missing go commands to install. Do so.
* Create a `hello.go` in `$GOPATH/src/github.com/{your-user-id/learn-go-with-tests/hello`} and re-create the first example from [here](https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/hello-world).
  * `package` declaration
  * `import` statement for the `fmt` package
  * declaration of a mandatory `main` function
  * Noteworthy Syntax: No semicolons! Exported method names start with a capital letter!
* Run code with `F5`

## How to test

* How to  test?
  * Separate 'domain' code from side effects (in this case, printing).
  * In this case, domain code is supplying the hello world string.
    * Extract to function.
  * Note: Only an example, so we can see how tests work.
  * Tests reside in a separate file named `hello_test.go`
    * Name is convention
    * Same package
    * package testing
  * test is named `TestHello`. All tests start with `Test`
* Run test via VSCode (click ui above test definition), or run `go test` in your shell.
  * Note: Must be in the `hello` folder to do so.
* Test Conventions Summary
  * Needs to be in a file named `xxx_test.go`
  * Test function must start with `Test`.
  * Testing function takes one argument only, `t *testing.T`
    * t is our hook into the testing framework. See [here](https://golang.org/pkg/testing/#T).

## Godoc

* Launch with `godoc -http:8000`, then visit `localhost:8000/pkg`
  * Hm, doesn't work out of the box. Need to do `go get golang.org/x/tools/cmd/godoc` first. (It was removed, see [here](https://golang.org/doc/go1.13#godoc)).

## Iterating

* Test is in place, now we can expand/iterate.
* Change the test so it used the newly desired behavior
  * Naturally, you'll get a compile error when you change function arguments
* Fix compile errors by adjusting function interface
* Now you get a failing test &dash; implement the desired behaviour until the tests pass.

## Refactoring

* Use constants for strings you use in functions
  * Improves performance, because the string instance doesn't need to be created each function call.

## Multiple test cases with subtests

* you can use `t.Run("some description", func t *Testing.T) { ... })` within your test functions to create as many test cases as you like.
* You can share code between test cases. `*testing.T` must be an argument, and `t.Helper()` must be  called within the helper.
  * If it fails, the line reported will be in our function call and not the helper.
* note how the types for `got` and `want` are declared with one `string` instead of two
* also, `assertCorrectMessage` is defined in a JS-likeway as a variable
  * that is because it was declared within another function
  * the normal way does not work

## Iterating II

The cycle

* Write a test (that fails or does not compiled).
* Make the compiler pass
* Run the test, see what fails and check the error message is meaningful
* Write enough code to make the test pass
* Refactor

## Switch

* Works as expected, but syntax is rather clean
