package board

import (
	"io"
	"math/rand"
	"os"
)

const STARTVALUE = 2

type Render func(io.Writer, [4][4]int) error
type Move interface {
	Top([4][4]int) [4][4]int
	Right([4][4]int) [4][4]int
	Bottom([4][4]int) [4][4]int
	Left([4][4]int) [4][4]int
}

type Board struct {
	move   Move
	state  [4][4]int
	render Render
}

func (b *Board) MoveTop() {
	b.state = b.move.Top(b.state)
}

func (b *Board) MoveRight() {
	b.state = b.move.Right(b.state)
}

func (b *Board) MoveBottom() {
	b.state = b.move.Bottom(b.state)
}

func (b *Board) MoveLeft() {
	b.state = b.move.Left(b.state)
}

func (b *Board) SetAction(moveAction Move) {
	b.move = moveAction
}

func (b *Board) newCell() {
	for {
		row, col := rand.Intn(4), rand.Intn(4)
		if b.state[row][col] == 0 {
			b.state[row][col] = STARTVALUE
			break
		}
	}
}

func (b *Board) Start() {
	b.newCell()
	b.Render()
}

func (b *Board) SetRender(render Render) {
	b.render = render
}

func (b *Board) Render() {
	b.newCell()
	b.render(os.Stdout, b.state)
}

func New() *Board {
	return &Board{}
}
