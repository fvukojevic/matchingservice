package domain

var (
	UsersMap = map[int]string{}
)

type User struct {
	Name string `json:"name"`
}
