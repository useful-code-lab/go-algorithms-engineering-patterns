package main

import (
	"errors"
	"fmt"
)

type Account struct {
	ID      int
	Balance int64
}

func Transfer(from *Account, to *Account, amount int64) error {
	if amount < 0 {
		return errors.New("Сумма перевода не может быть отрицательной.")
	}

	if from.Balance < amount {
		return fmt.Errorf("Недостаточно денег для перевода на счёте %d", from.ID)
	}

	from.Balance -= amount
	to.Balance += amount

	return nil
}

func main() {
	petr := &Account{ID: 1, Balance: 50000}
	ivan := &Account{ID: 2, Balance: 10000}
	fmt.Printf("До перевода: Пётр(%d) - %d, Иван(%d) - %d\n", petr.ID, petr.Balance, ivan.ID, ivan.Balance)
	err := Transfer(petr, ivan, 5000)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("После перевода: Пётр(%d) - %d, Иван(%d) - %d\n", petr.ID, petr.Balance, ivan.ID, ivan.Balance)
	}

	err = Transfer(petr, ivan, -4000)
	if err != nil {
		fmt.Println(err)

	} else {
		fmt.Printf("После перевода: Пётр(%d) - %d, Иван(%d) - %d\n", petr.ID, petr.Balance, ivan.ID, ivan.Balance)
	}

	err = Transfer(petr, ivan, 774000)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("После перевода: Пётр(%d) - %d, Иван(%d) - %d\n", petr.ID, petr.Balance, ivan.ID, ivan.Balance)
	}

}
