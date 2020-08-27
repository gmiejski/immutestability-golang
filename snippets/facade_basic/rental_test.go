package facade_basic

import (
	"github.com/stretchr/testify/require"
	"immutestable-golang/snippets/library/rental"
	"testing"
)

var userID = 1
var movieID = 1

func TestRentedMovieIsSavedAsUsersRentedMovies(t *testing.T) {
	// skipped, because it's used just on basic example of how it works
	return
	// given
	rentalFacade := rental.NewRentalFacade(nil) // create a new facade

	// when
	err := rentalFacade.Rent(userID, movieID)

	// then
	require.NoError(t, err)
	rentedMovies, err := rentalFacade.ListRented(userID)
	require.NoError(t, err)
	require.ElementsMatch(t, []rental.RentedMovie{{
		MovieID: movieID,
		UserID:  userID,
	}}, rentedMovies.Movies)

}
