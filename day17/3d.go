package main

import "fmt"

// Vector3 xyz
type Vector3 = [3]int

// Input3d data structure
type Input3d struct {
	cells map[string]Vector3
}

func (i *Input3d) aliveNeighbours(v3 Vector3) int {
	sum := 0
	for z := -1; z < 2; z++ {
		for y := -1; y < 2; y++ {
			for x := -1; x < 2; x++ {
				if x == 0 && y == 0 && z == 0 {
					continue
				}
				if i.getCellState(Vector3{v3[0] + x, v3[1] + y, v3[2] + z}) {
					sum++
				}
			}
		}
	}
	return sum
}

func (i *Input3d) cycle() Input3d {
	newI := Input3d{
		cells: map[string]Vector3{},
	}

	minBound, maxBound := i.getBounds()

	for z := minBound[2] - 1; z < maxBound[2]+2; z++ {
		for y := minBound[1] - 1; y < maxBound[1]+2; y++ {
			for x := minBound[0] - 1; x < maxBound[0]+2; x++ {
				cell := Vector3{x, y, z}
				alive := i.getCellState(cell)
				neighbours := i.aliveNeighbours(cell)
				newI.setCellState(cell, nextCellState(alive, neighbours))
			}
		}
	}

	return newI
}

func (i *Input3d) totalAlive() int {
	return len(i.cells)
}

func (i *Input3d) getBounds() (Vector3, Vector3) {
	min := Vector3{}
	max := Vector3{}

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
		if v[0] > max[0] {
			max[0] = v[0]
		}
		if v[1] > max[1] {
			max[1] = v[1]
		}
		if v[2] > max[2] {
			max[2] = v[2]
		}
	}

	return min, max
}

func parseInput3d(input []string) Input3d {
	i := Input3d{
		cells: map[string]Vector3{},
	}

	for y, row := range input {
		for x, cell := range row {
			if string(cell) == "#" {
				i.setCellState(Vector3{x, y, 0}, true)
			}
		}
	}

	return i
}

func (i *Input3d) getCellState(v3 Vector3) bool {
	_, ok := i.cells[hashVector3(v3)]
	return ok
}

func (i *Input3d) setCellState(v3 Vector3, alive bool) {
	hash := hashVector3(v3)
	if alive {
		i.cells[hash] = v3
	} else {
		delete(i.cells, hash)
	}
}

func hashVector3(v3 Vector3) string {
	return fmt.Sprint(v3[0]) + "," + fmt.Sprint(v3[1]) + "," + fmt.Sprint(v3[2])
}
