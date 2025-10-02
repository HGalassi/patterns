package main

type AdminUser struct {
	Name string
	Age  int
	Type string
}

type NormalUser struct {
	Name string
	Age  int
	Type string
}

type User interface {
	ValidateRole() string
}

func (a AdminUser) ValidateRole() string {
	return "Admin"
}

func (n NormalUser) ValidateRole() string {
	return "Normal"
}
func NewUser(name string, age int, userType string) User {
	if userType == "admin" {
		return AdminUser{
			Name: name,
			Age:  age,
			Type: userType,
		}
	}
	return NormalUser{
		Name: name,
		Age:  age,
		Type: userType,
	}
}
