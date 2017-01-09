package queenattack

import (
	"fmt"
	"math"
	"strconv"
)

const testVersion = 1

type coord [2]int

// CanQueenAttack check if one queen could attack the second
func CanQueenAttack(first, second string) (bool, error) {
	if len(first) < 2 || len(second) < 2 {
		return false, fmt.Errorf("Invalid coordinates")
	}

	firstCoords, err := convertCoordinates(first)
	if err != nil {
		return false, fmt.Errorf("Invalid coordinates")
	}
	secondCoords, err := convertCoordinates(second)
	if err != nil {
		return false, fmt.Errorf("Invalid coordinates")
	}

	if firstCoords[0] == secondCoords[0] && firstCoords[1] == secondCoords[1] {
		return false, fmt.Errorf("Same coordinates")
	}

	if firstCoords[1] > 7 || secondCoords[1] > 7 {
		return false, fmt.Errorf("Off board")
	}

	if math.Abs(float64(firstCoords[0]-secondCoords[0])) == math.Abs(float64(firstCoords[1]-secondCoords[1])) ||
		firstCoords[0] == secondCoords[0] ||
		firstCoords[1] == secondCoords[1] {
		return true, nil
	}

	return false, nil
}

func convertCoordinates(str string) (coord, error) {
	first := int(str[0]) - 97
	second, err := strconv.Atoi(string(str[1]))
	if err != nil {
		return [2]int{0, 0}, err
	}
	second--
	return [2]int{first, second}, nil
}
