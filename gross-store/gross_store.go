package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	gsum := make(map[string]int)
	gsum["quarter_of_a_dozen"] = 3
	gsum["half_of_a_dozen"] = 6
	gsum["dozen"] = 12
	gsum["small_gross"] = 120
	gsum["gross"] = 144
	gsum["great_gross"] = 1728

	return gsum
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	empty := make(map[string]int)

	return empty
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item string, unit string) bool {
	exists := false

	for s := range units {
		if s == unit {
			exists = true
			break
		}
	}

	if !exists {
		return false
	}

	bill[item] += Units()[unit]

	return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	itemInBill := false

	for s := range bill {
		if s == item {
			itemInBill = true
			break
		}
	}

	validUnit := false
	for x := range units {
		if x == unit {
			validUnit = true
			break
		}
	}

	if !validUnit || !itemInBill {
		return false
	}

	newQuantity := bill[item] - Units()[unit]
	if newQuantity < 0 {
		return false
	} else if newQuantity == 0 {
		delete(bill, item)

		return true
	}

	bill[item] -= Units()[unit]

	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	for s := range bill {
		if s == item {
			return bill[item], true
		}
	}
	return 0, false
}
