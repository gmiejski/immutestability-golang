package inventory

import (
	"errors"
	"fmt"
)

type inMemoryInventory struct {
	movies map[int]int
}

func (i *inMemoryInventory) Rent(movieID int) error {
	ok, err := i.IsAvailableToRent(movieID)
	if err != nil {
		return err
	}
	if !ok {
		return errors.New("no more copies to rent right now")
	}
	i.movies[movieID] = i.movies[movieID] - 1
	return nil
}

func (i *inMemoryInventory) IsAvailableToRent(movieID int) (bool, error) {
	if count, ok := i.movies[movieID]; ok {
		return count > 0, nil
	}
	return false, errors.New(fmt.Sprintf("no movie with id %d", movieID))
}

func NewInMemoryInventory(movies map[int]int) InventoryService {
	return &inMemoryInventory{movies: movies}
}
