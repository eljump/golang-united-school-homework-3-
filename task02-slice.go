package homework

func reverse(input []int64) (result []int64) {

	res := make([]int64, len(input))
	j := 0
	for i := len(input) - 1; i >= 0; i-- {
		res[j] = input[i]
		j++
	}

	return res
}
