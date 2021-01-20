package domain

var (
	UsersSlice []User
)

type User struct {
	Uuid string `json:"id"`
	Name string `json:"username"`
	GameId string `json:"game_id"`
}

func RemovePlayer(username string) {
	for i := range UsersSlice {
		if UsersSlice[i].Name == username {
			UsersSlice = append(UsersSlice[:i], UsersSlice[i+1:]...)
			break
		}
	}

	return
}