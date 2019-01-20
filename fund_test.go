package funding_tuto

import (
	"sync"
	"testing"
)

const WORKERS = 10

func BenchmarkWithWithdrawals(b *testing.B) {
	if b.N < WORKERS {
		return
	}

	server := NewFundServer(b.N)

	dollarsPerFounder := b.N / WORKERS

	var wg sync.WaitGroup

	for i := 0; i < WORKERS; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			pizzaTime := false

			for i := 0; i < dollarsPerFounder; i++ {
				server.Transact(func(managedValue interface{}) {
					fund := managedValue.(*Fund)
					if fund.Balance() <= 10 {
						pizzaTime = true
						return
					}

					fund.Withdraw(1)
				})

				if pizzaTime {
					break
				}
			}
		}()
	}

	wg.Wait()

	balance := server.Balance()

	if balance != 10 {
		b.Error("Balance was not ten: ", balance)
	}
}
