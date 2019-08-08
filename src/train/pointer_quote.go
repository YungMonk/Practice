package train

import "fmt"

type point struct {
	x, y int
}

func (p point) getX() int {
	return p.x
}

func (p point) getY() int {
	return p.y
}

func (p *point) setX(ix int) {
	(*p).x = ix
}

func (p *point) setY(iy int) {
	(*p).y = iy
}

func (p *point) setY2(iy int) {
	p.y = iy
}

func (p point) setY3(iy int) {
	p.y = iy
}

func test1() {
	p := point{1, 2}
	fmt.Print(p, "\t")

	fmt.Print("x=", p.getX(), "\t")
	fmt.Print("y=", p.getY(), "\t")

	p.setX(3)
	p.setY(4)
	fmt.Print(p, "\t")

	(&p).setX(5)
	(&p).setY(6)
	fmt.Print(p, "\t")

	fmt.Println()
}

// setY() -> setY2()
func test2() {
	p := point{1, 2}
	fmt.Print(p, "\t")

	fmt.Print("x=", p.getX(), "\t")
	fmt.Print("y=", p.getY(), "\t")

	p.setX(3)
	p.setY2(4)
	fmt.Print(p, "\t")

	(&p).setX(5)
	(&p).setY2(6)
	fmt.Print(p, "\t")

	fmt.Println()
}

// setY() -> setY3()
func test3() {
	p := point{1, 2}
	fmt.Print(p, "\t")

	fmt.Print("x=", p.getX(), "\t")
	fmt.Print("y=", p.getY(), "\t")

	p.setX(3)
	p.setY3(4)
	fmt.Print(p, "\t")

	(&p).setX(5)
	(&p).setY3(6)
	fmt.Print(p, "\t")

	fmt.Println()
}

func Test4() {
	p := &point{1, 2}
	fmt.Print(p, "\t")

	fmt.Print("x=", p.getX(), "\t")
	fmt.Print("y=", p.getY(), "\t")

	p.setX(3)
	p.setY(4)
	fmt.Print(p, "\t")

	// calling method setX with receiver &p (type **point) requires explicit dereference
	(*p).setX(5)
	p.setY3(6)
	fmt.Print(p, "\t")

	fmt.Println()
}
