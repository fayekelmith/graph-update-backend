package models

// Bounty is a unit of work someone can claim and get paid for.
type Bounty struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Amount int    `json:"amount"`
	Status string `json:"status"`
}
