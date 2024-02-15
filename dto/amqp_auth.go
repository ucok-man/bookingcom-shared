package dto

type AmqpAuthMessageDTO struct {
	ReceiverEmail string `json:"receiver_email"`
	Username      string `json:"username"`
	VerifyLink    string `json:"verify_link"`
	ResetLink     string `json:"reset_link"`
	Template      string `json:"template"`
}
