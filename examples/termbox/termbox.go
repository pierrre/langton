package main

import (
	termbox "github.com/nsf/termbox-go"
	"github.com/pierrre/langton"
)

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	evQueue := make(chan termbox.Event)
	go func() {
		for {
			evQueue <- termbox.PollEvent()
		}
	}()

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
			switch ev.Type {
			case termbox.EventKey:
				return
			}
		default:
		}

		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
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
		err = termbox.Flush()
		if err != nil {
			panic(err)
		}

		game.Step()
	}
}
