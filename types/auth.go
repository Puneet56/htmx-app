package types

type LoginRequest struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

type LoginError struct {
	Username string
	Password string
}
