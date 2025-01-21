package main

import "fmt"

func main() {
    // Create a map
    colors := map[string]string{
        "red":   "#FF0000",
        "green": "#00FF00",
        "blue":  "#0000FF",
    }
    fmt.Println("Colors map:", colors)

    // Accessing a value
    fmt.Println("Hex code for green:", colors["green"])

    // Adding a key-value pair
    colors["yellow"] = "#FFFF00"
    fmt.Println("After adding yellow:", colors)

    // Deleting a key
    delete(colors, "red")
    fmt.Println("After deleting red:", colors)

    // Check if a key exists
    hex, exists := colors["purple"]
    if exists {
        fmt.Println("Purple exists with hex:", hex)
    } else {
        fmt.Println("Purple does not exist.")
    }

    // Loop through a map
    fmt.Println("Looping through the map:")
    for key, value := range colors {
        fmt.Printf("Color: %s, Hex: %s\n", key, value)
    }
}

