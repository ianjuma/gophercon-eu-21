package example1

// example1 receives its own copy of an integer value
//go:noinline
func foo(i int) {
	i++
}

// example2 receives a pointer to the original integer value j
//go:noinline
func bar(j *int) {
	*j++
}

// go:noinline pragma
// advantages and disadvantages of inlining
