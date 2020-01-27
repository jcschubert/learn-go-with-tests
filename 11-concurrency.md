# Concurrency

## Goroutines

```golang
go func(value string) {
    // do something
}()
```

Start a goroutine by putting `go` in front of a function, which can either
be anonymous or a named function

## Anonymous functions

```golang
func(){
    // do something
}()
```

A function declaration that is invoked directly with `()` after it.

## Channels

Example on creating and using a channel. Channels are commonly used
to handle results of concurrent operations. Maybe for communications
between goroutines, too?

```golang

// Channels will contain items of this type
type result struct {
    url string
    valid bool
}

// Make a channel with the type above
resultChannel := make(chan result)

// Put a result into the channel. Normally, you put this into a goroutine
resultChannel <- result{"http://www.google.de", true}

// Fetch an item from the channel
result := <-resultChannel

```

## Testing for race conditions

`go test` has a builtin race checker. Invoke with `go test -race`.
