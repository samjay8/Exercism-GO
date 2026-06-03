package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	units := map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen":    6,
		"dozen":              12,
		"small_gross":        120,
		"gross":              144,
		"great_gross":        1728,
	}
	return units
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	bill := map[string]int{}
	return bill
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	Val, ok := units[unit]
	if ok {
		bill[item] = bill[item] + Val
		return true
	}
	return false
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	billcheck, ok1 := bill[item]
	unitcheck, ok2 := units[unit]
	if !ok1 || !ok2 {
		return false
	}
	quantity := billcheck - unitcheck
	if quantity < 0 {
		return false
	} else if quantity == 0 {
		delete(bill, item)
		return true
	}
	bill[item] = quantity
	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	billcheck, ok := bill[item]
	if ok {
		return billcheck, true
	}
	return 0, false
}
