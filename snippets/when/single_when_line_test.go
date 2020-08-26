package when

import (
	"github.com/stretchr/testify/assert"
	"immutestable-golang/snippets/shopping"
	"testing"
)

var testUser = 1

func TestRemovingSingleItemWithAllQuantityClearsShoppingCart(t *testing.T) {
	// given
	shoppingCartFacade := shopping.NewShoppingCartFacade()
	shoppingCartFacade.Add(testUser, shopping.CartItem{ProductID: 1, Count: 2})

	// when
	err := shoppingCartFacade.Remove(testUser, shopping.CartItem{ProductID: 1, Count: 2})

	// then
	assert.NoError(t, err)
	shoppingCart := shoppingCartFacade.GetAllItems(testUser)
	assert.Empty(t, shoppingCart)
}

func TestRemovingSingleItemWithAllQuantityClearsShoppingCart_many_when_sentences(t *testing.T) {
	// given
	shoppingCartFacade := shopping.NewShoppingCartFacade()

	// when
	shoppingCartFacade.Add(testUser, shopping.CartItem{ProductID: 1, Count: 2})
	_ = shoppingCartFacade.Remove(testUser, shopping.CartItem{ProductID: 1, Count: 2})
	shoppingCart := shoppingCartFacade.GetAllItems(testUser)

	// then
	assert.Empty(t, shoppingCart)
}
