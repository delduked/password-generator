package inter

type Request struct {
	Length  int  `json:"length" query:"length"`
	Lower   bool `json:"lower" query:"lower"`
	Upper   bool `json:"upper" query:"upper"`
	Number  bool `json:"number" query:"number"`
	Special bool `json:"special" query:"special"`
}

type Response struct {
	Status   int    `json:"status"`
	Error    error  `json:"error"`
	Password string `json:"password"`
}
