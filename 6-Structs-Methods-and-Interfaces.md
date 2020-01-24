# Structs, Methods and Interfaces

## Structs

Structs are defined like this:

```golang
type Rectangle struct {
  Width  float64
  Height float64
}
```
To create a new instance? of a struct, there are two ways: 

```golang
// without naming the fields. This just goes by field order, so r.Width = 12.0,
// r.Height = 6.0
r := Rectangle{12.0, 6.0}

// with naming the fields
r := Rectangle{Width: 12.0, Height: 6.0}
```

Note the curly braces, same as with arrays.

## Methods

Methods are assigned to structs with their definition, allowing the dot syntax to use them.

```golang
func (r Rectangle) Area {
    return r.Width + r.Height
}
```

## Interfaces

Interfaces describe which methods a struct must have to implement it. There is no `implements` syntax, Go's interface resolution is 'implicit'. As long as a struct has all methods required by the interface, it automatically implements it.  

An interface is a type, and can be used everywhere other types can.

```golang
type Shape interface {
    Area() float64
}
```

## Anonymous structs

These seem similar to JS object literals.

```golang

// define a slice holding anonymous structs with the described signature
shapes := []struct {
    name    string
    shape   Shape
} 

shapes.append({name: "Rectangle", shape: Rectangle{10.0, 10.0}})
shapes.append({name: "Circle", shape: Circle{10.0}})
shapes.append({name: "Triangle", shape: Triangle{10.0, 10.0}})
```
