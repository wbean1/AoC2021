package day17

import "fmt"

type Coord struct {
	x, y int
}

var targetTopLeft, targetBottomRight = Coord{25, -200}, Coord{67, -260}

func Run() {
	maxY := 0
	for xVelocity := 1; xVelocity < 68; xVelocity++ {
		for yVelocity := 1; yVelocity < 1000; yVelocity++ {
			coords := simulateShot(xVelocity, yVelocity)
			maxYinPath, hitsTarget := doesShotHitTarget(coords)
			if hitsTarget {
				if maxYinPath > maxY {
					maxY = maxYinPath
				}
			}
		}
	}
	fmt.Printf("part1: highest Y while still hitting target is: %d\n", maxY)
	successfulVelocities := []Coord{}
	for xVelocity := 1; xVelocity < 68; xVelocity++ {
		for yVelocity := -1000; yVelocity < 1000; yVelocity++ {
			coords := simulateShot(xVelocity, yVelocity)
			_, hitsTarget := doesShotHitTarget(coords)
			if hitsTarget {
				successfulVelocities = append(successfulVelocities, Coord{xVelocity, yVelocity})
			}
		}
	}
	fmt.Printf("part2: number of successful initial velocities: %d\n", len(successfulVelocities))
}

func doesShotHitTarget(c []Coord) (int, bool) {
	maxY := 0
	for _, coord := range c {
		if coord.y > maxY {
			maxY = coord.y
		}
		if coord.x >= targetTopLeft.x &&
			coord.x <= targetBottomRight.x &&
			coord.y <= targetTopLeft.y &&
			coord.y >= targetBottomRight.y {
			return maxY, true
		}
	}
	return maxY, false
}

func simulateShot(xVel, yVel int) []Coord {
	currentPosition := Coord{0, 0}
	positions := []Coord{currentPosition}
	for {
		currentPosition = Coord{
			currentPosition.x + xVel,
			currentPosition.y + yVel,
		}
		positions = append(positions, currentPosition)
		if xVel > 0 {
			xVel -= 1
		} else if xVel < 0 {
			xVel += 1
		}
		yVel -= 1
		if currentPosition.x > targetBottomRight.x {
			break
		}
		if currentPosition.y < targetBottomRight.y {
			break
		}
	}
	return positions
}
