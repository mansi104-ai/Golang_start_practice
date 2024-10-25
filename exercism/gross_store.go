package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	// panic("Please implement the Units() function")
    return map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen":    6,
		"dozen":              12,
		"small_gross":       120,
		"gross":              144,
		"great_gross":       1728,
	}
   }

// NewBill creates a new bill.
func NewBill() map[string]int {
	// panic("Please implement the NewBill() function")

    return make(map[string]int)
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	// panic("Please implement the AddItem() function")
    quantity ,exists := units[unit]// check if the unit exists in the units map
    if !exists{
        return false
    }
    bill[item] += quantity //Add the quantity of the item to the bill
    return true
}

// RemoveItem removes an item from customer bill.
// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	quantity, exists := units[unit] // Check if the unit exists in the units map
	if !exists {
		return false // If the unit doesn't exist, return false
	}

	// Check if the item exists in the bill
	if currentQty, inBill := bill[item]; inBill {
		if currentQty < quantity {
			return false // Cannot remove more than currently exists
		}
		
		// Reduce the quantity of the item in the bill
		bill[item] -= quantity
		if bill[item] <= 0 {
			delete(bill, item) // Remove the item if its quantity is zero or less
		}
		return true // Successfully removed item from the bill
	}

	return false // Item not found in the bill
}


// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	quantity ,exists := bill[item]
    return quantity,exists
}