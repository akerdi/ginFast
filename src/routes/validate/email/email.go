package email

type Email struct {
	 Email string `form:"email" json:"email" validate:"required"`
	 Subject string `form:"subject" json:"subject" validate:"required"`
	 Msg string `form:"msg" json:"msg" validate:"required"`
}
