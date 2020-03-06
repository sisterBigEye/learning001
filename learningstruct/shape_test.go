package learningstruct

import "testing"

func TestPerimeter(t *testing.T){
	//got := Perimeter(10.0,10.0)
	rectangle := Rectangle{10.0,10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want{
		t.Errorf("got %.2f want %.2f",got,want)
	}
}

func TestArea(t *testing.T){
	//got := Area(12.0,6.0)
	checkArea := func(t *testing.T,shape Shape,want float64){
		t.Helper()
		got := shape.Area()
		if got != want{
			t.Errorf("got %.2f want %.2f",got,want)
		}

	}
	t.Run("rectangles", func(t *testing.T) {
		retangle := Rectangle{12.0 ,6.0}
		checkArea(t,retangle,72.0)

	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		checkArea(t,circle,314.15926)
	})

}