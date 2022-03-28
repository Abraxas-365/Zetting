package user_models

type Contact struct {
	Email                   string `json:"email"`
	Phone                   string `json:"phone"`
	Country                 string `json:"country"`
	IdentifierDocumentation string `json:"identifier_document"`
}
