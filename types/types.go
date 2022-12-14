package types

type NewPasswordRequest struct {
	Length  int  `json:"Length" query:"Length"`
	Lower   bool `json:"Lower" query:"Lower"`
	Upper   bool `json:"Upper" query:"Upper"`
	Number  bool `json:"Number" query:"Number"`
	Special bool `json:"Special" query:"Special"`
}

type NewPasswordResponse struct {
	Status   int    `json:"Status" query:"Status"`
	Error    error  `json:"Error" query:"Error"`
	Password string `json:"Password" query:"Password"`
}
type Response struct {
	Status int   `json:"Status"`
	Error  error `json:"Error"`
}

type SavedField struct {
	Key      string `json:"Key"`
	Account  string `json:"Account"`
	Username string `json:"Username"`
	Password string `json:"Password"`
}
type NewPasswordReqSave struct {
	Account  string `json:"Account"`
	Username string `json:"username"`
	Password string `json:"Password"`
}
type KeyedField struct {
	Account  string `json:"Account" redis:"Account"`
	Username string `json:"Username" redis:"Username"`
	Password string `json:"Password" redis:"Password"`
}
type KeyedResponse struct {
	Status int   `json:"Status"`
	Error  error `json:"Error"`
	Fields KeyedField
}
type SavedFieldResponse struct {
	Status int   `json:"Status"`
	Error  error `json:"Error"`
	Field  SavedField
}
type SavedFieldsResponse struct {
	Status int   `json:"Status"`
	Error  error `json:"Error"`
	Fields []SavedField
}
