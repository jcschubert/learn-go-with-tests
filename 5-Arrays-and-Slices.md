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
