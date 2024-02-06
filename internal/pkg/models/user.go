package models

type User struct {
	Username    string `db:"useername" json:"username"`
	Email       string `db:"email" json:"email"`
	Fullname    string `db:"fullname" json:"fullname"`
	Password    string `db:"password" json:"password"`
	Address     string `db:"address" json:"address"`
	PhoneNumber string `db:"phone_number" json:"phone_number"`
	ID          int64  `db:"id" json:"id"`
}

type UserRegisterRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Fullname string `json:"fullname"`
	Password string `json:"password"`
}

type UserRegisterResponse struct {
	Message string `json:"message"`
}

type UserLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
