package schema

type UserAccount struct {
	Username string `json:"Username"`
	Password string `json:"Password"`
}
type SignUp struct {
	Username string `json:"Username"`
	Password string `json:"Password"`
	Secret   string `json:"Secret"`
}
