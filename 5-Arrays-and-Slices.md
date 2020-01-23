# Arrays and Slices

## Array declarations

Arrays have a 'fixed capacity' which is defined when declaring the array

```golang
// arraytype: [5]int
numbers := [5]int{1,2,3,4,5}
```

## Array access

```golang
// Elements are accessed via the standard name[index]
a := numbers[1]
numbers[3] = 1
```

## Using range to iterate over arrays

```golang
// range returns index, value. If index is not needed, we can discard it with _
for _, number := range numbers {
    sum += number
}
```

## Array types

The size of an array is encoded in its type. `[4]int` is different than `[5]int`.

## Slices

Slices are like arrays, but do not encode the size as part of their type. Seems to
work the same way as array, but without the size. so: `[]int` is a slice an arbitrary amount of ints.

But: You can't use an `[4]int` where a `[]int` is expected.

## Test Coverage

`go test -cover`

## Variadic functions

Go allows the definition of 'variadic functions', which take an arbitrary amount of
elements of the specified type.

```golang
    func Sum(numbers ... int) (sum int) {
        for _, num := range numbers {
            sum += num
        }
        return
    }
```

## Creating slices with `make`

You can create a slice with a specified starting capacity.

```golang
    // sums has a starting capacity of 10
    sums := make([]int, 10)
```

## append

Usually you don't want to worry about the slice capacity, so you should use `append`

```golang
    sums = append(sums, Sum(numbers))
```
