package problem

// import (
// 	"testing"

// 	"github.com/germanfilipp/transportation/controllers/transportation"
// 	// "github.com/germanfilipp/transportation/controllers/transportation"
// )

// func TestIsValidParams(t *testing.T) {
// 	params := transportation.ProblemParams{
// 		Demand: []int{20, 40, 30, 10, 50, 25},
// 		Supply: []int{30, 50, 75, 20},
// 		Costs: [][]float64{
// 			{1, 2, 1, 4, 5, 2},
// 			{3, 3, 2, 1, 4, 3},
// 			{4, 2, 5, 9, 6, 2},
// 			{3, 1, 7, 3, 4, 6},
// 		},
// 	}
// 	if _, isVal := IsValidParams(params); !isVal {
// 		t.Errorf("TestIsValidParams test expected [%+v], actual [%+v]", false, isVal)
// 	}
// }

// func TestIsInvalidParamsCostsDemand(t *testing.T) {
// 	params := transportation.ProblemParams{
// 		Demand: []int{20, 40, 30, 10},
// 		Supply: []int{30, 50, 75, 20},
// 		Costs: [][]float64{
// 			{1, 2, 1, 4, 5, 2},
// 			{3, 3, 2, 1, 4, 3},
// 			{4, 2, 5, 9, 6, 2},
// 			{3, 1, 7, 3, 4, 6},
// 		},
// 	}
// 	if _, isVal := IsValidParams(params); isVal {
// 		t.Errorf("TestIsValidParams test expected [%+v], actual [%+v]", false, isVal)
// 	}
// 	m := map[string]string{"costs": "length of costs columns should be equal to demand"}
// 	if err, _ := IsValidParams(params); m["costs"] != err.FieldName["costs"] {
// 		t.Errorf("TestIsValidParams test expected [%+v], actual [%+v]", m, err.FieldName)
// 	}
// }

// func TestIsInvalidParamsCostsSupply(t *testing.T) {
// 	params := transportation.ProblemParams{
// 		Demand: []int{20, 40, 30, 10, 50, 25},
// 		Supply: []int{30, 50, 75},
// 		Costs: [][]float64{
// 			{1, 2, 1, 4, 5, 2},
// 			{3, 3, 2, 1, 4, 3},
// 			{4, 2, 5, 9, 6, 2},
// 			{3, 1, 7, 3, 4, 6},
// 		},
// 	}
// 	if _, isVal := IsValidParams(params); isVal {
// 		t.Errorf("TestIsValidParams test expected [%+v], actual [%+v]", false, isVal)
// 	}
// 	m := map[string]string{"costs": "length of costs rows should be equal to supply"}
// 	if err, _ := IsValidParams(params); m["costs"] != err.FieldName["costs"] {
// 		t.Errorf("TestIsValidParams test expected [%+v], actual [%+v]", m, err.FieldName)
// 	}
// }

// func TestIsInvalidSupply(t *testing.T) {
// 	params := transportation.ProblemParams{
// 		Demand: []int{20, 40, 30, 10, 50, 25},
// 		Costs: [][]float64{
// 			{1, 2, 1, 4, 5, 2},
// 			{3, 3, 2, 1, 4, 3},
// 			{4, 2, 5, 9, 6, 2},
// 			{3, 1, 7, 3, 4, 6},
// 		},
// 	}
// 	if _, isVal := IsValidParams(params); isVal {
// 		t.Errorf("TestIsValidParams test expected [%+v], actual [%+v]", false, isVal)
// 	}
// 	m := map[string]string{"supply": "cannot be empty"}
// 	if err, _ := IsValidParams(params); m["supply"] != err.FieldName["supply"] {
// 		t.Errorf("TestIsValidParams test expected [%+v], actual [%+v]", m, err.FieldName)
// 	}
// }
// func TestIsInvalidDemand(t *testing.T) {
// 	params := transportation.ProblemParams{
// 		Supply: []int{30, 50, 75},
// 		Costs: [][]float64{
// 			{1, 2, 1, 4, 5, 2},
// 			{3, 3, 2, 1, 4, 3},
// 			{4, 2, 5, 9, 6, 2},
// 			{3, 1, 7, 3, 4, 6},
// 		},
// 	}
// 	if _, isVal := IsValidParams(params); isVal {
// 		t.Errorf("TestIsValidParams test expected [%+v], actual [%+v]", false, isVal)
// 	}
// 	m := map[string]string{"demand": "cannot be empty"}
// 	if err, _ := IsValidParams(params); m["demand"] != err.FieldName["demand"] {
// 		t.Errorf("TestIsValidParams test expected [%+v], actual [%+v]", m, err.FieldName)
// 	}
// }

// func TestIsInvalidEmptyCosts(t *testing.T) {
// 	params := transportation.ProblemParams{
// 		Demand: []int{20, 40, 30, 10, 50, 25},
// 		Supply: []int{30, 50, 75},
// 	}
// 	if _, isVal := IsValidParams(params); isVal {
// 		t.Errorf("TestIsValidParams test expected [%+v], actual [%+v]", false, isVal)
// 	}
// 	m := map[string]string{"costs": "cannot be empty"}
// 	if err, _ := IsValidParams(params); m["costs"] != err.FieldName["costs"] {
// 		t.Errorf("TestIsValidParams test expected [%+v], actual [%+v]", m, err.FieldName)
// 	}
// }

// func TestAsJSON(t *testing.T) {
// 	pE := ParamsError{map[string]string{"costs": "cannot be empty"}}
// 	if string(pE.AsJSON()) != `{"error":{"costs":"cannot be empty"}}` {
// 		t.Errorf("TestAsJSON test expected [%+v], actual [%+v]", `{"error":{"costs":"cannot be empty"}}`, string(pE.AsJSON()))
// 	}
// }
