package bubbleSort

func BubbleSort(arr []int) []int {

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			jValue := arr[j]
			jPlusOneValue := arr[j+1]

			if jValue > jPlusOneValue {
				arr[j] = jPlusOneValue
				arr[j+1] = jValue
			}
		}
	}
	return arr
}
