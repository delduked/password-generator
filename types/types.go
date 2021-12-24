package types

type NewPasswordRequest struct {
	Length  int  `json:"length" query:"length"`
	Lower   bool `json:"lower" query:"lower"`
	Upper   bool `json:"upper" query:"upper"`
	Number  bool `json:"number" query:"number"`
	Special bool `json:"special" query:"special"`
}

type NewPasswordResponse struct {
	Status   int    `json:"status" query:"status"`
	Error    error  `json:"error" query:"error"`
	Password string `json:"password" query:"password"`
}
type Response struct {
	Status int   `json:"status"`
	Error  error `json:"error"`
}
type SavedFields struct {
	Key      string `json:"key"`
	Account  string `json:"account"`
	Username string `json:"username"`
	Password string `json:"password"`
}
type NewPasswordReqSave struct {
	Account  string `json:"account"`
	Username string `json:"username"`
	Password string `json:"password"`
}
type ModelResponse struct {
	Error    error  `json:"error"`
	Account  error  `json:"account"`
	Username string `json:"username"`
	Password string `json:"password"`
}
type AllPasswordResponse struct {
	Status    int   `json:"status"`
	Error     error `json:"error"`
	Passwords []SavedFields
}
