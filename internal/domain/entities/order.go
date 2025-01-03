package entities

import "time"

type Order struct {
	ID             string    `json:"id"`
	RelatedClient  string    `json:"relatedClient"`
	MailingAddress string    `json:"mailingAddress"`
	BillingAddress string    `json:"billingAddress"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	Products       []Product `json:"products"`
}
