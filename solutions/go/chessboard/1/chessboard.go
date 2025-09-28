// Package chessboard provides types and functions for counting
// squares and pieces on a chessboard.
package chessboard

// File represents a vertical column of squares on a chessboard.
type File []bool

// Chessboard represents the whole chessboard, mapping file names to Files.
type Chessboard map[string]File

// CountInFile returns the number of occupied squares in the given file.
func CountInFile(cb Chessboard, file string) int {
	sum := 0
	for _, occupied := range cb[file] {
		if occupied {
			sum++
		}
	}
	return sum
}

// CountInRank returns the number of occupied squares in the given rank (1–8).
// Returns 0 if rank is out of range.
func CountInRank(cb Chessboard, rank int) int {
	if rank < 1 || rank > 8 {
		return 0
	}
	sum := 0
	for _, f := range cb {
		if f[rank-1] {
			sum++
		}
	}
	return sum
}

// CountAll returns the total number of squares on the chessboard.
func CountAll(cb Chessboard) int {
	sum := 0
	for _, f := range cb {
		sum += len(f)
	}
	return sum
}

// CountOccupied returns the total number of occupied squares on the chessboard.
func CountOccupied(cb Chessboard) int {
	sum := 0
	for _, f := range cb {
		for _, occupied := range f {
			if occupied {
				sum++
			}
		}
	}
	return sum
}
