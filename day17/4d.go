package main

import "fmt"

// Vector4 xyz
type Vector4 = [4]int

// Input4d data structure
type Input4d struct {
	cells map[string]Vector4
}

func (i *Input4d) aliveNeighbours(v4 Vector4) int {
	sum := 0
	for w := -1; w < 2; w++ {
		for z := -1; z < 2; z++ {
			for y := -1; y < 2; y++ {
				for x := -1; x < 2; x++ {
					if x == 0 && y == 0 && z == 0 && w == 0 {
						continue
					}
					if i.getCellState(Vector4{v4[0] + x, v4[1] + y, v4[2] + z, v4[3] + w}) {
						sum++
					}
				}
			}
		}
	}
	return sum
}

func (i *Input4d) cycle() Input4d {
	newI := Input4d{
		cells: map[string]Vector4{},
	}

	minBound, maxBound := i.getBounds()

	for w := minBound[3] - 1; w < maxBound[2]+2; w++ {
		for z := minBound[2] - 1; z < maxBound[2]+2; z++ {
			for y := minBound[1] - 1; y < maxBound[1]+2; y++ {
				for x := minBound[0] - 1; x < maxBound[0]+2; x++ {
					cell := Vector4{x, y, z, w}
					alive := i.getCellState(cell)
					neighbours := i.aliveNeighbours(cell)
					newI.setCellState(cell, nextCellState(alive, neighbours))
				}
			}
		}
	}

	return newI
}

func (i *Input4d) totalAlive() int {
	return len(i.cells)
}

func (i *Input4d) getBounds() (Vector4, Vector4) {
	min := Vector4{}
	max := Vector4{}

	for _, v := range i.cells {
		if v[0] < min[0] {
			min[0] = v[0]
		}
		if v[1] < min[1] {
			min[1] = v[1]
		}
		if v[2] < min[2] {
			min[2] = v[2]
		}
		if v[3] < min[3] {
			min[3] = v[3]
		}
		if v[0] > max[0] {
			max[0] = v[0]
		}
		if v[1] > max[1] {
			max[1] = v[1]
		}
		if v[2] > max[2] {
			max[2] = v[2]
		}
		if v[3] > max[3] {
			max[3] = v[3]
		}
	}

	return min, max
}

func parseInput4d(input []string) Input4d {
	i := Input4d{
		cells: map[string]Vector4{},
	}

	for y, row := range input {
		for x, cell := range row {
			if string(cell) == "#" {
				i.setCellState(Vector4{x, y, 0, 0}, true)
			}
		}
	}

	return i
}

func (i *Input4d) getCellState(v4 Vector4) bool {
	_, ok := i.cells[hashVector4(v4)]
	return ok
}

func (i *Input4d) setCellState(v4 Vector4, alive bool) {
	hash := hashVector4(v4)
	if alive {
		i.cells[hash] = v4
	} else {
		delete(i.cells, hash)
	}
}

func hashVector4(v4 Vector4) string {
	return fmt.Sprint(v4[0]) + "," + fmt.Sprint(v4[1]) + "," + fmt.Sprint(v4[2]) + "," + fmt.Sprint(v4[3])
}
