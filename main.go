package main

import (
	"fmt"
	"sort"
)

type Animals struct {
	Species string
	Age     int
}

func (a Animals) String() string {
	return fmt.Sprintf("%s: %d", a.Species, a.Age)
}

type SortByAge []Animals

func (a SortByAge) Len() int           { return len(a) }
func (a SortByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

// func (a Animals) Less(i, j int) bool {return a[i].Species < a[j] }
func main() {

	animalGroup := []Animals{
		{"Kangaroo", 3},
		{"Wombat", 5},
		{"Koala", 1},
		{"Tasmanian devil", 7},
	}

	fmt.Println(animalGroup)
	fmt.Println(animalGroup[1])
	sort.Sort(SortByAge(animalGroup))
	fmt.Println(animalGroup)

}
