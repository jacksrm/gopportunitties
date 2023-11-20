package dto

type CreateOpening struct {
	Role     string `json:"role" binding:"required,min=10"`
	Company  string `json:"company" binding:"required,min=10"`
	Location string `json:"location" binding:"required,min=10"`
	Remote   *bool  `json:"remote" binding:"required"`
	Link     string `json:"link" binding:"required,url"`
	Salary   int64  `json:"salary" binding:"required,gt=0"`
}
