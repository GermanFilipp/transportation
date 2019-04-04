package solver

import (
	"container/list"
	"math"
	"sort"

	"github.com/germanfilipp/transportation/utils"
	"github.com/germanfilipp/transportation/utils/logger"
)

//Shipment should
type shipment struct {
	quantity, costPerUnit     float64
	indexInRow, indexinColumn int
}

var shipZero = shipment{}

//TransportTable save
type TransportTable struct {
	supply, demand, origSup, origDem []int
	supplyLen, demandLen             int
	costs                            [][]float64
	matrix                           [][]shipment
	rowsDone, columnsDone            []bool
}

type potential struct {
	value       int
	isAvailable bool
	i, j        int
}

//ResponseJSON json metadatas
type ResponseJSON struct {
	TotalCost int         `json:"totalCost"`
	Result    [][]int     `json:"result"`
	Supply    []int       `json:"supply"`
	Demand    []int       `json:"demand"`
	Costs     [][]float64 `json:"costs"`
}

//NewTransportTable initialize new TransportTable struct and made balancing validations
func NewTransportTable(supply, demand []int, costs [][]float64) TransportTable {
	if diff := utils.TotalSum(supply) - utils.TotalSum(demand); diff > 0 {
		demand = append(demand, diff)
		for i, v := range costs {
			costs[i] = append(v, 0)
		}
	} else if diff < 0 {
		supply = append(supply, -diff)
		costs = append(costs, make([]float64, len(demand)))
	}
	sLen, dLen := len(supply), len(demand)
	return TransportTable{
		supply,
		demand,
		utils.CopyArray(supply),
		utils.CopyArray(demand),
		sLen,
		dLen,
		costs,
		initMatrix(sLen, dLen),
		make([]bool, sLen),
		make([]bool, dLen),
	}
}

func initMatrix(lSup, lDem int) [][]shipment {
	matrix := make([][]shipment, lSup)
	for i := 0; i < lSup; i++ {
		matrix[i] = make([]shipment, lDem)
	}
	return matrix
}

//SolveByPotentialsMethod should
func (t *TransportTable) SolveByPotentialsMethod() {
	if t.isDegeneracySolution() {
		logger.Warn("Degeneracy solution ", string(t.ResultToJSON()))
		t.fixDegeneracy()
	}

	potentialCandidate := t.searchCandidate()
	logger.Info("Search candidate ", string(t.ResultToJSON()))
	if potentialCandidate.isAvailable {
		logger.Info("Start solving with potentials ", string(t.ResultToJSON()))
		freePoint := shipment{
			0,
			t.costs[potentialCandidate.i][potentialCandidate.j],
			potentialCandidate.i,
			potentialCandidate.j,
		}
		move, leaving := t.buildSquare(freePoint)
		if move != nil {
			plus := true
			for _, s := range move {
				if plus {
					s.quantity += leaving.quantity
				} else {
					s.quantity -= leaving.quantity
				}
				t.matrix[s.indexInRow][s.indexinColumn] = shipZero
				if s.quantity != 0 {
					t.matrix[s.indexInRow][s.indexinColumn] = s
				}
				plus = !plus
			}
			t.SolveByPotentialsMethod()
		}
	}
}

func (t *TransportTable) buildSquare(fP shipment) ([]shipment, shipment) {
	path := t.getClosedPath(fP)
	reduction := 0.0
	leavingCandidate := shipZero
	plus := true
	lowestQuantity := float64(math.MaxInt32)
	for _, s := range path {
		if plus {
			reduction += s.costPerUnit
		} else {
			reduction -= s.costPerUnit
			if s.quantity < lowestQuantity {
				leavingCandidate = s
				lowestQuantity = s.quantity
			}
		}
		plus = !plus
	}
	return path, leavingCandidate
}

func (t *TransportTable) searchCandidate() potential {
	u, v := t.potentialMethod()
	var candidates []potential
	for i, cRow := range t.costs {
		for j, c := range cRow {
			if t.matrix[i][j].quantity != 0 && t.matrix[i][j].costPerUnit != 0 {
				continue
			}
			//formula Ui + Vj > Cij
			if u[i].value+v[j].value > int(c) {
				candidates = append(candidates, potential{u[i].value + v[j].value - int(c), true, i, j})
			}
		}
	}
	if len(candidates) != 0 {
		sort.Slice(candidates, func(i, j int) bool { return candidates[i].value > candidates[j].value })
		return candidates[0]
	}
	return potential{}
}

//PotentialMethod should
func (t *TransportTable) potentialMethod() ([]potential, []potential) {
	u, v := make([]potential, t.supplyLen), make([]potential, t.demandLen)
	fb := t.firstBasis()
	u[0] = potential{0, true, fb.indexInRow, fb.indexinColumn}
	for isAnyAvailable(u) && isAnyAvailable(v) {
		for i, mRow := range t.matrix {
			for j, mx := range mRow {
				if mx != shipZero {
					if !u[i].isAvailable && !v[j].isAvailable {
						continue
					}
					//formula = Uij + Vij = Cij
					if u[i].isAvailable {
						v[j] = potential{int(mx.costPerUnit) - u[i].value, true, mx.indexInRow, fb.indexinColumn}
					}
					if v[j].isAvailable {
						u[i] = potential{int(mx.costPerUnit) - v[j].value, true, mx.indexInRow, fb.indexinColumn}
					}
				}
			}
		}
	}
	return u, v
}

func isAnyAvailable(p []potential) bool {
	for _, v := range p {
		if !v.isAvailable {
			return true
		}
	}
	return false
}

func (t *TransportTable) firstBasis() shipment {
	for _, v := range t.matrix {
		for _, m := range v {
			if m != shipZero {
				return m
			}
		}
	}
	return shipZero
}

func (t *TransportTable) getClosedPath(s shipment) []shipment {
	matrix := t.matrixToList()
	matrix.PushFront(s)
	var next *list.Element
	for {
		removals := 0
		for e := matrix.Front(); e != nil; e = next {
			next = e.Next()
			nbrs := t.getNeighbors(e.Value.(shipment), matrix)
			if nbrs[0] == shipZero || nbrs[1] == shipZero {
				matrix.Remove(e)
				removals++
			}
		}
		if removals == 0 {
			break
		}
	}
	stones := make([]shipment, matrix.Len())
	prev := s
	for i := 0; i < len(stones); i++ {
		stones[i] = prev
		prev = t.getNeighbors(prev, matrix)[i%2]
	}
	return stones
}

func (t *TransportTable) getNeighbors(s shipment, lst *list.List) [2]shipment {
	var nbrs [2]shipment
	for e := lst.Front(); e != nil; e = e.Next() {
		o := e.Value.(shipment)
		if o != s {
			if o.indexInRow == s.indexInRow && nbrs[0] == shipZero {
				nbrs[0] = o
			} else if o.indexinColumn == s.indexinColumn && nbrs[1] == shipZero {
				nbrs[1] = o
			}
			if nbrs[0] != shipZero && nbrs[1] != shipZero {
				break
			}
		}
	}
	return nbrs
}

func (t *TransportTable) matrixToList() *list.List {
	l := list.New()
	for _, m := range t.matrix {
		for _, s := range m {
			if s != shipZero {
				l.PushBack(s)
			}
		}
	}
	return l
}
