package main

// START BSORTG OMIT
type AlphaNum interface {
	int | int8 | int16 | int32 | int64 | float32 | float64 | byte | string
}

func BubbleSortAlphaNum[T AlphaNum](a *[]T) {
	// END BSORTG OMIT
	alen := len((*a))
	// loop to access each array element
	for i := 0; i < alen; i++ {
		// loop to compare array elements
		for j := 0; j < (alen - i - 1); j++ {
			// compare two adjacent elements
			// change > to < to sort in descending order
			if (*a)[j] > (*a)[j+1] {
				// swapping elements if elements
				// are not in the intended order
				temp := (*a)[j]
				(*a)[j] = (*a)[j+1]
				(*a)[j+1] = temp
			}
		}
	}
}
