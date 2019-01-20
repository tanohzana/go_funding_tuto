package funding_tuto

type FundServer struct {
	commands chan TransactionCommand
	fund     *Fund
}

type Transactor func(interface{})

type TransactionCommand struct {
	Transactor Transactor
	Done       chan bool
}

func (s *FundServer) Balance() int {
	var balance int
	s.Transact(func(managedValue interface{}) {
		fund := managedValue.(*Fund)
		balance = fund.Balance()
	})

	return balance
}

func (s *FundServer) Withdraw(amount int) {
	s.Transact(func(managedValue interface{}) {
		fund := managedValue.(*Fund)
		fund.Withdraw(amount)
	})
}

func (s *FundServer) Transact(transactor Transactor) {
	command := TransactionCommand{
		Transactor: transactor,
		Done:       make(chan bool),
	}

	s.commands <- command
	<-command.Done
}

type WithdrawCommand struct {
	Amount int
}

type BalanceCommand struct {
	Response chan int
}

func NewFundServer(initialBalance int) *FundServer {
	server := &FundServer{
		commands: make(chan TransactionCommand),
		fund:     NewFund(initialBalance), // pull request ?
	}

	go server.loop()

	return server
}

func (s *FundServer) loop() {
	for transaction := range s.commands {
		transaction.Transactor(s.fund) // pull request
		transaction.Done <- true
	}
}
