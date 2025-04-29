package request

type Login struct {
	Username string `json:"username" gorm:"unique"`
	Password string `json:"password"`
}