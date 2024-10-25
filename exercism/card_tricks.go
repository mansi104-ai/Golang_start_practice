package cards

// FavoriteCards returns a slice with the cards 2, 6 and 9 in that order.
func FavoriteCards() []int {
	// panic("Please implement the FavoriteCards function")
    list := []int{2,6,9}
    return list
}

// GetItem retrieves an item from a slice at given position.
// If the index is out of range, we want it to return -1.
func GetItem(slice []int, index int) int {
	// panic("Please implement the GetItem function")
    if index < 0 || index >= len(slice) {
		return -1
	}
    card := slice[index]
    return card
}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(slice []int, index, value int) []int {
	// panic("Please implement the SetItem function")
    if index < 0 || index >= len(slice){
        slice = append(slice,value)
    }else{
        slice[index] = value
    }
    return slice
}

// PrependItems adds an arbitrary number of values at the front of a slice.
func PrependItems(slice []int, values ...int) []int {
	// Loop over the values and prepend each one
	for i := len(values) - 1; i >= 0; i-- {
		slice = append([]int{values[i]}, slice...)
	}
	return slice
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
	// Check if the index is within valid bounds
	if index < 0 || index >= len(slice) {
		return slice
	}
	// Remove the item at the given index
	return append(slice[:index], slice[index+1:]...)
}
