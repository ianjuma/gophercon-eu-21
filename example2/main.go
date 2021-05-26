package example2

// example1 returns a new integer value
//go:noinline
func func1() int {
	i := 10
	return i
}

// example2 returns a pointer to the original integer value j
//go:noinline
func func2() *int {
	j := 10
	return &j
}

// go:noinline pragma

// talk about ownership models
// determination of lifetime - value still has to exist on return -> heap
