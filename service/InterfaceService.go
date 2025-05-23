package service

type EmailServiceInterface interface {
	SendEmail(email string) error
}
