package inventory

type InventoryService interface {
	Rent(movieID int) error
	IsAvailableToRent(movieID int) (bool, error)
}
