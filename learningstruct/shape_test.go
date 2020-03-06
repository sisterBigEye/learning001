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
	t.Run("rectangles", func(t *testing.T) {
		retangle := Rectangle{12.0 ,6.0}
		got := retangle.Area()
		want := 72.0

		if got != want {
			t.Errorf("got %.2f want %.2f",got ,want)
		}

	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		got := circle.Area()
		want := 314.15926
		if got != want{
			t.Errorf("got %.2f want %.2f",got,want)
		}
	})

}