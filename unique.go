package mapvsslice

// UniqueWithMap finds the unique strings using an algorithm consisting of map
func UniqueWithMap(input []string) []string {
	intermediateOutput := make(map[string]struct{})
	output := make([]string, 0, len(input))

	for _, element := range input {
		intermediateOutput[element] = struct{}{}
	}

	for element := range intermediateOutput {
		output = append(output, element)
	}

	return output
}

// UniqueWithSlice finds the unique strings using an algorithm consisting of slice
func UniqueWithSlice(input []string) []string {
	output := make([]string, 0, len(input))

	for _, element := range input {
		i := findIndexOfElement(output, element)
		if i == -1 {
			output = append(output, element)
		}
	}

	return output
}

func findIndexOfElement(input []string, elementToFind string) int {
	for i, element := range input {
		if element == elementToFind {
			return i
		}
	}

	return -1
}
