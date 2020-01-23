# Iteration

Go has no `while` or `do ... until`. Iteration is done with `for` only.

## Declaring vs declaring/initializing shorthand

```golang
// Declares a variable, but does not initialize it
var declared string
```

```golang
// Declares and initializes a variable (short hand, typ is inferred from value)
initialized := "A string"
```

## Shorthand operators

```golang
// Go has the same shorthand operators that most other languages have
a := 1
a += 1 // shorthand for adding 1, and re-assigning a
```

## Benchmarks

Go has a built-in benchmark tool.

```golang
func BenchmarkRepeat(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Repeat("a")
    }
}
```

Note: we use `testing.B` instead of `testing.T`. `B` provides us with `N`, which
the test frameworks uses to configure the amount of repetitions.

Output:

```bash
Running tool: C:\Go\bin\go.exe test -benchmem -run=^$ github.com\jcschubert\learn-go-with-tests\iteration -bench ^(BenchmarkRepeat)$

goos: windows
goarch: amd64
pkg: github.com/jcschubert/learn-go-with-tests/iteration
BenchmarkRepeat-8   	10938315	       109 ns/op	      16 B/op	       4 allocs/op
PASS
ok  	github.com/jcschubert/learn-go-with-tests/iteration	1.440s
Success: Benchmarks passed
```