package solver

import (
	"reflect"
	"testing"
)

func TestVogelApproximation(t *testing.T) {
	for _, test := range vogelTestCases {
		input := test.input
		actual := NewTransporTable(input.supply, input.demand, input.costs)
		actual.VogelApproximation()
		// if !reflect.DeepEqual(actual.supply, test.expected.supply) {
		// 	t.Errorf("Encrypt test [%+v], expected [%+v], actual [%+v]", test.input.supply, test.expected.supply, actual.supply)
		// }
		// if !reflect.DeepEqual(actual.demand, test.expected.demand) {
		// 	t.Errorf("Encrypt test [%+v], expected [%+v], actual [%+v]", test.input.demand, test.expected.demand, actual.demand)
		// }
		if !reflect.DeepEqual(actual.matrix, test.expected.matrix) {
			t.Errorf("SolverVogelApproximation test expected [%+v],\n actual [%+v]", test.expected.matrix, actual.matrix)
		}
		// if !reflect.DeepEqual(actual.costs, test.expected.costs) {
		// 	t.Errorf("Encrypt test [%+v], expected [%+v], actual [%+v]", test.input.costs, test.expected.costs, actual.costs)
		// }
	}
}

func BenchmarkSolverVogel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range vogelTestCases {
			actual := NewTransporTable(test.input.supply, test.input.demand, test.input.costs)
			actual.VogelApproximation()
		}
	}
}
