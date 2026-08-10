# Langton

Langton's ant automaton library.

[![Go Reference](https://pkg.go.dev/badge/github.com/pierrre/langton.svg)](https://pkg.go.dev/github.com/pierrre/langton)

## Features

- Grid with configurable number of states
- Ants with orientation (up, right, down, left)
- Customizable rules
- Multiple ants per game
- Demo commands: `cmd/termbox` (interactive terminal) and `cmd/text` (text output)

## Usage

```bash
# Local build
make build
./build/termbox
./build/text

# Remote install
go install github.com/pierrre/langton/cmd/termbox@latest
go install github.com/pierrre/langton/cmd/text@latest

# Module install
go get github.com/pierrre/langton@latest
```
