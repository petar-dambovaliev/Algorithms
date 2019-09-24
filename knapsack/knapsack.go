package knapsack

import (
	"fmt"
	"strings"
)

// Pick picks the best combo from items
// to fill the given knapsack cap as much as possible
// returns weight of picked items and a collection of the individual items
func Pick(items Items, icap int) (chosen Knapsack) {
	icap += 1
	grid := matrix(len(items.items)+1, icap)

	for row := 1; row < len(grid); row++ {
		upperRow := row - 1
		i := items.items[upperRow]
		weight := i.Weight()
		value := i.Value()

		for col := 1; col < len(grid[row]); col++ {
			//set the default to be the cell just above
			// just in case the current item doesn't fit into the capacity
			grid[row][col] = grid[upperRow][col]

			// check if we can fit the current item
			if weight <= col {
				// the new value is the current item
				// plus the last cell that had the best result
				v := value + grid[upperRow][col-weight]
				// check not to overwrite the cell with a lesser value
				if v > grid[row][col] {
					grid[row][col] = v
				}
			}
		}
	}

	// backtrack to get the picked items
	// that madeup the total weight
	col := icap - 1
	for row := len(grid) - 1; row > 1; row-- {
		if grid[row][col] > grid[row-1][col] {
			chosen.Add(items.items[row-1])
			col -= items.items[row-1].Weight()
		}
	}

	return chosen
}

type Knapsack struct {
	items  Items
	weight int
	value  int
}

func (i *Knapsack) Add(it Item) {
	i.items.Add(it)
	i.weight += it.Weight()
	i.value += it.Value()
}

func (i *Knapsack) Weight() int   { return i.weight }
func (i *Knapsack) Value() int    { return i.value }
func (i Knapsack) String() string { return i.items.String() }

// Item is a single entity in a knapsack
type Item interface {
	Value() int
	Weight() int
}

// Items is a collection of Item interface
// to be filled into a knapsack
type Items struct {
	items []Item
}

func (i *Items) Len() int                     { return len(i.items) }
func (i *Items) Add(it Item)                  { i.items = append(i.items, it) }
func (i *Items) AddFromInt(value, weight int) { i.Add(&item{value: value, weight: weight}) }

func (i Items) String() string {
	var s strings.Builder
	for k, it := range i.items {
		s.WriteString(
			fmt.Sprintf("%v => value(%v) weight(%v)\n", k, it.Value(), it.Weight()),
		)
	}
	return s.String()
}

func matrix(ilen, icap int) [][]int {
	matrix := make([][]int, ilen)

	for k := range matrix {
		matrix[k] = make([]int, icap)
	}
	return matrix
}

type item struct {
	value  int
	weight int
}

func (i *item) Value() int  { return i.value }
func (i *item) Weight() int { return i.weight }
