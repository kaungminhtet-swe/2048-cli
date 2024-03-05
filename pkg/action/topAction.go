package action

func topAction(input [4][4]int) [4][4]int {
	var output [4][4]int

	for colId := 0; colId < len(input); colId++ {
		outputRowId := 1
		for _, row := range input {
			cell := row[colId]
			if cell != EMPTYCELL {
				if output[outputRowId-1][colId] == EMPTYCELL {
					output[outputRowId-1][colId] = cell
				} else if output[outputRowId-1][colId] == cell {
					output[outputRowId-1][colId] += cell
					outputRowId++
				} else {
					output[outputRowId][colId] = cell
					outputRowId++
				}
			}
		}
	}
	return output
}
