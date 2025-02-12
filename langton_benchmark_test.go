package langton

import (
	"testing"
)

func BenchmarkGridGetInc(b *testing.B) {
	g := NewGrid(Pt(10, 10), 2)
	for b.Loop() {
		g.GetInc(Pt(5, 5))
	}
}

func BenchmarkOrientationRotateRight1(b *testing.B) {
	o := OrientationUp
	for b.Loop() {
		o = o.Rotate(1)
	}
}

func BenchmarkOrientationRotateRight5(b *testing.B) {
	o := OrientationUp
	for b.Loop() {
		o = o.Rotate(5)
	}
}

func BenchmarkOrientationRotateLeft1(b *testing.B) {
	o := OrientationUp
	for b.Loop() {
		o = o.Rotate(-1)
	}
}

func BenchmarkOrientationRotateLeft5(b *testing.B) {
	o := OrientationUp
	for b.Loop() {
		o = o.Rotate(-5)
	}
}

func BenchmarkAntMoveUp(b *testing.B) {
	a := &Ant{Orientation: OrientationUp}
	for b.Loop() {
		a.Move(1)
	}
}

func BenchmarkAntMoveRight(b *testing.B) {
	a := &Ant{Orientation: OrientationRight}
	for b.Loop() {
		a.Move(1)
	}
}

func BenchmarkAntMoveDown(b *testing.B) {
	a := &Ant{Orientation: OrientationDown}
	for b.Loop() {
		a.Move(1)
	}
}

func BenchmarkAntMoveLeft(b *testing.B) {
	a := &Ant{Orientation: OrientationLeft}
	for b.Loop() {
		a.Move(1)
	}
}

func BenchmarkGameStep(b *testing.B) {
	g := &Game{
		Rules: RulesBasic,
		Grid:  NewGrid(Pt(20, 20), 2),
		Ants: []*Ant{{
			Location:    Pt(10, 10),
			Orientation: OrientationUp,
		}},
	}
	for b.Loop() {
		g.Step()
	}
}
