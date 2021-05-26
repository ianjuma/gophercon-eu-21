package example7

type Printer interface {
	Print()
}

type Canon struct {
	name string
}

func (c Canon) Print() {}

func example1() {
	x1 := Canon{"x50"}
	var i1 Printer = x1
	i1.Print() // FIXED 1.16 - x1 no longer allocated
}

func example2() {
	x2 := Canon{"prisma"}
	var i2 Printer = &x2
	i2.Print() // FIXED 1.16 - x2 no longer allocated
}

// use this in diff go versions - show the fix
// devirtualizing interface call to concrete call
