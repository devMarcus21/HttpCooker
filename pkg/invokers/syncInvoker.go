package invokers

func InvokeSync(callback func() int, calls int) map[int]int {
	results := map[int]int{}
	for call := 0; call < calls; call++ {
		res := callback()
		results[res] = results[res] + 1
	}

	return results
}
