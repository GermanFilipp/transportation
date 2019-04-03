package solver

import (
	"math"
	"sort"

	"github.com/germanfilipp/transportation/utils"
)

type minDiff struct {
	diff, minNum, col, row int
}

//VogelApproximation algoritm should find first rude results for problem
func (t *TransportTable) VogelApproximation() {
	totalSupply := utils.TotalSum(t.supply)
	for totalSupply > 0 {
		penalty := t.getPenalty()
		col, row := penalty.col, penalty.row
		demQuantity := t.demand[col]
		supQuantity := t.supply[row]
		if demQuantity > supQuantity {
			demQuantity = supQuantity
		}

		t.demand[col] -= demQuantity
		t.supply[row] -= demQuantity
		t.columnsDone[col] = t.demand[col] == 0
		t.rowsDone[row] = t.supply[row] == 0

		t.matrix[row][col] = shipment{
			quantity:      float64(demQuantity),
			costPerUnit:   t.costs[row][col],
			indexInRow:    row,
			indexinColumn: col,
		}
		totalSupply -= demQuantity
	}
}

func (t *TransportTable) getPenalty() (used minDiff) {
	used = t.findMaxPenalty(t.supplyLen, true)
	maxPInCol := t.findMaxPenalty(t.demandLen, false)
	if used.diff < maxPInCol.diff {
		used = maxPInCol
	}
	return
}

func (t *TransportTable) findMaxPenalty(lN int, isRow bool) minDiff {
	maxPenalties := []minDiff{}
	for i := 0; i < lN; i++ {
		if isRow && t.rowsDone[i] || !isRow && t.columnsDone[i] {
			continue
		}
		if isRow {
			maxPenalties = append(maxPenalties, t.diffInRow(i))
		} else {
			maxPenalties = append(maxPenalties, t.diffInColumn(i))
		}
	}
	sort.Slice(maxPenalties, func(i, j int) bool { return maxPenalties[i].diff > maxPenalties[j].diff })
	return maxPenalties[0]
}

func (t *TransportTable) diffInRow(index int) minDiff {
	var minNumIndex int
	minFirst, minSecond := math.MaxInt32, math.MaxInt32
	for i := 0; i < t.demandLen; i++ {
		if t.columnsDone[i] {
			continue
		}
		c := int(t.costs[index][i])
		if c < minSecond {
			minSecond = c
		}
		if c < minFirst {
			minSecond, minFirst, minNumIndex = minFirst, c, i
		}
	}
	return minDiff{minSecond - minFirst, minFirst, minNumIndex, index}
}

func (t *TransportTable) diffInColumn(index int) minDiff {
	var minNumIndex int
	minFirst, minSecond := math.MaxInt32, math.MaxInt32
	for i := 0; i < t.supplyLen; i++ {
		if t.rowsDone[i] {
			continue
		}
		c := int(t.costs[i][index])
		if c < minSecond {
			minSecond = c
		}
		if c < minFirst {
			minSecond, minFirst, minNumIndex = minFirst, c, i
		}
	}
	return minDiff{minSecond - minFirst, minFirst, index, minNumIndex}
}
