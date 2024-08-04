package generics

import (
	"errors"
	"fmt"
	"time"
)

func ChargeForLineItem[T lineItem](newItem T, oldItems []T, balance float64) ([]T, float64, error) {
	newBalance := balance - newItem.GetCost()
	if newBalance < 0 {
		return nil, 0.0, errors.New("insufficient funds")
	}
	oldItems = append(oldItems, newItem)
	return oldItems, newBalance, nil
}

type lineItem interface {
	GetCost() float64
	GetName() string
}

type subscription struct {
	userEmail string
	startDate time.Time
	interval  string
}

func (s subscription) GetName() string {
	return fmt.Sprintf("%s subscription", s.interval)
}

func (s subscription) GetCost() float64 {
	if s.interval == "monthly" {
		return 25.00
	}
	if s.interval == "yearly" {
		return 250.00
	}
	return 0.0
}

func ConstraintDemo() {
	var oldItems []subscription
	var balance float64 = 100.00
	var newItem subscription

	oldItems, balance, _ = ChargeForLineItem(newItem, oldItems, balance)
	fmt.Println(oldItems, balance)
}
