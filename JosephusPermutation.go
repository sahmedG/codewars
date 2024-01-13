package kata

func Josephus(arr []interface{}, k int) []interface{} {
	// Your code:
	if len(arr) == 0 {
		return []interface{}{}
	}
	if len(arr) == 1 {
		return arr
	}

	var result []interface{}
	var n interface{}
	for len(arr) > 0 {
		arr, n = joseOne(arr, k)
		result = append(result, n)
	}

	return result
}

func joseOne(arr []interface{}, k int) ([]interface{}, interface{}) {
	k = k % len(arr)
	if k == 0 {
		return arr[:len(arr)-1], arr[len(arr)-1].(interface{})
	}

	return append(arr[k:], arr[:k-1]...), arr[k-1].(interface{})
}
