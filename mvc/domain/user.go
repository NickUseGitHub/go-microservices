package domain

type User struct {
	Id       int64  `json:"id"`
	Title    string `json:"title"`
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
}
