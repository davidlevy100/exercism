// Package cards demonstrates basic slice operations with card values.
package cards

// FavoriteCards returns a fixed slice of favorite cards: 2, 6, and 9.
func FavoriteCards() []int {
	return []int{2, 6, 9}
}

// GetItem returns the element at the given index.
// If the index is out of range, it returns -1.
func GetItem(slice []int, index int) int {
	if index >= 0 && index < len(slice) {
		return slice[index]
	}
	return -1
}

// SetItem updates the slice at the given index with value.
// If the index is out of range, the value is appended instead.
func SetItem(slice []int, index, value int) []int {
	if index >= 0 && index < len(slice) {
		slice[index] = value
		return slice
	}
	return append(slice, value)
}

// PrependItems returns a new slice with the given values added in front.
func PrependItems(slice []int, values ...int) []int {
	return append(values, slice...)
}

// RemoveItem returns a new slice with the item at the given index removed.
// If the index is out of range, the original slice is returned unchanged.
func RemoveItem(slice []int, index int) []int {
	if index >= 0 && index < len(slice) {
		return append(slice[:index], slice[index+1:]...)
	}
	return slice
}
