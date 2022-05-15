package homework

func average(input [15]float32) (result float32) {

	var sum float32 = 0
	for _, item := range input {
		sum += item
	}
	return sum / 15
}
