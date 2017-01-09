package account

import "sync"

// Account is a bank account
type Account struct {
	balance int64
	closed  bool
	mu      sync.Mutex
}

// Open returns opened account
func Open(initialDeposit int64) *Account {
	if initialDeposit < 0 {
		return nil
	}
	acc := new(Account)
	acc.balance = initialDeposit
	return acc
}

// Balance return account's ballance
func (acc *Account) Balance() (balance int64, ok bool) {
	if acc.closed {
		return 0, false
	}
	return acc.balance, true
}

// Close closes account
func (acc *Account) Close() (payout int64, ok bool) {
	acc.mu.Lock()
	defer acc.mu.Unlock()
	if acc.closed {
		return 0, false
	}
	acc.closed = true
	return acc.balance, true
}

// Deposit adds deposit to account
func (acc *Account) Deposit(amount int64) (newBalance int64, ok bool) {
	acc.mu.Lock()
	defer acc.mu.Unlock()
	if acc.closed {
		return 0, false
	}
	newBalance = acc.balance + amount
	if newBalance < 0 {
		return acc.balance, false
	}
	acc.balance = newBalance
	return acc.balance, true
}
