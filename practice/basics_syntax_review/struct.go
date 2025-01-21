package main

import "fmt"

// Define a struct
type User struct {
    Name  string
    Age   int
    Email string
}

// Method for the User struct (value receiver)
func (u User) Greet() string {
    return fmt.Sprintf("Hello, my name is %s and I am %d years old.", u.Name, u.Age)
}

// Method for the User struct (pointer receiver)
func (u *User) UpdateEmail(newEmail string) {
    u.Email = newEmail
}

func main() {
    // Create a User instance
    user := User{Name: "Alice", Age: 25, Email: "alice@example.com"}

    // Call the Greet method
    fmt.Println(user.Greet())

    // Update email using the pointer receiver method
    user.UpdateEmail("newalice@example.com")
    fmt.Println("Updated Email:", user.Email)
}
