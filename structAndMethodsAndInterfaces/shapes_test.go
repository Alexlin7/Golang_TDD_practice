package shapes

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {

	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got '%.2f' want '%.2f'", got, want)
		}
	}
	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{12.0, 6.0}
		want := 72.0
		checkArea(t, rectangle, want)
	})

	t.Run("circles", func(t *testing.T) {
		circles := Circles{10}
		want := 314.1592653589793
		checkArea(t, circles, want)
	})
}

func TestArea2(t *testing.T) {

	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "Rectangle", shape: Rectangle{12, 6}, want: 72.0},
		{name: "Rectangle", shape: Circles{10}, want: 314.1592653589793},
		{name: "Rectangle", shape: Triangle{12, 6}, want: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.want {
				t.Errorf("got %g want %g", got, tt.want)
			}
		})
	}
}
