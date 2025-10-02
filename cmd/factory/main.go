package main

func main() {
	println("=== Teste Factory Method ===")

	user1 := NewUser("Alice", 28, "admin")
	user2 := NewUser("Bob", 25, "normal")

	println("User 1:", user1.ValidateRole())
	println("User 2:", user2.ValidateRole())

}
