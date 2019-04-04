package solver

import (
	"reflect"
	"testing"
)

func TestResultToJSON(t *testing.T) {
	for _, test := range jsonTests {
		actual := NewTransportTable(test.input.supply, test.input.demand, test.input.costs)
		actual.VogelApproximation()
		actual.SolveByPotentialsMethod()
		result := actual.ResultToJSON()
		if !reflect.DeepEqual(string(result), test.expected) {
			t.Errorf("TestResultToJSON test expected [%+v], \n actual [%+v]", string(result), test.expected)
		}
	}
}
