package problem

import (
	"encoding/json"

	"github.com/germanfilipp/transportation/utils/logger"
)

//Params interface
type Params interface {
	GetCosts() [][]float64
	GetSupply() []int
	GetDemand() []int
}

//ParamsError params
type ParamsError struct {
	FieldName map[string]string `json:"error"`
}

//NewError initialize new error
func NewError(message string) *ParamsError {
	err := ParamsError{map[string]string{"error": message}}
	return &err
}

//IsValidParams validate params
func IsValidParams(params Params) (ParamsError, bool) {
	pErr := ParamsError{map[string]string{}}
	pErr.validateEmpty(params)
	if len(pErr.FieldName) != 0 {
		return pErr, false
	}
	pErr.validateCosts(params)
	if len(pErr.FieldName) != 0 {
		return pErr, false
	}
	pErr.validateNegative(params)
	if len(pErr.FieldName) != 0 {
		return pErr, false
	}
	return pErr, true
}

//AsJSON represent as json
func (pE *ParamsError) AsJSON() []byte {
	data, err := json.Marshal(pE)
	if err != nil {
		logger.Error("invalid json ", err)
	}
	return data
}

func (pE *ParamsError) validateEmpty(params Params) {
	erStr := "cannot be empty"
	switch 0 {
	case len(params.GetDemand()):
		pE.FieldName["demand"] = erStr
	case len(params.GetSupply()):
		pE.FieldName["supply"] = erStr
	case len(params.GetCosts()):
		pE.FieldName["costs"] = erStr
	}
}

func (pE *ParamsError) validateCosts(params Params) {
	if len(params.GetCosts()) != len(params.GetSupply()) {
		pE.FieldName["costs"] = "length of costs rows should be equal to supply"
	}
	for _, cRow := range params.GetCosts() {
		if len(cRow) != len(params.GetDemand()) {
			pE.FieldName["costs"] = "length of costs columns should be equal to demand"
		}
	}
}

func (pE *ParamsError) validateNegative(params Params) {
	for _, val := range params.GetSupply() {
		if val < 0 {
			pE.FieldName["supply"] = "supply cannot include negative values"
		}
	}
	for _, val := range params.GetDemand() {
		if val < 0 {
			pE.FieldName["demand"] = "demand cannot include negative values"
		}
	}

	for _, cost := range params.GetCosts() {
		for _, val := range cost {
			if val < 0 {
				pE.FieldName["costs"] = "costs cannot include negative values"
			}
		}
	}
}
