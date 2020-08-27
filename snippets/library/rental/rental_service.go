package rental

import (
	"errors"
	"fmt"
	"immutestable-golang/snippets/library/inventory"
)

type RentedMovie struct {
	MovieID int
	UserID  int
}

type RentedMovies struct {
	Movies []RentedMovie
}

type RentalFacade interface {
	Rent(userID int, movieID int) error
	ListRented(userID int) (RentedMovies, error)
	Return(userID int, movieID int) error
}

type inMemoryRentalFacade struct {
	rentedMovies     map[int][]RentedMovie
	inventoryService inventory.InventoryService
}

func (r *inMemoryRentalFacade) Rent(userID int, movieID int) error {
	isAvailable, err := r.inventoryService.IsAvailableToRent(movieID)
	if err != nil {
		return errors.New(fmt.Sprintf("cannot rent a movie with ID %d. Reason: %s", movieID, err.Error()))
	}
	if !isAvailable {
		return errors.New(fmt.Sprintf("cannot rent a movie with ID %d. Reason: no more copies to rent", movieID))
	}
	if _, ok := r.rentedMovies[userID]; !ok {
		r.rentedMovies[userID] = make([]RentedMovie, 0)
	}
	err = r.inventoryService.Rent(movieID)
	if err != nil {
		return errors.New(fmt.Sprintf("cannot rent a movie with ID %d. Reason: %s", movieID, err.Error()))
	}
	r.rentedMovies[userID] = append(r.rentedMovies[userID], RentedMovie{MovieID: movieID, UserID: userID})
	return nil
}

func (r *inMemoryRentalFacade) ListRented(userID int) (RentedMovies, error) {
	if movies, ok := r.rentedMovies[userID]; ok {
		return RentedMovies{movies}, nil
	} else {
		return RentedMovies{nil}, nil
	}
}

func (r *inMemoryRentalFacade) Return(userID int, movieID int) error {
	panic("implement me")

}

func NewRentalFacade(inventoryService inventory.InventoryService) RentalFacade {
	return &inMemoryRentalFacade{rentedMovies: make(map[int][]RentedMovie), inventoryService: inventoryService}
}
