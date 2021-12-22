package inter

type NewPasswordRequest struct {
	Length  int  `json:"length" query:"length"`
	Lower   bool `json:"lower" query:"lower"`
	Upper   bool `json:"upper" query:"upper"`
	Number  bool `json:"number" query:"number"`
	Special bool `json:"special" query:"special"`
}

type NewPasswordResponse struct {
	Status   int    `json:"status"`
	Error    error  `json:"error"`
	Password string `json:"password"`
}
type SaveResponse struct {
	Status int   `json:"status"`
	Error  error `json:"error"`
}
type SavedFields struct {
	Key      string `json:"key"`
	Account  error  `json:"account"`
	Username string `json:"username"`
	Password string `json:"password"`
}
type NewPasswordReqSave struct {
	Account  error  `json:"account"`
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
	Status    int `json:"status"`
	Passwords []SavedFields
}
