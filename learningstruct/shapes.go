package learningstruct

import "math"

type Rectangle struct {
	Width float64
	Height float64
}


//func Perimeter(width float64,height float64) float64 {
//	return 2*(width+height)
//}
//
//func Area(width float64,height float64) float64 {
//	return width * height
//}

func Perimeter(rectangle Rectangle) float64{
	return 2*(rectangle.Width+rectangle.Height)
}

//func Area(rectangle Rectangle) float64{
//	return rectangle.Width*rectangle.Height
//}

func(r Rectangle) Area() float64{
	return r.Width*r.Height
}

type Circle struct {
	Radius float64
}

//func(receiverName ReceiverType) MethodName(args)
func (c Circle) Area() float64{
	return math.Pi * c.Radius *c.Radius
}

type Triangle struct {
	Base float64
	Height float64
}

func (c Triangle) Area() float64{
	return 0.5*c.Base*c.Height
}

type Shape interface {
	Area() float64
}