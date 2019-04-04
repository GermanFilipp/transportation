package solver

import (
	"encoding/json"

	"github.com/germanfilipp/transportation/utils/logger"
)

//ResponseJSON json metadata
type ResponseJSON struct {
	TotalCost int         `json:"totalCost"`
	Result    [][]int     `json:"result"`
	Supply    []int       `json:"supply"`
	Demand    []int       `json:"demand"`
	Costs     [][]float64 `json:"costs"`
}

//ResultToJSON return json with solution results
func (t *TransportTable) ResultToJSON() []byte {
	var totalCosts float64
	var matrix [][]int
	for _, mRow := range t.matrix {
		var mx []int
		for _, v := range mRow {
			mx = append(mx, int(v.quantity))
			totalCosts += v.quantity * v.costPerUnit
		}
		matrix = append(matrix, mx)
	}
	rJ := ResponseJSON{int(totalCosts), matrix, t.origSup, t.origDem, t.costs}
	data, err := json.Marshal(rJ)
	if err != nil {
		logger.Fatal("Marshaling Error: ", err)
	}
	return data
}
