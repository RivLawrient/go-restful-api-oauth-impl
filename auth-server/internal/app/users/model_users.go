package users

type NewUsersRequest struct {
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	BirthDate int64  `json:"birth_date"`
}
