package models

type CreateUserRequest struct {
	Name string `json:"name" validate:"required,min=2,max=100"`
	Dob  string `json:"dob" validate:"required"`
}

type UpdateUserRequest struct {
	Name string `json:"name" validate:"required,min=2,max=100"`
	Dob  string `json:"dob" validate:"required"`
}
