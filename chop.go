package main

import (
	"errors"
	"fmt"
	"slices"
)

const (
	dataLimit = 10000
)

func main() {
	data := make([]uint32, 0, 5)
	data = append(data, 0, 7, 53, 9, 1)
	result, err := chop(9, data)
	if err != nil {
		fmt.Printf(fmt.Errorf("chop error:%w", err).Error())
	}

	fmt.Printf("the result is in index: %d", result)

}

func chop(obj uint32, data []uint32) (int, error) {

	if len(data) > dataLimit {
		err := errors.New("the data lenght is superior to data limit")
		return 0, fmt.Errorf(err.Error()+"%d>%d", len(data), dataLimit)
	}

	slices.Sort(data)

	for range data {
		index, found := slices.BinarySearch(data, obj)
		if found {
			fmt.Printf("element %d found at index %d\n", obj, index)

			return index, nil
		}

	}

	return 0, errors.New("number not found")

}
