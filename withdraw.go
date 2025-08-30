package main

import "fmt"

// Card represents a bank card with its properties.
type Card struct {
	Number   string
	IsActive bool
	Balance  float64
}

// Withdraw allows a user to withdraw money from a card.
func Withdraw(card *Card, amount float64) {
	if !card.IsActive {
		fmt.Println("Error: The card is not active. The operation is not possible.")
		return
	}

	if amount <= 0 {
		fmt.Println("Error: The withdrawal amount must be greater than zero.")
		return
	}

	if card.Balance >= amount {
		card.Balance -= amount
		fmt.Printf("Withdrawal successful. Current balance: %.2f somoni.\n", card.Balance)
	} else {
		fmt.Println("Insufficient funds.")
	}
}

// Deposit allows a user to deposit money into a card.
func Deposit(card *Card, amount float64) {
	// Requirements for the Deposit function:
	// 1. Money can only be deposited into an active card.
	// 2. Money can be deposited even with a negative balance.
	// 3. A maximum of 50,000 somoni can be deposited at a time.

	// 1. Check if the card is active.
	if !card.IsActive {
		fmt.Println("Error: The card is not active. Deposit is not possible.")
		return
	}

	// Check if the amount is a positive number.
	if amount <= 0 {
		fmt.Println("Error: The deposit amount must be a positive number.")
		return
	}

	// 3. Check if the amount exceeds the limit.
	const maxDepositAmount = 50000.0
	if amount > maxDepositAmount {
		fmt.Printf("Error: A maximum of %.2f somoni can be deposited at a time.\n", maxDepositAmount)
		return
	}

	// If all checks pass, add the amount to the balance.
	// 2. This works even with a negative balance, as the `+=` operator handles it correctly.
	card.Balance += amount
	fmt.Printf("Deposit successful. Card %s received %.2f somoni. Current balance: %.2f somoni.\n", card.Number, amount, card.Balance)
}

func main() {
	// Create a sample active card
	myCard := &Card{
		Number:   "1234-5678-9012-3456",
		IsActive: true,
		Balance:  1000.0,
	}

	// Create a sample inactive card
	inactiveCard := &Card{
		Number:   "9876-5432-1098-7654",
		IsActive: false,
		Balance:  500.0,
	}

	fmt.Println("--- Testing Withdraw function ---")
	Withdraw(myCard, 500)
	Withdraw(myCard, 1000) // Insufficient funds

	fmt.Println("\n--- Testing Deposit function ---")
	// Successful deposit
	Deposit(myCard, 20000)

	// Deposit with a negative balance
	myCard.Balance = -500.0
	Deposit(myCard, 1000)

	// Attempt to deposit to an inactive card
	Deposit(inactiveCard, 1000)

	// Attempt to exceed the deposit limit
	Deposit(myCard, 60000)
}
