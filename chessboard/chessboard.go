package chessboard

type Rank []bool

type Chessboard map[string]Rank

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank
func CountInRank(cb Chessboard, rank string) int {
	count := 0

	for s := range cb {
		if s == rank {
			for _, i := range cb[s] {
				if i {
					count++
				}
			}
		}
	}

	return count
}

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file
func CountInFile(cb Chessboard, file int) int {
	count := 0

	index := file - 1
	if index < 0 || index > 7 {
		return count
	}

	for _, s := range cb {
		if s[index] {
			count++
		}
	}

	return count
}

// CountAll should count how many squares are present in the chessboard
func CountAll(cb Chessboard) int {
	count := 0

	for _, r := range cb {
		for range r {
			count++
		}
	}

	return count
}

// CountOccupied returns how many squares are occupied in the chessboard
func CountOccupied(cb Chessboard) int {
	count := 0

	for _, r := range cb {
		for _, occupied := range r {
			if occupied {
				count++
			}
		}
	}

	return count
}
