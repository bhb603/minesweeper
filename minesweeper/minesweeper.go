package minesweeper

//func NewGame(n, m, x int) (string, error) {
//	// validate n,m

//	// empty ""
//	// mine: "*"
//	// adj: "2"
//	grid := make([][]string, n)
//	for _, row := range grid {
//		row := make([]string, m)
//	}

//	// n*m cells x random
//	// [0, 0, 1, 0, 1, 0, 1]
//	randomList := []int{0, 1, 0, 0, 0 1} // m*n long
//	for i := 0; i < n; i++ {
//		for j := 0; j < m; j++ {
//			cellNum := i*n + j
//			if randomList[cellNum] == 1 {
//				grid[i][j] = "*"
//				for _, adj := range getAdjCells(i, j, n, m) {
//					adjCellVal := grid[adj[0]][adj[1]]

//					if adjCellVal == "" {
//						grid[adj[0]][adj[1]] = "1"
//					} else if intVal, err := strconv.Atoi(adjCellVal); err != nil {
//						grid[adj[0]][adj[1]] = strconv.ItoA(intVal++)
//					}
//				}
//			}
//		}
//	}

//	uuid := db.Save(grid)

//	return uuid, nil
//}

//func getAdjCells(i, j, n, m int) [][]int {
//	adj := [][]int{}

//	deltas := [][]int{
//		{-1, -1},
//		{-1, 0},
//		{-1, 1},
//		{0, -1},
//		{0, 1},
//		{1, -1}
//		{1, 0}
//		{1, 1}
//	}

//	for _, delta := range deltas {
//		di, dj := delta[0], delta[1]
//		newI := i + di
//		newJ := j + dj
//		if newI >= 0 && newI <n && newJ>=0 && newJ < m {
//			adj = append(adj, []int{newI, newj}
//		}
//	}

//	return adj
//}

//type Cell struct {
//	I     int
//	J     int
//	Value string
//}

//type MoveResponse struct {
//	CurrentCell   *Cell
//	AdjBlankCells []*Cell
//}

//func Move(uuid string, i, j int) (*MoveResponse, error) {
//	grid := db.Find(uuid)
//	if grid == nil {
//		return grid, errors.New("not found")
//	}

//	// validate i, j input

//	cellVal := grid[i][j]
//	if cellVal == "*" {
//		// gameover store state indicatating game over

//		return &MoveResponse{
//			CurrentCell: {I: i, J: j, Value: cellVal}
//		}
//	}
//	if cellVal == "" {
//		adjBlankCells = []*Cell{}
//		// getadjcells ->
//		// - if numbered add to adjcells
//		// - if blank, add to queue of cells to check
//		//
//	}

//	return &MoveResponse{
//		CurrentCell: {I: i, J: j, Value: cellVal}
//	}
//}
