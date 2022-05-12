package main

func reverse(input []int64) (result []int64) {
	inputLen := len(input)
	for idx := inputLen - 1; idx >= 0; idx-- {
		result = append(result, input[idx])
	}
	return
}
