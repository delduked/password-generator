package schema

type Response struct {
	Status int   `json:"Status"`
	Error  error `json:"Error"`
}
