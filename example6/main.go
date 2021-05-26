package example6

type Printer interface {
	Print()
}

type Canon struct {
	name string
}

func (c Canon) Print() {}

func example1() {
	x1 := Canon{"m50"}
	polymorphic(x1)
}

func example2() {
	x2 := Canon{"prisma"}
	polymorphic(&x2)
}

//go:noinline
func polymorphic(i Printer) {
	i.Print() // BAD: value passed escapes because i escapes
}

// i escapes because of interface call
