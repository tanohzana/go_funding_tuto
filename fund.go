package funding_tuto

type Fund struct {
	balance int
}

func NewFund(initialBalance int) *Fund {
	return &Fund{
		balance: initialBalance,
	}
}

func (f *Fund) Balance() int {
	return f.balance
}

func (f *Fund) Withdraw(amount int) {
	f.balance -= amount
}
