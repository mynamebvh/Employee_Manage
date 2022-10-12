package auth

type RequestLogin struct {
	Email    string `json:"email" example:"mynamebvh@gmail.com" validate:"required"`
	Password string `json:"password" example:"hoangdz" validate:"required"`
}

type RequestSendCodeResetPassword struct {
	Email string `json:"email" example:"mynamebvh@gmail.com" validate:"required"`
}

type RequestResetPassword struct {
	Code        string `json:"code"`
	NewPassword string `json:"new_password"`
}
