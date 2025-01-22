package ch

import "fmt"

func ExampleNew() {
	type UserData struct {
		Name  string
		Email string
	}

	chStruct := New[UserData](3, nil)
	chStruct.AddNode("NodeA")
	chStruct.AddNode("NodeB")

	chStruct.AddKey("user123", UserData{Name: "Alice", Email: "alice@example.com"})
	user, exists := chStruct.GetKey("user123")
	if exists {
		fmt.Println("User Data:", user.Name, user.Email)
	}
}
