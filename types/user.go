package types

type User struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	Name      string `json:"name"`
	Role      string `json:"role" default:"user"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type RegisterPayload struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type LoginPayload struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password"`
}
