package action

func rightAction(input [4][4]int) [4][4]int {
	var output [4][4]int
	for rowId := 0; rowId < len(input); rowId++ {
		for colId, outputColId := len(input[rowId])-1, 2; colId >= 0; colId-- {
			cell := input[rowId][colId]
			if cell != EMPTYCELL {
				if output[rowId][outputColId+1] == EMPTYCELL {
					output[rowId][outputColId+1] = cell
				} else if output[rowId][outputColId+1] == cell {
					output[rowId][outputColId+1] += cell
					outputColId--
				} else {
					output[rowId][outputColId] = cell
					outputColId--
				}
			}
		}
	}
	return output
}
