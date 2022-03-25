package user_models

type AuthUser struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}
