// Package langton provides an implementation of the Langton's ant.
package langton

import (
	"fmt"
)

// Point is a point on a [Grid].
type Point struct {
	X, Y int
}

// Pt returns a [Point].
func Pt(x, y int) Point {
	return Point{x, y}
}

// Grid is a grid of squares with multiple states.
//
// By convention, the X-axis goes from left to right, and the Y-axis goes from top to bottom.
type Grid struct {
	Size    Point
	Squares []uint8
	States  uint8
}

// NewGrid creates a new [Grid].
func NewGrid(size Point, states uint8) *Grid {
	return &Grid{
		Size:    size,
		Squares: make([]uint8, size.X*size.Y),
		States:  states,
	}
}

func (g *Grid) index(p Point) int {
	return p.Y*g.Size.X + p.X
}

// Get returns the value of a [Point].
func (g *Grid) Get(p Point) uint8 {
	return g.Squares[g.index(p)]
}

// Set sets the value of a [Point].
func (g *Grid) Set(p Point, v uint8) {
	g.Squares[g.index(p)] = v
}

// GetInc increments the value of a [Point] and return the previous value.
func (g *Grid) GetInc(p Point) uint8 {
	i := g.index(p)
	v := g.Squares[i]
	w := v + 1
	if w >= g.States {
		w = 0
	}
	g.Squares[i] = w
	return v
}

// Orientation is the orientation of an [Ant].
type Orientation int

// Orientation values.
//
// The 0 value is "up", and increase clockwise.
const (
	OrientationUp Orientation = iota
	OrientationRight
	OrientationDown
	OrientationLeft
)

func (o Orientation) String() string {
	switch o {
	case OrientationUp:
		return "up"
	case OrientationRight:
		return "right"
	case OrientationDown:
		return "down"
	case OrientationLeft:
		return "left"
	default:
		return fmt.Sprintf("invalid (%d)", o)
	}
}

// Rotate rotates the orientation.
func (o Orientation) Rotate(v int) Orientation {
	o += Orientation(v)
	if o < 0 {
		o = o%4 + 4
	}
	if o >= 4 {
		o %= 4
	}
	return o
}

// Ant is an ant on a [Grid].
type Ant struct {
	Location    Point
	Orientation Orientation
}

// Move moves the Ant for the current orientation.
func (a *Ant) Move(v int) {
	switch a.Orientation {
	case OrientationUp:
		a.Location.Y -= v
	case OrientationRight:
		a.Location.X += v
	case OrientationDown:
		a.Location.Y += v
	case OrientationLeft:
		a.Location.X -= v
	}
}

// Turn changes the [Orientation] of the [Ant].
func (a *Ant) Turn(v int) {
	a.Orientation = a.Orientation.Rotate(v)
}

// Rule is a moving instruction for an [Ant].
type Rule interface {
	Apply(a *Ant)
}

// RuleFunc is a [Rule] func.
type RuleFunc func(a *Ant)

// Apply implements [Rule].
func (f RuleFunc) Apply(a *Ant) {
	f(a)
}

// RuleTurnRightMove is a simple [Rule].
var RuleTurnRightMove = RuleFunc(func(a *Ant) {
	a.Turn(1)
	a.Move(1)
})

// RuleTurnLeftMove is a simple [Rule].
var RuleTurnLeftMove = RuleFunc(func(a *Ant) {
	a.Turn(-1)
	a.Move(1)
})

// RulesBasic are the basic rules of Langton's ant.
var RulesBasic = []Rule{
	RuleTurnRightMove,
	RuleTurnLeftMove,
}

// Game is a Langton's ant game.
type Game struct {
	Rules []Rule
	Grid  *Grid
	Ants  []*Ant
}

// Step runs the [Game] for 1 step.
func (g *Game) Step() {
	for _, a := range g.Ants {
		v := g.Grid.GetInc(a.Location)
		r := g.Rules[v]
		r.Apply(a)
		a.Location = normalizePoint(a.Location, g.Grid.Size)
	}
}

func normalize(v, maxValue int) int {
	if v < 0 {
		v = v%maxValue + maxValue
	}
	if v >= maxValue {
		v %= maxValue
	}
	return v
}

func normalizePoint(v, maxPoint Point) Point {
	v.X = normalize(v.X, maxPoint.X)
	v.Y = normalize(v.Y, maxPoint.Y)
	return v
}
