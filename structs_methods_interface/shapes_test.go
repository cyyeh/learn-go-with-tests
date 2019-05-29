package structs_methods_interface

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	/*
		checkArea := func(t *testing.T, shape Shape, want float64) {
			t.Helper()
			got := shape.Area()
			if got != want {
				t.Errorf("got %.2f want %.2f", got, want)
			}
		}

		t.Run("rectangles", func(t *testing.T) {
			rectangle := Rectangle{12.0, 6.0}
			checkArea(t, rectangle, 72.0)
		})

		t.Run("circles", func(t *testing.T) {
			circle := Circle{10.0}
			checkArea(t, circle, 314.1592653589793)
		})
	*/

	/*
		Table driven tests are useful when you want to build a list of test cases
		that can be tested in the same manner.

		Table based tests can be a great item in your toolbox but be sure that
		you have a need for the extra noise in the tests. If you wish to test various
		implementations of an interface, or if the data being passed in to a function has
		lots of different requirements that need testing then they are a great fit.
	*/
	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, hasArea: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, hasArea: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v got %.2f want %.2f", tt.shape, got, tt.hasArea)
			}
		})
	}
}
