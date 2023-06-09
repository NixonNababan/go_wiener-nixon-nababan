package payload

import uuid "github.com/satori/go.uuid"

type GetProfil struct {
	Name     string `json:"name" form:"name"`
	Username string `json:"username" form:"username"`
	Email    string `json:"email" form:"email"`
}

type GetItem struct {
	ID          uuid.UUID
	Name        string `json:"name" form:"name"`
	Description string `json:"description" form:"description"`
	Stock       int    `json:"stock" form:"stock"`
	Price       int    `json:"price" form:"price"`
	Category    string `json:"category" form:"category"`
}
