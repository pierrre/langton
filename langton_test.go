package langton

import "testing"

func TestNewGrid(t *testing.T) {
	g := NewGrid(Pt(10, 15), 2)
	if g.Size.X != 10 {
		t.Fatal("not equal")
	}
	if g.Size.Y != 15 {
		t.Fatal("not equal")
	}
	if len(g.Squares) != 150 {
		t.Fatal("not equal")
	}
	if g.States != 2 {
		t.Fatal("not equal")
	}
	for y := 0; y < g.Size.Y; y++ {
		for x := 0; x < g.Size.X; x++ {
			if g.Get(Pt(x, y)) != 0 {
				t.Fatalf("not blank: %d,%d", x, y)
			}
		}
	}
}

func TestGridSquareIndex(t *testing.T) {
	g := NewGrid(Pt(10, 15), 2)
	i := g.SquareIndex(Pt(5, 5))
	if i != 55 {
		t.Fatal("not equal")
	}
}

func TestGridGetSet(t *testing.T) {
	g := NewGrid(Pt(10, 15), 2)
	v := g.Get(Pt(5, 5))
	if v != 0 {
		t.Fatal("not equal")
	}
	g.Set(Pt(5, 5), 1)
	v = g.Get(Pt(5, 5))
	if v != 1 {
		t.Fatal("not equal")
	}
}

func TestGridGetInc(t *testing.T) {
	g := NewGrid(Pt(10, 15), 2)
	v := g.GetInc(Pt(5, 5))
	if v != 0 {
		t.Fatal("not equal")
	}
	v = g.GetInc(Pt(5, 5))
	if v != 1 {
		t.Fatal("not equal")
	}
	v = g.GetInc(Pt(5, 5))
	if v != 0 {
		t.Fatal("not equal")
	}
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
		if res != tc.expected {
			t.Errorf("%#v: got %s, want %s", tc, res, tc.expected)
		}
	}
}

func TestOrientationRotate(t *testing.T) {
	o := Orientation(0)
	if o != OrientationUp {
		t.Fatal("not equal")
	}
	o = o.Rotate(1)
	if o != OrientationRight {
		t.Fatal("not equal")
	}
	o = o.Rotate(1)
	if o != OrientationDown {
		t.Fatal("not equal")
	}
	o = o.Rotate(1)
	if o != OrientationLeft {
		t.Fatal("not equal")
	}
	o = o.Rotate(1)
	if o != OrientationUp {
		t.Fatal("not equal")
	}
	o = o.Rotate(-1)
	if o != OrientationLeft {
		t.Fatal("not equal")
	}
	o = o.Rotate(-1)
	if o != OrientationDown {
		t.Fatal("not equal")
	}
	o = o.Rotate(-1)
	if o != OrientationRight {
		t.Fatal("not equal")
	}
	o = o.Rotate(-1)
	if o != OrientationUp {
		t.Fatal("not equal")
	}
	o = o.Rotate(5)
	if o != OrientationRight {
		t.Fatal("not equal")
	}
	o = o.Rotate(-5)
	if o != OrientationUp {
		t.Fatal(o)
		t.Fatal("not equal")
	}
}

func TestAntMoveTurn(t *testing.T) {
	a := &Ant{Location: Pt(1, 1), Orientation: OrientationUp}
	a.Move(1)
	if a.Location != Pt(1, 0) {
		t.Fatal("not equal")
	}
	a.Turn(1)
	if a.Orientation != OrientationRight {
		t.Fatal("not equal")
	}
	a.Move(1)
	if a.Location != Pt(2, 0) {
		t.Fatal("not equal")
	}
	a.Turn(1)
	if a.Orientation != OrientationDown {
		t.Fatal("not equal")
	}
	a.Move(1)
	if a.Location != Pt(2, 1) {
		t.Fatal("not equal")
	}
	a.Turn(1)
	if a.Orientation != OrientationLeft {
		t.Fatal("not equal")
	}
	a.Move(1)
	if a.Location != Pt(1, 1) {
		t.Fatal("not equal")
	}
}

var _ Rule = RuleFunc(nil)

func TestRuleFunc(t *testing.T) {
	called := false
	r := RuleFunc(func(a *Ant) {
		called = true
	})
	r.Apply(&Ant{})
	if !called {
		t.Fatal("not called")
	}
}

func TestRuleTurnRightMove(t *testing.T) {
	a := &Ant{
		Location:    Pt(0, 0),
		Orientation: OrientationUp,
	}
	RuleTurnRightMove.Apply(a)
	if a.Location != Pt(1, 0) {
		t.Fatal("not equal")
	}
	if a.Orientation != OrientationRight {
		t.Fatal("not equal")
	}
}

func TestRuleTurnLeftMove(t *testing.T) {
	a := &Ant{
		Location:    Pt(0, 0),
		Orientation: OrientationUp,
	}
	RuleTurnLeftMove.Apply(a)
	if a.Location != Pt(-1, 0) {
		t.Fatal("not equal")
	}
	if a.Orientation != OrientationLeft {
		t.Fatal("not equal")
	}
}

func TestGameStep(t *testing.T) {
	g := &Game{
		Rules: RulesBasic,
		Grid:  NewGrid(Pt(50, 50), 2),
		Ants: []*Ant{
			&Ant{
				Location:    Pt(25, 25),
				Orientation: OrientationUp,
			},
		},
	}
	for i := 0; i < 100000; i++ {
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
		if res != tc.expected {
			t.Errorf("%#v: got %d, want %d", tc, res, tc.expected)
		}
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
		if res != tc.expected {
			t.Errorf("%#v: got %#v, want %#v", tc, res, tc.expected)
		}
	}
}
