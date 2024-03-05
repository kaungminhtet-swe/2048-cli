package action

func leftAction(input [4][4]int) [4][4]int {
	var output [4][4]int
	for rowId, row := range input {
		outputColId := 1
		for _, cell := range row {
			if cell != EMPTYCELL {
				if output[rowId][outputColId-1] == EMPTYCELL {
					output[rowId][outputColId-1] = cell
				} else if output[rowId][outputColId-1] == cell {
					output[rowId][outputColId-1] += cell
					outputColId++
				} else {
					output[rowId][outputColId] = cell
					outputColId++
				}
			}
		}
	}
	return output
}
