package main

import (
	"fmt"
	"sync"
)

type BankAccount struct {
	RWMutex sync.RWMutex
	Balance int
}

func (account *BankAccount) AddBalance(amount int) {
	account.RWMutex.Lock()
	defer account.RWMutex.Unlock()

	account.Balance = account.Balance + amount
}

func (account *BankAccount) GetBalance() int {
	account.RWMutex.RLock()
	defer account.RWMutex.RUnlock()

	balance := account.Balance

	return balance
}

func main() {
	account := BankAccount{}
	group := &sync.WaitGroup{}

	for i := 1; i <= 1000; i++ {
		go func(i int) {
			for j := 1; j <= 100; j++ {
				defer group.Done()
				group.Add(1)

				fmt.Printf("Before Balance %d : %d\n", i*100+j, account.GetBalance())
				account.AddBalance(1)
				fmt.Printf("After Balance %d : %d\n", i*100+j, account.GetBalance())
			}
		}(i)
	}

	group.Wait()
	fmt.Println("Final Balance : ", account.GetBalance())
}
