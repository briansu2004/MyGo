package main

import "fmt"

// Enum definition using const and iota
type Color int

const (
    Red    Color = iota // 0
    Green               // 1 (automatically increments)
    Blue                // 2 (automatically increments)
    Black               // 3 (automatically increments)
    White               // 4 (automatically increments)
)

func (c Color) String() string {
    // Convert enum value to string
    names := [...]string{
        "Red",
        "Green",
        "Blue",
        "Black",
        "White",
    }

    if c < Red || c > White {
        return "Unknown"
    }

    return names[c]
}

func main() {
    // Using enums
    var color1 Color = Red
    var color2 Color = Blue
    var color3 Color = Green
    var color4 Color = Black
    var color5 Color = White

    fmt.Println("Color 1:", color1) // Output: Color 1: Red
    fmt.Println("Color 2:", color2) // Output: Color 2: Blue
    fmt.Println("Color 3:", color3) // Output: Color 3: Green
    fmt.Println("Color 4:", color4) // Output: Color 4: Black
    fmt.Println("Color 5:", color5) // Output: Color 5: White
}
