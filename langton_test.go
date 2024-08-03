package langton

import (
	"testing"

	"github.com/pierrre/assert"
)

func TestNewGrid(t *testing.T) {
	g := NewGrid(Pt(10, 15), 2)
	assert.Equal(t, g.Size, Pt(10, 15))
	assert.Equal(t, len(g.Squares), 150)
	assert.Equal(t, g.States, 2)
	for y := range g.Size.Y {
		for x := range g.Size.X {
			assert.Equal(t, g.Get(Pt(x, y)), 0, assert.MessageWrapf("%d,%d", x, y))
		}
	}
}

func TestGridIndex(t *testing.T) {
	g := NewGrid(Pt(10, 15), 2)
	i := g.index(Pt(5, 5))
	assert.Equal(t, i, 55)
}

func TestGridGetSet(t *testing.T) {
	g := NewGrid(Pt(10, 15), 2)
	v := g.Get(Pt(5, 5))
	assert.Equal(t, v, 0)
	g.Set(Pt(5, 5), 1)
	v = g.Get(Pt(5, 5))
	assert.Equal(t, v, 1)
}

func TestGridGetInc(t *testing.T) {
	g := NewGrid(Pt(10, 15), 2)
	v := g.GetInc(Pt(5, 5))
	assert.Equal(t, v, 0)
	v = g.GetInc(Pt(5, 5))
	assert.Equal(t, v, 1)
	v = g.GetInc(Pt(5, 5))
	assert.Equal(t, v, 0)
}

func TestOrientationString(t *testing.T) {
	type TC struct {
		orientation Orientation
		expected    string
	}
	for _, tc := range []TC{
		{OrientationUp, "up"},
		{OrientationRight, "right"},
		{OrientationDown, "down"},
		{OrientationLeft, "left"},
		{Orientation(-1), "invalid (-1)"},
	} {
		res := tc.orientation.String()
		assert.Equal(t, res, tc.expected)
	}
}

func TestOrientationRotate(t *testing.T) {
	o := Orientation(0)
	for i, v := range []struct {
		rotate   int
		expected Orientation
	}{
		{
			rotate:   0,
			expected: OrientationUp,
		},
		{
			rotate:   1,
			expected: OrientationRight,
		},
		{
			rotate:   1,
			expected: OrientationDown,
		},
		{
			rotate:   1,
			expected: OrientationLeft,
		},
		{
			rotate:   1,
			expected: OrientationUp,
		},
		{
			rotate:   -1,
			expected: OrientationLeft,
		},
		{
			rotate:   -1,
			expected: OrientationDown,
		},
		{
			rotate:   -1,
			expected: OrientationRight,
		},
		{
			rotate:   -1,
			expected: OrientationUp,
		},
		{
			rotate:   5,
			expected: OrientationRight,
		},
		{
			rotate:   -5,
			expected: OrientationUp,
		},
	} {
		o = o.Rotate(v.rotate)
		assert.Equal(t, o, v.expected, assert.MessageWrapf("step %d", i))
	}
}

func TestAntMoveTurn(t *testing.T) {
	a := &Ant{Location: Pt(1, 1), Orientation: OrientationUp}
	a.Move(1)
	assert.Equal(t, a.Location, Pt(1, 0))
	a.Turn(1)
	assert.Equal(t, a.Orientation, OrientationRight)
	a.Move(1)
	assert.Equal(t, a.Location, Pt(2, 0))
	a.Turn(1)
	assert.Equal(t, a.Orientation, OrientationDown)
	a.Move(1)
	assert.Equal(t, a.Location, Pt(2, 1))
	a.Turn(1)
	assert.Equal(t, a.Orientation, OrientationLeft)
	a.Move(1)
	assert.Equal(t, a.Location, Pt(1, 1))
}

var _ Rule = RuleFunc(nil)

func TestRuleFunc(t *testing.T) {
	called := false
	r := RuleFunc(func(a *Ant) {
		called = true
	})
	r.Apply(&Ant{})
	assert.True(t, called)
}

func TestRuleTurnRightMove(t *testing.T) {
	a := &Ant{
		Location:    Pt(0, 0),
		Orientation: OrientationUp,
	}
	RuleTurnRightMove.Apply(a)
	assert.Equal(t, a.Location, Pt(1, 0))
	assert.Equal(t, a.Orientation, OrientationRight)
}

func TestRuleTurnLeftMove(t *testing.T) {
	a := &Ant{
		Location:    Pt(0, 0),
		Orientation: OrientationUp,
	}
	RuleTurnLeftMove.Apply(a)
	assert.Equal(t, a.Location, Pt(-1, 0))
	assert.Equal(t, a.Orientation, OrientationLeft)
}

func TestGameStep(t *testing.T) {
	g := &Game{
		Rules: RulesBasic,
		Grid:  NewGrid(Pt(50, 50), 2),
		Ants: []*Ant{
			{
				Location:    Pt(25, 25),
				Orientation: OrientationUp,
			},
		},
	}
	for range 100000 {
		g.Step()
	}
}

func TestNormalize(t *testing.T) {
	type TC struct {
		v, max, expected int
	}
	for _, tc := range []TC{
		{0, 10, 0},
		{5, 10, 5},
		{10, 10, 0},
		{-3, 10, 7},
		{-10, 10, 0},
	} {
		res := normalize(tc.v, tc.max)
		assert.Equal(t, res, tc.expected)
	}
}

func TestNormalizePoint(t *testing.T) {
	type TC struct {
		v, max, expected Point
	}
	for _, tc := range []TC{
		{Pt(0, 0), Pt(10, 10), Pt(0, 0)},
		{Pt(5, 3), Pt(10, 10), Pt(5, 3)},
		{Pt(10, 10), Pt(10, 10), Pt(0, 0)},
		{Pt(-5, -3), Pt(10, 10), Pt(5, 7)},
	} {
		res := normalizePoint(tc.v, tc.max)
		assert.Equal(t, res, tc.expected)
	}
}
