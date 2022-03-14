package transport

type InsertBook struct {
	Name    string `json:"bookName" validate:"required"`
	Creator string `json:"bookCreator" validate:"required"`
}

type UpdateBook struct {
	Name    string `json:"bookName"`
	Creator string `json:"bookCreator"`
}

// users

type InsertUser struct {
	Username string `json:"username" validate:"required"`
	Email string `json:"email" validate:"required"`
}

type UpdateUser struct {
	Username string `json:"username"`
	Email string `json:"email"`
}
