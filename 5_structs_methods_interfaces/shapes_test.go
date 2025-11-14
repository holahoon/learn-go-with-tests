package structsmethodsinterfaces

import "testing"

type AreaTest []struct {
	shape Shape
	want  float64
}

func TestArea(t *testing.T) {
	areaTests := AreaTest{
		{shape: Rectangle{Width: 12, Height: 6}, want: 72.10},
		{shape: Circle{Radius: 10}, want: 314.1592653589793},
		{shape: Triangle{Base: 12, Height: 6}, want: 36.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		want := tt.want

		if got != want {
			t.Errorf("%#v got %g want %g", tt.shape, got, want)
		}
	}

	// checkArea := func(t testing.TB, shape Shape, want float64) {
	// 	t.Helper()
	// 	got := shape.Area()

	// 	if got != want {
	// 		t.Errorf("got %g want %g", got, want)
	// 	}
	// }

	// t.Run("rectangles", func(t *testing.T) {
	// 	rectangle := Rectangle{12.0, 6.0}
	// 	checkArea(t, rectangle, 72.0)
	// })

	// t.Run("circles", func(t *testing.T) {
	// 	circle := Circle{10}
	// 	checkArea(t, circle, 314.1592653589793)
	// })
}
