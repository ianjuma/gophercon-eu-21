package example4

//go:noinline
func func1() {
	var i int
	fn := getFunc()
	fn(&i, 42) // BAD: i escapes: Indirect Call Flaw
}

func getFunc() func(p *int, x int) {
	return func(p *int, x int) {
		*p = x
	}
}

//go:noinline
func func2() {
	var j int
	p := foo  // assignment and indirect call
	p(&j, 42) // FIXED 1.16 - j no longer allocated -flaw name?
}

func foo(p *int, x int) {
	*p = x
}
