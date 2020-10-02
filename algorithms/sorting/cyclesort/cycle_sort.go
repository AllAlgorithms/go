package cyclesort

func Sort(slice []int) []int {
	var element, position int

	for index, count := 0, len(slice)-1; index < count; index++ {
		element, position = slice[index], index

		for i, iCount := index+1, len(slice); i < iCount; i++ {
			if slice[i] < element {
				position++
			}
		}

		if position == index {
			continue
		}

		for element == slice[position] {
			position++
		}

		if position != index {
			element, slice[position] = slice[position], element
		}

		for position != index {
			position = index

			for i, iCount := index+1, len(slice); i < iCount; i++ {
				if slice[i] < element {
					position++
				}
			}

			for element == slice[position] {
				position++
			}

			if element != slice[position] {
				element, slice[position] = slice[position], element
			}
		}
	}

	return slice
}
