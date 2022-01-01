package schema

type JWT struct {
	Status int    `json:"Status"`
	Error  error  `json:"Error"`
	Valid  bool   `json:"Valid"`
	Bearer string `json:"Bearer"`
}
