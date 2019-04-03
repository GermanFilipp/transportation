package solver

type solverInput struct {
	demand []int
	supply []int
	costs  [][]float64
}
type solverTest struct {
	input    solverInput
	expected TransportTable
}
type jsonResponseTest struct {
	input    solverInput
	expected string
}

var solverTestCases = []solverTest{
	{
		input: solverInput{
			demand: []int{20, 40, 30, 10, 50, 25},
			supply: []int{30, 50, 75, 20},
			costs: [][]float64{
				{1, 2, 1, 4, 5, 2},
				{3, 3, 2, 1, 4, 3},
				{4, 2, 5, 9, 6, 2},
				{3, 1, 7, 3, 4, 6},
			},
		},
		expected: TransportTable{
			supply:    []int{30, 50, 75, 20},
			demand:    []int{20, 40, 30, 10, 50, 25},
			origSup:   []int{30, 50, 75, 20},
			origDem:   []int{20, 40, 30, 10, 50, 25},
			supplyLen: 4,
			demandLen: 6,
			costs: [][]float64{
				{1, 2, 1, 4, 5, 2},
				{3, 3, 2, 1, 4, 3},
				{4, 2, 5, 9, 6, 2},
				{3, 1, 7, 3, 4, 6},
			},
			matrix: [][]shipment{
				make([]shipment, 6),
				make([]shipment, 6),
				make([]shipment, 6),
				make([]shipment, 6),
			},
			rowsDone:    []bool{false, false, false, false},
			columnsDone: []bool{false, false, false, false, false, false},
		},
	},
	{ // CLOSED TASK EXAMPLE
		input: solverInput{
			demand: []int{16, 18, 30, 25},
			supply: []int{19, 37, 34},
			costs: [][]float64{
				{5, 3, 6, 2},
				{4, 7, 9, 1},
				{3, 4, 7, 5},
			},
		},
		expected: TransportTable{
			supply:    []int{19, 37, 34},
			demand:    []int{16, 18, 30, 25, 1},
			origSup:   []int{19, 37, 34},
			origDem:   []int{16, 18, 30, 25, 1},
			supplyLen: 3,
			demandLen: 5,
			costs: [][]float64{
				{5, 3, 6, 2, 0},
				{4, 7, 9, 1, 0},
				{3, 4, 7, 5, 0},
			},
			matrix: [][]shipment{
				make([]shipment, 5),
				make([]shipment, 5),
				make([]shipment, 5),
			},
			rowsDone:    []bool{false, false, false, false},
			columnsDone: []bool{false, false, false, false},
		},
	},
	{ // CLOSED TASK EXAMPLE
		input: solverInput{
			demand: []int{17, 18, 30, 25},
			supply: []int{18, 37, 34},
			costs: [][]float64{
				{5, 3, 6, 2},
				{4, 7, 9, 1},
				{3, 4, 7, 5},
			},
		},
		expected: TransportTable{
			supply:    []int{18, 37, 34, 1},
			demand:    []int{17, 18, 30, 25},
			origSup:   []int{18, 37, 34, 1},
			origDem:   []int{17, 18, 30, 25},
			supplyLen: 4,
			demandLen: 4,
			costs: [][]float64{
				{5, 3, 6, 2},
				{4, 7, 9, 1},
				{3, 4, 7, 5},
				{0, 0, 0, 0},
			},
			matrix: [][]shipment{
				make([]shipment, 4),
				make([]shipment, 4),
				make([]shipment, 4),
				make([]shipment, 4),
			},
			rowsDone:    []bool{false, false, false, false},
			columnsDone: []bool{false, false, false, false},
		},
	},
}

