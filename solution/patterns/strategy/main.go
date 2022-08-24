package main

import (
	"fmt"
	"math"
)

type Vector2D struct {
	X float64
	Y float64
}

func (v *Vector2D) GetMagnitude() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vector2D) Normalize() Vector2D {
	x := v.X / v.GetMagnitude()
	y := v.Y / v.GetMagnitude()
	return Vector2D{x, y}
}

func Sub(firstVector Vector2D, secondVector Vector2D) Vector2D {
	return Vector2D{firstVector.X - secondVector.X, firstVector.Y - secondVector.Y}
}

type MovePattern interface {
	Move(position Vector2D) Vector2D
}

type OneDirectionPattern struct {
	direction Vector2D
}

func NewOneDirectionPattern(direction Vector2D) *OneDirectionPattern {
	return &OneDirectionPattern{direction: direction.Normalize()}
}

func (od *OneDirectionPattern) Move(position Vector2D) Vector2D {
	position.X += od.direction.X
	position.Y += od.direction.Y
	return position
}

type PointFollowingPattern struct {
	pointPosition Vector2D
}

func NewPointFollowingPattern(pointPosition Vector2D) *PointFollowingPattern {
	return &PointFollowingPattern{pointPosition: pointPosition}
}

func (pf *PointFollowingPattern) Move(position Vector2D) Vector2D {
	direction := Sub(pf.pointPosition, position)
	direction = direction.Normalize()

	position.X += direction.X
	position.Y += direction.Y
	return position
}

type Entity struct {
	position Vector2D
	moving   MovePattern
}

func NewEntity(moving MovePattern) Entity {
	return Entity{Vector2D{0, 0}, moving}
}

func (e *Entity) Move() {
	e.position = e.moving.Move(e.position)
}

func (e Entity) String() string {
	return fmt.Sprintf("Current position of entity: (%f, %f)", e.position.X, e.position.Y)
}

func main() {
	first := NewEntity(NewOneDirectionPattern(Vector2D{1, 1}))
	second := NewEntity(NewPointFollowingPattern(Vector2D{2, 5}))

	for true {
		fmt.Println("First entity:", first)
		fmt.Println("Second entity:", second)

		first.Move()
		second.Move()
	}
}
