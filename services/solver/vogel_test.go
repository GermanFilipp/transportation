package solver

import (
	"reflect"
	"testing"
)

func TestVogelApproximation(t *testing.T) {
	for _, test := range vogelTestCases {
		input := test.input
		actual := NewTransportTable(input.supply, input.demand, input.costs)
		actual.VogelApproximation()
		if !reflect.DeepEqual(actual.matrix, test.expected.matrix) {
			t.Errorf("TestVogelApproximation test expected [%+v],\n actual [%+v]", test.expected.matrix, actual.matrix)
		}
	}
}

func BenchmarkSolverVogel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range vogelTestCases {
			actual := NewTransportTable(test.input.supply, test.input.demand, test.input.costs)
			actual.VogelApproximation()
		}
	}
}
