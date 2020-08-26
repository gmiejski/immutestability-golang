package shopping

import "errors"

type CartItem struct {
	ProductID int
	Count     int
}

type ShoppingCartFacade interface {
	GetAllItems(user int) []CartItem
	Add(user int, item CartItem)
	Remove(user int, item CartItem) error
}

func NewShoppingCartFacade() ShoppingCartFacade {
	return &inMemoryShoppingCardFacade{Carts: make(map[int][]CartItem)}
}

func NewShoppingCartFacade_with_internals() *inMemoryShoppingCardFacade {
	return &inMemoryShoppingCardFacade{Carts: make(map[int][]CartItem)}
}

type inMemoryShoppingCardFacade struct {
	Carts map[int][]CartItem
}

func (i *inMemoryShoppingCardFacade) Remove(user int, itemRemoved CartItem) error {
	if _, ok := i.Carts[user]; !ok {
		return errors.New("user has no shopping cart")
	}
	var newShoppingCart []CartItem

	for _, item := range i.Carts[user] {
		if itemRemoved.ProductID == item.ProductID {
			count := max(0, item.Count-itemRemoved.Count)
			if count > 0 {
				newShoppingCart = append(newShoppingCart, CartItem{ProductID: item.ProductID, Count: count})
			}
		} else {
			newShoppingCart = append(newShoppingCart, item)
		}
	}
	i.Carts[user] = newShoppingCart
	return nil
}

func (i *inMemoryShoppingCardFacade) GetAllItems(user int) []CartItem {
	if cart, ok := i.Carts[user]; ok {
		return cart
	}
	return []CartItem{}
}

func (i *inMemoryShoppingCardFacade) Add(user int, item CartItem) {
	if _, ok := i.Carts[user]; !ok {
		i.Carts[user] = []CartItem{}
	}
	i.Carts[user] = append(i.Carts[user], item)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
