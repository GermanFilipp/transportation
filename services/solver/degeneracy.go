package solver

import "math"

//isDegeneracySolution check if solution degenerate
func (t *TransportTable) isDegeneracySolution() bool {
	var counter int
	for _, mx := range t.matrix {
		for _, v := range mx {
			if v != shipZero {
				counter++
			}
		}
	}
	if t.demandLen+t.supplyLen-1 > counter {
		return true
	}
	return false
}

//fixDegeneracy fix infinity loop
func (t *TransportTable) fixDegeneracy() {
	for i := 0; i < t.supplyLen; i++ {
		for j := 0; j < t.demandLen; j++ {
			if t.matrix[i][j] == shipZero {
				fakeShipment := shipment{math.SmallestNonzeroFloat64, t.costs[i][j], i, j}
				if len(t.getClosedPath(fakeShipment)) == 0 {
					t.matrix[i][j] = fakeShipment
					return
				}
			}
		}
	}
}
