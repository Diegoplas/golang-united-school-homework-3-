package homework

func average(input [15]float32) (result float32) {
	for _, value := range input {
		if value != 0 {
			result += value
		}
	}
	result = result / float32(len(input))
	return
}
