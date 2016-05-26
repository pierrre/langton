package main

import (
	"bytes"
	"fmt"
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
		if step%1 == 0 {
			buf.Reset()
			for y := 0; y < game.Grid.Size.Y; y++ {
				for x := 0; x < game.Grid.Size.X; x++ {
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
			fmt.Println(buf)
		}
		game.Step()
	}
}
