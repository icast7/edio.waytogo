package main

import "fmt"

type Shaper interface {
	Area() float32
}

type Square struct {
	side float32
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

type Rectangle struct {
	length, width float32
}

func (r Rectangle) Area() float32 {
	return r.length * r.width
}

func main() {
	r := Rectangle{10, 5}
	s := &Square{7}
	fmt.Println("Looping through shapes to calculate area...")

	//shapes := []Shaper{r, s}
	shapes := []Shaper{Shaper(r), Shaper(s)}
	for n, _ := range shapes {
		fmt.Println("Shape details:", shapes[n])
		fmt.Println("Area:", shapes[n].Area())
	}
}
