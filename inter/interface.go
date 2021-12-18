package inter

type Response struct {
	Password string `json:"password"`
}
type StatusResponse struct {
	Status int   `json:"status"`
	Error  error `json:"error"`
}
