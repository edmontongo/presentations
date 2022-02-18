package main

// START BSORTI OMIT
func BubbleSortInt(a *[]int) {
	// END BSORTI OMIT
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

// START BSORTI64 OMIT
func BubbleSortInt64(a *[]int64) {
	// END BSORTI64 OMIT
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

// START BSORTF32 OMIT
func BubbleSortFloat32(a *[]float32) {
	// END BSORTF32 OMIT
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

// START BSORTF64 OMIT
func BubbleSortFloat64(a *[]float64) {
	// END BSORTF64 OMIT
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

// START BSORTS OMIT
func BubbleSortString(a *[]string) {
	// END BSORTS OMIT
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
