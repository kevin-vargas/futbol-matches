package model

type Foo struct {
	Bar string
}

// TODO: field validations
type User struct {
	Username string `json:"username" validate:"min=5,max=40,required"`
	Password string `json:"password" validate:"min=5,max=40,required"`
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
}