var vogelTestCases = []solverTest{
	{
		input: solverInput{
			demand: []int{20, 40, 30, 10, 50, 25},
			supply: []int{30, 50, 75, 20},
			costs: [][]float64{
				{1, 2, 1, 4, 5, 2},
				{3, 3, 2, 1, 4, 3},
				{4, 2, 5, 9, 6, 2},
				{3, 1, 7, 3, 4, 6},
			},
		},
		expected: TransportTable{
			supply:    []int{30, 50, 75, 20},
			demand:    []int{20, 40, 30, 10, 50, 25},
			origSup:   []int{30, 50, 75, 20},
			origDem:   []int{20, 40, 30, 10, 50, 25},
			supplyLen: 4,
			demandLen: 6,
			costs: [][]float64{
				{1, 2, 1, 4, 5, 2},
				{3, 3, 2, 1, 4, 3},
				{4, 2, 5, 9, 6, 2},
				{3, 1, 7, 3, 4, 6},
			},
			matrix: [][]shipment{
				{
					shipment{quantity: 20, costPerUnit: 1, indexInRow: 0, indexinColumn: 0},
					shipment{quantity: 0, costPerUnit: 0, indexInRow: 0, indexinColumn: 0},
					shipment{quantity: 10, costPerUnit: 1, indexInRow: 0, indexinColumn: 2},
					shipment{quantity: 0, costPerUnit: 0, indexInRow: 0, indexinColumn: 0},
					shipment{quantity: 0, costPerUnit: 0, indexInRow: 0, indexinColumn: 0},
					shipment{quantity: 0, costPerUnit: 0, indexInRow: 0, indexinColumn: 0},
				},
				{
					shipment{quantity: 0, costPerUnit: 0, indexInRow: 0, indexinColumn: 0},
					shipment{quantity: 0, costPerUnit: 0, indexInRow: 0, indexinColumn: 0},
					shipment{quantity: 20, costPerUnit: 2, indexInRow: 1, indexinColumn: 2},
					shipment{quantity: 10, costPerUnit: 1, indexInRow: 1, indexinColumn: 3},
					shipment{quantity: 20, costPerUnit: 4, indexInRow: 1, indexinColumn: 4},
					shipment{quantity: 0, costPerUnit: 0, indexInRow: 0, indexinColumn: 0},
				}, {
					shipment{quantity: 0, costPerUnit: 0, indexInRow: 0, indexinColumn: 0},
					shipment{quantity: 20, costPerUnit: 2, indexInRow: 2, indexinColumn: 1},
					shipment{quantity: 0, costPerUnit: 0, indexInRow: 0, indexinColumn: 0},
					shipment{quantity: 0, costPerUnit: 0, indexInRow: 0, indexinColumn: 0},
					shipment{quantity: 30, costPerUnit: 6, indexInRow: 2, indexinColumn: 4},
					shipment{quantity: 25, costPerUnit: 2, indexInRow: 2, indexinColumn: 5},
				},
				{
					shipment{quantity: 0, costPerUnit: 0, indexInRow: 0, indexinColumn: 0},
					shipment{quantity: 20, costPerUnit: 1, indexInRow: 3, indexinColumn: 1},
					shipment{quantity: 0, costPerUnit: 0, indexInRow: 0, indexinColumn: 0},
					shipment{quantity: 0, costPerUnit: 0, indexInRow: 0, indexinColumn: 0},
					shipment{quantity: 0, costPerUnit: 0, indexInRow: 0, indexinColumn: 0},
					shipment{quantity: 0, costPerUnit: 0, indexInRow: 0, indexinColumn: 0},
				},
			},
			rowsDone:    []bool{true, true, true, true},
			columnsDone: []bool{true, true, true, true, true, true},
		},
	},
	// TODO: Add this cases
	// {
	// 	input: solverInput{
	// 		demand: []int{16, 18, 30, 25},
	// 		supply: []int{19, 37, 34},
	// 		costs: [][]float64{
	// 			{5, 3, 6, 2},
	// 			{4, 7, 9, 1},
	// 			{3, 4, 7, 5},
	// 		},
	// 	},
	// },
	// {
	// 	input: solverInput{
	// 		demand: []int{17, 18, 30, 25},
	// 		supply: []int{18, 37, 34},
	// 		costs: [][]float64{
	// 			{5, 3, 6, 2},
	// 			{4, 7, 9, 1},
	// 			{3, 4, 7, 5},
	// 		},
	// 	},
	// },

}

