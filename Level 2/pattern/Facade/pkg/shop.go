package pkg

import (
	"errors"
	"fmt"
	"time"
)

type Shop struct {
	Name     string
	Products []Product
}

func (shop Shop) Sell(user User, product string) error {
	fmt.Println("[Магазин] Запрос к пользователю, для получения остатка по карте")
	time.Sleep(time.Millisecond * 500)
	err := user.Card.CheckBalance()
	if err != nil {
		return err
	}

	fmt.Println(fmt.Sprintf("[Магазин] Проверка - может ли [%s] купить товар!", user.Name))
	time.Sleep(time.Millisecond * 500)
	for _, prod := range shop.Products {
		if prod.Name != product {
			continue
		}

		if prod.Price > user.GetBalance() {
			return errors.New("[Магазин] Недостаточно средств для покупки товара!")
		}

		fmt.Println(fmt.Sprintf("[Магазин] Товар - куплен | [%s]", prod.Name))
	}

	return nil
}
