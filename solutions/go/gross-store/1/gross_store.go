// Package gross provides functions to manage bills and unit measurements.
package gross

// Units returns available unit measurements.
func Units() map[string]int {
	return map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen":    6,
		"dozen":              12,
		"small_gross":        120,
		"gross":              144,
		"great_gross":        1728,
	}
}

// NewBill returns an empty bill.
func NewBill() map[string]int {
	return make(map[string]int)
}

// AddItem adds an item by unit; returns false if the unit is invalid.
func AddItem(bill, units map[string]int, item, unit string) bool {
	value, ok := units[unit]
	if !ok {
		return false
	}
	bill[item] += value
	return true
}

// RemoveItem removes an item by unit.
// Returns false if the item or unit is invalid, or if removal would go below zero.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	amount, ok := bill[item]
	if !ok {
		return false
	}

	value, ok := units[unit]
	if !ok {
		return false
	}

	switch {
	case amount < value:
		return false
	case amount == value:
		delete(bill, item)
	default:
		bill[item] = amount - value
	}
	return true
}

// GetItem returns the quantity of an item and whether it exists.
func GetItem(bill map[string]int, item string) (int, bool) {
	qty, ok := bill[item]
	return qty, ok
}
