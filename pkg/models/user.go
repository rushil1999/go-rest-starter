package models

type Person struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	Password string `json:"password"`
}

type PersonDTO struct {
	Name string `json:"name"`
	Email string `json:"email"`
	phone string `json:"phone"`
	password string `json:"password"`
}

var UserData []Person

// func (p *Person) getId () string {
// 	return p.id
// }

// func (p *Person) getName () string {
// 	return p.name
// }

// func (p *Person) getEmail () string {
// 	return p.email
// }

// func (p *Person) getPhone () string {
// 	return p.phone
// }

// func (p *Person) setName (name string) string {
// 	return p.name = name
// }

// func (p *Person) getEmail (email string) string {
// 	return p.email = email
// }

// func (p *Person) getPhone (phone string) string {
// 	return p.phone = phone
// }
// func (p *Person) setId (id string) string {
// 	return p.id = id
// }


