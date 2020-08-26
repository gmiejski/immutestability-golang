package given

import (
	"github.com/stretchr/testify/assert"
	"immutestable-golang/snippets/shopping"
	"testing"
)

var testUser = 1

func TestAddingItemToShoppingCart_single_level_of_abstraction(t *testing.T) {
	// given
	shoppingCartFacade := shopping.NewShoppingCartFacade()

	// when
	shoppingCartFacade.Add(testUser, shopping.CartItem{ProductID: 1, Count: 2})

	// then
	shoppingCart := shoppingCartFacade.GetAllItems(testUser)
	assert.ElementsMatch(t, []shopping.CartItem{{1, 2}}, shoppingCart)
}

func TestAddingItemToShoppingCart_mixing_levels_of_abstractions(t *testing.T) {
	// given
	shoppingCartFacade := shopping.NewShoppingCartFacade_with_internals()

	// when
	shoppingCartFacade.Carts[1] = []shopping.CartItem{{1, 2}} // Here we access structure's internals

	// then
	shoppingCart := shoppingCartFacade.GetAllItems(testUser)
	assert.ElementsMatch(t, []shopping.CartItem{{1, 2}}, shoppingCart)
}
