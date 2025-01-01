package entities

type Client struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	LastName       string `json:"lastName"`
	SecondLastName string `json:"secondLastName"`
	Email          string `json:"email"`
	MailingAddress string `json:"mailingAddress"`
	BillingAddress string `json:"billingAddress"`
}