var potentialTestCases = []solverTest{
	{
		input: solverInput{
			demand: []int{20, 40, 30, 10, 50, 25},
			supply: []int{30, 50, 75, 20},
			costs: [][]float64{
				{1, 2, 1, 4, 5, 2},
				{3, 3, 2, 1, 4, 3},
				{4, 2, 5, 9, 6, 2},
				{3, 1, 7, 3, 4, 6},
			},
		},
		expected: TransportTable{
			supply:    []int{30, 50, 75, 20},
			demand:    []int{20, 40, 30, 10, 50, 25},
			origSup:   []int{30, 50, 75, 20},
			origDem:   []int{20, 40, 30, 10, 50, 25},
			supplyLen: 4,
			demandLen: 6,
			costs: [][]float64{
				{1, 2, 1, 4, 5, 2},
				{3, 3, 2, 1, 4, 3},
				{4, 2, 5, 9, 6, 2},
				{3, 1, 7, 3, 4, 6},
			},
			matrix: [][]shipment{
				{
					shipment{quantity: 20, costPerUnit: 1, indexInRow: 0, indexinColumn: 0},
					shipment{quantity: 0, costPerUnit: 0, indexInRow: 0, indexinColumn: 0},
					shipment{quantity: 10, costPerUnit: 1, indexInRow: 0, indexinColumn: 2},
					shipment{quantity: 0, costPerUnit: 0, indexInRow: 0, indexinColumn: 0},
					shipment{quantity: 0, costPerUnit: 0, indexInRow: 0, indexinColumn: 0},
					shipment{quantity: 0, costPerUnit: 0, indexInRow: 0, indexinColumn: 0},
				},
				{
					shipment{quantity: 0, costPerUnit: 0, indexInRow: 0, indexinColumn: 0},
					shipment{quantity: 0, costPerUnit: 0, indexInRow: 0, indexinColumn: 0},
					shipment{quantity: 20, costPerUnit: 2, indexInRow: 1, indexinColumn: 2},
					shipment{quantity: 10, costPerUnit: 1, indexInRow: 1, indexinColumn: 3},
					shipment{quantity: 20, costPerUnit: 4, indexInRow: 1, indexinColumn: 4},
					shipment{quantity: 0, costPerUnit: 0, indexInRow: 0, indexinColumn: 0},
				},
				{
					shipment{quantity: 0, costPerUnit: 0, indexInRow: 0, indexinColumn: 0},
					shipment{quantity: 40, costPerUnit: 2, indexInRow: 2, indexinColumn: 1},
					shipment{quantity: 0, costPerUnit: 0, indexInRow: 0, indexinColumn: 0},
					shipment{quantity: 0, costPerUnit: 0, indexInRow: 0, indexinColumn: 0},
					shipment{quantity: 10, costPerUnit: 6, indexInRow: 2, indexinColumn: 4},
					shipment{quantity: 25, costPerUnit: 2, indexInRow: 2, indexinColumn: 5},
				},
				{
					shipment{quantity: 0, costPerUnit: 0, indexInRow: 0, indexinColumn: 0},
					shipment{quantity: 0, costPerUnit: 0, indexInRow: 0, indexinColumn: 0},
					shipment{quantity: 0, costPerUnit: 0, indexInRow: 0, indexinColumn: 0},
					shipment{quantity: 0, costPerUnit: 0, indexInRow: 0, indexinColumn: 0},
					shipment{quantity: 20, costPerUnit: 4, indexInRow: 3, indexinColumn: 4},
					shipment{quantity: 0, costPerUnit: 0, indexInRow: 0, indexinColumn: 0},
				},
			},
			rowsDone:    []bool{false, false, false, false},
			columnsDone: []bool{false, false, false, false, false, false},
		},
	},
}

var jsonTests = []jsonResponseTest{
	{
		input: solverInput{
			demand: []int{16, 18, 30, 25},
			supply: []int{19, 37, 34},
			costs: [][]float64{
				{5, 3, 6, 2},
				{4, 7, 9, 1},
				{3, 4, 7, 5},
			},
		},
		expected: `{"totalCost":347,"result":[[0,18,1,0,0],[11,0,0,25,1],[5,0,29,0,0]],"supply":[19,37,34],"demand":[16,18,30,25,1],"costs":[[5,3,6,2,0],[4,7,9,1,0],[3,4,7,5,0]]}`,
	},
}
