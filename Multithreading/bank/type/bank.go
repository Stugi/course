package bank

import (
	"errors"
	"fmt"
)

type BankClient interface {
	// Deposit deposits given amount to clients account
	Deposit(amount int)

	// Withdrawal withdraws given amount from clients account.
	// return error if clients balance less the withdrawal amount
	Withdrawal(amount int) error

	// Balance returns clients balance
	Balance() int
}

type Client struct {
	balance int
}

func (c *Client) Deposit(amount int) {
	if amount <= 0 {
		return
	}
	c.balance += amount
	fmt.Println("Deposit", amount, "Summ", c.balance)
}

func (c *Client) Withdrawal(amount int) error {
	if amount <= 0 {
		return errors.New("сумма списания дб больше 0")
	}
	if c.balance < amount {
		return errors.New("недостаточно средств")
	}
	c.balance -= amount
	fmt.Println("Withdrawal", amount, "Summ", c.balance)
	return nil
}

func (c *Client) Balance() int {
	return c.balance
}
