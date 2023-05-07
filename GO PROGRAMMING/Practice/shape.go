package main

import (
    "fmt"
)

// shape interface defines methods for calculating area and volume
type shape interface {
    area() float64
    volume() float64
}

// square type implements the shape interface
type square struct {
    side float64
}

func (s square) area() float64 {
    return s.side * s.side
}

func (s square) volume() float64 {
    return 0
}

// rectangle type implements the shape interface
type rectangle struct {
    length float64
    width  float64
    height float64
}

func (r rectangle) area() float64 {
    return r.length * r.width
}

func (r rectangle) volume() float64 {
    return r.length * r.width * r.height
}

func main() {
    var s shape
    var length, width, height, side float64
    var shapeType string

    fmt.Println("Enter shape type (square/rectangle):")
    fmt.Scanln(&shapeType)

    switch shapeType {
    case "square":
        fmt.Println("Enter side length:")
        fmt.Scanln(&side)
        s = square{side: side}
    case "rectangle":
        fmt.Println("Enter length:")
        fmt.Scanln(&length)
        fmt.Println("Enter width:")
        fmt.Scanln(&width)
        fmt.Println("Enter height:")
        fmt.Scanln(&height)
        s = rectangle{length: length, width: width, height: height}
    default:
        fmt.Println("Invalid shape type!")
        return
    }

    fmt.Printf("Area of %s: %.2f\n", shapeType, s.area())
    fmt.Printf("Volume of %s: %.2f\n", shapeType, s.volume())
}
