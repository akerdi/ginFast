package email

type Email struct {
	 Email string `form:"email" json:"email" validate:"required"`
}
