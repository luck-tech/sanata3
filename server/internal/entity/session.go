package entity

type Session struct {
	ID           string
	UserID       string
	AccessToken  string
	RefreshToken string
}
