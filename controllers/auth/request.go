package auth

type RequestLogin struct {
	Email    string `json:"email" example:"mynamebvh@gmail.com" validate:"required"`
	Password string `json:"password" example:"hoangdz" validate:"required"`
}
