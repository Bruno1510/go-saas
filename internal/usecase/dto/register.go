package dto

type RegisterRequest struct {
	CompanyName string `json:"companyName"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
}