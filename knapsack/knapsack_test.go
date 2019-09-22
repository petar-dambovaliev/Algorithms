package knapsack

import "testing"

func TestPick(t *testing.T) {
	makeNapsack := func(ints ...int) (k Knapsack) {
		for i := 0; i < len(ints)-1; i += 2 {
			k.AddFromInt(ints[i], ints[i+1])
		}
		return k
	}

	input := makeNapsack(2, 3, 2, 1, 4, 3, 5, 4, 3, 2)
	output := makeNapsack(3, 2, 5, 4, 2, 1)

	input1 := makeNapsack(1, 3, 4, 3, 8, 5, 5, 6)
	output1 := makeNapsack(8, 5, 4, 3)

	tests := map[int][2]Knapsack{
		7: {
			input, output,
		},
		10: {
			input1, output1,
		},
	}

	weights := []int{10, 12}
	lenths := []int{3, 2}
	var i int

	for icap, v := range tests {
		in := v
		chosen := Pick(in[0], icap)

		if chosen.weight != weights[i] {
			t.Errorf("weigth=%v should be %v", chosen.weight, weights[i])
		}
		if chosen.Len() != lenths[i] {
			t.Fatalf("chosen items lenth=%v but should be %v", chosen.Len(), lenths[i])
		}

		for k, vv := range chosen.items {
			if vv.Value() != in[1].items[k].Value() {
				t.Errorf("Index:%v value=%v should be %v", k, vv.Value(), in[1].items[k].Value())
			}
			if vv.Weight() != in[1].items[k].Weight() {
				t.Errorf("Index:%v weigth=%v should be %v", k, vv.Weight(), in[1].items[k].Weight())
			}
		}
		i++
	}
}
