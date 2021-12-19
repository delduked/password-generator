package inter

type Request struct {
	Length  int  `json:"length"`
	Lower   bool `json:"lower"`
	Upper   bool `json:"Upper"`
	Number  bool `json:"number"`
	Special bool `json:"special"`
}

type Response struct {
	Password string `json:"password"`
}
type StatusResponse struct {
	Status int   `json:"status"`
	Error  error `json:"error"`
}
