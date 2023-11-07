package user

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	FullName string `json:"full_name"`
	Phone    string `json:"phone"`
}
