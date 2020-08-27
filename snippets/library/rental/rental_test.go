package rental

import (
	"github.com/stretchr/testify/require"
	"immutestable-golang/snippets/library/inventory"
	"testing"
)

var userID = 1
var movieID = 1

func TestRentedMovieIsSavedAsUsersRentedMovies(t *testing.T) {
	// given
	movieID := 100
	availableMovies := map[int]int{movieID: 1}
	rentalFacade := NewRentalFacade(inventory.NewInMemoryInventory(availableMovies)) // create a new facade

	// when
	err := rentalFacade.Rent(userID, movieID)

	// then
	require.NoError(t, err)
	rentedMovies, err := rentalFacade.ListRented(userID)
	require.NoError(t, err)
	require.ElementsMatch(t, []RentedMovie{{
		MovieID: movieID,
		UserID:  userID,
	}}, rentedMovies.Movies)
}

func TestCannotRentAMovieWhenThereAreNoCopiesAvailable(t *testing.T) {
	// given
	movieID := 100
	availableMovies := map[int]int{movieID: 1}
	rentalFacade := NewRentalFacade(inventory.NewInMemoryInventory(availableMovies))
	err := rentalFacade.Rent(userID, movieID)
	require.NoError(t, err)

	// when
	err = rentalFacade.Rent(userID, movieID)

	// then
	require.Error(t, err)
	require.Equal(t, "cannot rent a movie with ID 100. Reason: no more copies to rent", err.Error()) // don't check error messages
}
