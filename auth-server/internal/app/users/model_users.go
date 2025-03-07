package users

type NewUsersRequest struct {
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	BirthDate int64  `json:"birth_date"`
}

type NewUsersResponse struct {
	Id        string `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	BirthDate int64  `json:"birth_date"`
}

type LoginUsersRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginUsersResponse struct {
	TokenJwt string `json:"token_jwt"`
}

type JWTResponse struct {   
	Username  string `json:"username"`
	Email     string `json:"email"`
	BirthDate int64  `json:"birth_date"`
}
