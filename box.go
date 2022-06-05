package golang_united_school_homework

import "fmt"

var (
	nullShape  = fmt.Errorf("shape does not exist")
	outOfRange = fmt.Errorf("requested shape out of the range")
)

// box contains list of shapes and able to perform operations on them
type box struct {
	shapes         []Shape
	shapesCapacity int // Maximum quantity of shapes that can be inside the box.
}

// NewBox creates new instance of box
func NewBox(shapesCapacity int) *box {
	return &box{
		shapesCapacity: shapesCapacity,
	}
}

// AddShape adds shape to the box
// returns the error in case it goes out of the shapesCapacity range.
func (b *box) AddShape(shape Shape) error {
	if len(b.shapes) < b.shapesCapacity {
		b.shapes = append(b.shapes, shape)
		return nil
	}
	return fmt.Errorf("the box is full")
}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {
	if i <= len(b.shapes) {
		shape := b.shapes[i]
		if shape == nil {
			return nil, nullShape
		}
		return shape, nil
	} else {
		return nil, outOfRange
	}
}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {
	if i <= len(b.shapes) {
		shape := b.shapes[i]
		if shape == nil {
			return nil, nullShape
		}
		b.shapes = append(b.shapes[:i], b.shapes[i+1:]...)
		return shape, nil
	} else {
		return nil, outOfRange
	}
}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	if i <= len(b.shapes) {
		removedShape := b.shapes[i]
		if removedShape == nil {
			return nil, nullShape
		}
		b.shapes[i] = shape
		return removedShape, nil
	} else {
		return nil, outOfRange
	}
}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() float64 {
	var sum float64
	for _, shape := range b.shapes {
		sum += shape.CalcPerimeter()
	}
	return sum
}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() float64 {
	var sum float64
	for _, shape := range b.shapes {
		sum += shape.CalcArea()
	}
	return sum
}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {
	isCircleHere := false
	for index, shape := range b.shapes {
		_, ok := shape.(Circle)
		if ok {
			b.shapes = append(b.shapes[:index], b.shapes[index+1:]...)
			isCircleHere = true
		}
	}
	if isCircleHere {
		return nil
	}
	return fmt.Errorf("any circle was not found")
}
