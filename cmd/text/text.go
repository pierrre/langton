// Package text provides a text implementation of Langton's ant.
package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"

	"github.com/pierrre/langton"
)

func main() {
	game := &langton.Game{
		Rules: langton.RulesBasic,
		Grid:  langton.NewGrid(langton.Pt(80, 60), 2),
		Ants: []*langton.Ant{
			{
				Location:    langton.Pt(40, 30),
				Orientation: langton.OrientationUp,
			},
		},
	}
	buf := new(bytes.Buffer)
	for step := 0; ; step++ {
		buf.Reset()
		for y := range game.Grid.Size.Y {
			for x := range game.Grid.Size.X {
				p := langton.Pt(x, y)
				s := strconv.Itoa(int(game.Grid.Get(p)))
				for _, a := range game.Ants {
					if a.Location == p {
						s = "*"
						break
					}
				}
				_, _ = buf.WriteString(s)
			}
			_, _ = buf.WriteString("\n")
		}
		fmt.Fprintln(os.Stdout, buf)
		game.Step()
	}
}
