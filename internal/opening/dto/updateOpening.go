package dto

type UpdateOpening struct {
	Role     *string `json:"role" binding:"omitempty,min=10"`
	Company  *string `json:"company" binding:"omitempty,min=10"`
	Location *string `json:"location" binding:"omitempty,min=10"`
	Remote   *bool   `json:"remote" binding:"omitempty"`
	Link     *string `json:"link" binding:"omitempty,url"`
	Salary   *int64  `json:"salary" binding:"omitempty,gt=0"`
}
