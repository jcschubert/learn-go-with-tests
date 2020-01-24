# Pointers and Errors

## Go arguments are copied
If you call a function with a struct as argument, that struct is copied. You can
check this by printing the address with `&{variable-name}`.

``` golang
struct Point {
    x float64
    y float64
}


func (p Point) Distance(p2 Point) {
    fmt.Printf("address of p is %v\n", &p)
}

point := Point{10.0,10.0}
fmt.Printf("address of point is %v\n", &point)
```
