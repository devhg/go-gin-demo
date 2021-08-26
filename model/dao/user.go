package dao

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func CheckAuth(username, password string) (bool, error) {
	var user User
	err := db.Where("username=? AND password=?", username, password).First(&user).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
