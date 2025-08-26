package Card

import (
	"fmt"
	"image/color"
	"pkg/bank/card"
)
	type Card struct {
	ID       int
	PAN      string
	Balance  int
	Currency string
	Color    string
	Name     string
	Active   bool
}

func NewCard(currency string, color string, name string) types.Card {
	return types.Card{
		ID:       1000,
		PAN:      "5058 xxxx xxxx 0001",
		Balance:  0,
		Currency: currency,
		Color:    color,	
		Name:     name,
		Active:   true,
	}
}
