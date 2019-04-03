package solver

import (
	"reflect"
	"testing"
)

func TestSolverNewTransporTable(t *testing.T) {
	for _, test := range solverTestCases {
		input := test.input
		actual := NewTransporTable(input.supply, input.demand, input.costs)
		if !reflect.DeepEqual(actual.supply, test.expected.supply) {
			t.Errorf("NewTransporTable test [%+v], expected [%+v], actual [%+v]", test.input.supply, test.expected.supply, actual.supply)
		}
		if !reflect.DeepEqual(actual.demand, test.expected.demand) {
			t.Errorf("NewTransporTable test [%+v], expected [%+v], actual [%+v]", test.input.demand, test.expected.demand, actual.demand)
		}
		if !reflect.DeepEqual(actual.matrix, test.expected.matrix) {
			t.Errorf("NewTransporTable test [%+v], expected [%+v], actual [%+v]", test.input, test.expected.matrix, actual.matrix)
		}
		if !reflect.DeepEqual(actual.costs, test.expected.costs) {
			t.Errorf("NewTransporTable test [%+v], expected [%+v], actual [%+v]", test.input.costs, test.expected.costs, actual.costs)
		}
	}
}

func BenchmarkSolverNewTransportationTable(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range solverTestCases {
			NewTransporTable(test.input.supply, test.input.demand, test.input.costs)
		}
	}
}

// moore tests
func TestSolveByPotentialsMethod(t *testing.T) {
	for _, test := range potentialTestCases {
		actual := NewTransporTable(test.input.supply, test.input.demand, test.input.costs)
		actual.VogelApproximation()
		actual.SolveByPotentialsMethod()
		if !reflect.DeepEqual(actual.matrix, test.expected.matrix) {
			t.Errorf("SolveByPotentialsMethod test [%+v], \n expected [%+v], \n actual [%+v]", test.input, test.expected.matrix, actual.matrix)
		}
	}
}

// func BenchmarkSolveByPotentialsMethod(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		for _, test := range potentialTestCases {
// 			actual := NewTransporTable(test.input.supply, test.input.demand, test.input.costs)
// 			actual.VogelApproximation()
// 			actual.SolveByPotentialsMethod()
// 		}
// 	}
// }
