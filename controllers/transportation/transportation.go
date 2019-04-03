package transportation

import (
	"encoding/json"
	"net/http"

	"github.com/germanfilipp/transportation/services/solver"
	"github.com/germanfilipp/transportation/utils/logger"
	validation "github.com/germanfilipp/transportation/validations/problem"
)

//ProblemParams params
type ProblemParams struct {
	Demand []int       `json:"demand"`
	Supply []int       `json:"supply"`
	Costs  [][]float64 `json:"costs"`
}

//GetDemand method for interface
func (p ProblemParams) GetDemand() []int { return p.Demand }

//GetSupply method for interface
func (p ProblemParams) GetSupply() []int { return p.Supply }

//GetCosts method for interface
func (p ProblemParams) GetCosts() [][]float64 { return p.Costs }

//Create - post action that receive params for solving trasportation problem
func Create(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if r := recover(); r != nil {
			logger.Fatal(w, "Interval Server Error", http.StatusInternalServerError)
		}
	}()

	var p ProblemParams
	jsonErr := json.NewDecoder(r.Body).Decode(&p)
	logger.Info("params", p)
	if jsonErr != nil {
		logger.Error("invalid params", p)
		http.Error(w, "Unprocessable Entity", http.StatusUnprocessableEntity)
	} else {
		validation, isValidParams := validation.IsValidParams(p)
		if isValidParams {
			result := solve(p)
			w.Header().Set("Content-Type", "application/json")
			logger.Info("OK ", http.StatusOK, string(result))
			w.Write(result)
		} else {
			errJSON := validation.AsJSON()
			logger.Error("Unprocessable Entity ", http.StatusUnprocessableEntity, p, string(errJSON))
			http.Error(w, string(errJSON), http.StatusUnprocessableEntity)
		}
	}
}

func solve(p ProblemParams) []byte {
	logger.Info("Start calculation")
	tT := solver.NewTransporTable(p.Supply, p.Demand, p.Costs)
	logger.Info("Start Vogel Approximation")
	tT.VogelApproximation()
	logger.Info("Start Calculate Potentials")
	tT.SolveByPotentialsMethod()
	logger.Info("Done!")
	return tT.ResultToJSON()
}
