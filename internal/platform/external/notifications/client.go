package notifications

type EmailClient struct{}

func NewEmailClient() *EmailClient {
	return &EmailClient{}
}
func (e *EmailClient) SendWelcomeEmail(userID int) error {
	// HTTP call here
	return nil
}
