package main

import (
	"fmt"
	"math"
	"sync"
)

type Figure interface {
	Area(float64)
	Perimeter(float64)
}

type Rectangle struct {
	firstSide  float64
	secondSide float64
}

type Triangle struct {
	firstSide  float64
	secondSide float64
	thirdSide  float64
}

type Circle struct {
	diameter float64
}

func (t Triangle) Perimeter() (perimeter float64) {
	if t.firstSide == t.secondSide && t.firstSide == t.thirdSide {
		perimeter = 3 * t.firstSide
	} else if t.firstSide == t.secondSide {
		perimeter = 2*t.firstSide + t.thirdSide
	} else if t.firstSide == t.thirdSide {
		perimeter = 2*t.firstSide + t.secondSide
	} else if t.secondSide == t.thirdSide {
		perimeter = 2*t.secondSide + t.firstSide
	} else {
		perimeter = (t.firstSide + t.secondSide + t.thirdSide)
	}
	fmt.Println("Triangle perimeter: ", perimeter)

	return perimeter
}

func (t Triangle) Area() (area float64) {
	p := t.Perimeter() / 2
	//исправить расчет и добавить проверку для разных треугольников
	area = math.Sqrt(p * (p - t.firstSide) * (p - t.secondSide) * (p - t.thirdSide))
	fmt.Println("Triangle area: ", area)

	return area
}

func (r Rectangle) Area() (area float64) {
	area = r.firstSide * r.secondSide
	fmt.Println("Rectangle area: ", area)

	return area
}

func (r Rectangle) Perimeter() (perimeter float64) {
	perimeter = 2 * (r.firstSide + r.secondSide)
	fmt.Println("Rectangle perimeter: ", perimeter)

	return perimeter

}

func (c Circle) Area() (area float64) {
	r := c.diameter / 2
	const PI float64 = 3.1415

	area = PI * math.Pow(r, 2)
	fmt.Println("Circle area: ", area)

	return area
}

func (c Circle) Perimeter() (perimeter float64) {
	r := c.diameter / 2
	const PI float64 = 3.1415

	perimeter = 2 * PI * r
	fmt.Println("Circle perimeter: ", perimeter)

	return perimeter
}

func main() {
	var rec Rectangle = Rectangle{

		firstSide:  34,
		secondSide: 21,
	}
	var cir Circle = Circle{
		diameter: 29,
	}
	var tri Triangle = Triangle{
		firstSide:  15,
		secondSide: 16,
		thirdSide:  29,
	}
	var wg sync.WaitGroup

	wg.Add(6)
	go func() {
		defer wg.Done()
		rec.Area()
	}()
	go func() {
		defer wg.Done()
		rec.Perimeter()
	}()

	go func() {
		defer wg.Done()
		cir.Area()
	}()

	go func() {
		defer wg.Done()
		cir.Perimeter()
	}()

	go func() {
		defer wg.Done()
		tri.Area()
	}()

	go func() {
		defer wg.Done()
		tri.Perimeter()
	}()

	wg.Wait()
}
