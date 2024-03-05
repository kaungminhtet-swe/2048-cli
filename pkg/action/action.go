package action

const (
	EMPTYCELL = iota
	CELL
)

type Action struct{}

func (a *Action) Top(input [4][4]int) [4][4]int {
	return topAction(input)
}

func (a *Action) Right(input [4][4]int) [4][4]int {
	return rightAction(input)
}

func (a *Action) Bottom(input [4][4]int) [4][4]int {
	return bottomAction(input)
}

func (a *Action) Left(input [4][4]int) [4][4]int {
	return leftAction(input)
}

func New() *Action {
	return &Action{}
}
