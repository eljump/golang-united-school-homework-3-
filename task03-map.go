package homework

import "sort"

func sortMapValues(input map[int]string) (result []string) {

	keysSlice := make([]int, 0, len(input))
	for key, _ := range input {
		keysSlice = append(keysSlice, key)
	}
	sort.Slice(keysSlice, func(i, j int) bool {
		return keysSlice[i] > keysSlice[j]
	})

	sorted := make([]string, 0, len(input))

	for _, value := range keysSlice {
		sorted = append(sorted, input[value])
	}
	return sorted
}
