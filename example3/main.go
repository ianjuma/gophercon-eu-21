package example3

// example1 shares the backing array pointer back up
//go:noinline
func func1() []int {
	list1 := make([]int, 10)
	return list1
}

// example2
//go:noinline
func func2(size int) {
	list2 := make([]int, size)
	list2[0] = 1
}

// 1 - sharing the backing array pointer back up
// 2 - slice with variable size vs constant size slice and sharing

// 1 draw slice allocation
// returning a value that contains the value of the underlying backing array
