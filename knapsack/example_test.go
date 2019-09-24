package knapsack

import "fmt"

type mockItem int

func (i mockItem) Value() int  { return int(i) }
func (i mockItem) Weight() int { return int(i) }

func ExampleItems_Add() {
	var i Items
	i.Add(mockItem(1))
	fmt.Println(i)
	// Output: 0 => value(1) weight(1)
}

func ExampleItems_AddFromInt() {
	var i Items
	i.AddFromInt(1, 3)
	fmt.Println(i)
	// Output: 0 => value(1) weight(3)
}

func ExamplePick() {
	var items Items

	items.AddFromInt(2, 3)
	items.AddFromInt(2, 1)
	items.AddFromInt(4, 3)
	items.AddFromInt(5, 4)
	items.AddFromInt(3, 2)

	chosen := Pick(items, 7)
	fmt.Println(chosen)
	// Output: 0 => value(3) weight(2)
	//1 => value(5) weight(4)
	//2 => value(2) weight(1)
}
