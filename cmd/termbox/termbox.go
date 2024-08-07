// Package termbox provides a termbox implementation of Langton's ant.
package main

import (
	"context"

	termbox "github.com/nsf/termbox-go"
	"github.com/pierrre/go-libs/goroutine"
	"github.com/pierrre/langton"
)

func main() {
	ctx := context.Background()
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	evQueue := make(chan termbox.Event)
	goroutine.Start(ctx, func(context.Context) {
		for {
			evQueue <- termbox.PollEvent()
		}
	})

	width, height := termbox.Size()
	game := &langton.Game{
		Rules: langton.RulesBasic,
		Grid:  langton.NewGrid(langton.Pt(width, height), 2),
		Ants: []*langton.Ant{
			{
				Location:    langton.Pt(width/2, height/2),
				Orientation: langton.OrientationUp,
			},
		},
	}

	for step := 0; ; step++ {
		select {
		case ev := <-evQueue:
			if ev.Type == termbox.EventKey {
				return
			}
		default:
		}

		draw(game)
		err = termbox.Flush()
		if err != nil {
			panic(err)
		}

		game.Step()
	}
}

func draw(game *langton.Game) {
	for y := range game.Grid.Size.Y {
		for x := range game.Grid.Size.X {
			p := langton.Pt(x, y)
			var bg termbox.Attribute
			switch game.Grid.Get(p) {
			case 0:
				bg = termbox.ColorBlack
			default:
				bg = termbox.ColorWhite
			}
			for _, a := range game.Ants {
				if a.Location == p {
					bg = termbox.ColorRed
					continue
				}
			}
			termbox.SetCell(x, y, ' ', termbox.ColorDefault, bg)
		}
	}
}
