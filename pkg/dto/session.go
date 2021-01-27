package dto

type LoginCredential struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Language  string `json:"language"`
	CaptchaId string `json:"captchaId"`
	Code      string `json:"code"`
	// console or api
	AuthMethod string `json:"authMethod"`
}

type SessionUser struct {
	UserId   string `json:"userId"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Language string `json:"language"`
	IsActive bool   `json:"isActive"`
	IsAdmin  bool   `json:"isAdmin"`
}

type Profile struct {
	User  SessionUser `json:"user"`
	Token string      `json:"token,omitempty"`
}

type Captcha struct {
	Image     string `json:"image"`
	CaptchaId string `json:"captchaId"`
}
