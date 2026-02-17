package main

import "fmt"

type Account struct {
	ID int
	Balance int64
}

func Transfer(from *Account, to *Account) {

}

func main(){
	petr := &Account{ID: 1, Balance: 50000}
	ivan := &Account{ID: 2, Balance: 10000}
	fmt.Printf("До перевода: Пётр(%d) - %d, Иван(%d) - %d\n", petr.ID,  petr.Balance, ivan.ID, ivan.Balance)
	Transfer(petr, ivan)
	fmt.Printf("После первода: Пётр(%d) - %d, Иван(%d) - %d", petr.ID,  petr.Balance, ivan.ID, ivan.Balance)
}