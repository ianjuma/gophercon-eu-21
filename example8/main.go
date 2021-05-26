package example8

//go:noinline
func foo() {
	m := make(map[int]*int)
	var i1 int
	m[0] = &i1 // BAD: escape, address-of slice or map assignment flaw
}

//go:noinline
func bar() {
	s := make([]*int, 1)
	var i2 int
	s[0] = &i2 // BAD: x2 escape, address-of slice or map assignment flaw
}
