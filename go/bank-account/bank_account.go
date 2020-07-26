package account

import (
	"sync"
)

// Account manages all deposits of a customer
type Account struct {
	open    bool
	balance int64
	mutex   *sync.RWMutex
}

// Open opens up a new account with some initial funds
func Open(initialDeposit int64) *Account {
	if initialDeposit < 0 {
		return nil
	}
	return &Account{
		open:    true,
		balance: initialDeposit,
		mutex:   &sync.RWMutex{},
	}
}

// Close closes the account.
// If the account is already closed, there's no payout and
// this operation fails.
func (a *Account) Close() (payout int64, ok bool) {
	a.mutex.Lock()
	defer a.mutex.Unlock()

	if !a.open {
		return 0, false
	}

	// empty out and close the account
	payout = a.balance
	a.open = false
	a.balance = 0

	return payout, true
}

// Balance returns the current balance of the open account
func (a *Account) Balance() (balance int64, ok bool) {
	a.mutex.RLock()
	defer a.mutex.RUnlock()

	return a.balance, a.open
}

// Deposit deposits the given amount to the open account and
// returns the updated balance.
func (a *Account) Deposit(amount int64) (newBalance int64, ok bool) {
	// default assumption is that the deposit failed or account is closed
	ok = false

	a.mutex.Lock()
	defer a.mutex.Unlock()

	if a.open {
		possibleNewBalance := a.balance + amount

		// If amount is negative, action is permitted only if withdrawl
		// amount is less than original balance.
		if possibleNewBalance >= 0 {
			ok = true
			a.balance = possibleNewBalance
		}
	}

	return a.balance, ok
}
