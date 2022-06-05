package golang_united_school_homework

import "math"

// Triangle must satisfy to Shape interface
type Triangle struct {
	Side float64
}

func (t Triangle) CalcPerimeter() float64 {
	return 3 * t.Side
}

func (t Triangle) CalcArea() float64 {
	halfPerimeter := t.CalcPerimeter() / 2
	return math.Sqrt(halfPerimeter * (halfPerimeter - t.Side) * (halfPerimeter - t.Side) * (halfPerimeter - t.Side))
}
