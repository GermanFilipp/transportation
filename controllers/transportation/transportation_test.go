package transportation

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestSuccessCreate(t *testing.T) {
	rr := preparation(t, `{ 
		"demand": [40,30,10,50,25], 
		"supply": [30, 50, 75, 20], 
		"costs": [[40,30,10,50,25],[40,30,10,50,25],[40,30,10,50,25],[40,30,10,50,25]]
	}`)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("TestSuccessCreate returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := `{"totalCost":5725,"result":[[0,0,10,0,0,20],[0,25,0,0,25,0],[40,5,0,30,0,0],[0,0,0,20,0,0]],"supply":[30,50,75,20],"demand":[40,30,10,50,25,20],"costs":[[40,30,10,50,25,0],[40,30,10,50,25,0],[40,30,10,50,25,0],[40,30,10,50,25,0]]}`

	if rr.Body.String() != expected {
		t.Errorf("TestSuccessCreate returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestValidationErrorCreate(t *testing.T) {
	rr := preparation(t, `{ 
		"demand": [40,30,10,50,25], 
		"supply": [30, 50, 75, 20], 
		"costs": [[40,30,10,50,25],[40,30,10,50,25],[40,30,10,50,25],[40,30,10]]
	}`)
	if status := rr.Code; status != http.StatusUnprocessableEntity {
		t.Errorf("TestSuccessCreate returned wrong status code: got %v want %v", status, http.StatusUnprocessableEntity)
	}

	expected := `{"error":{"costs":"length of costs columns should be equal to demand"}}`

	if reflect.DeepEqual(rr.Body.String(), expected) {
		t.Errorf("TestSuccessCreate returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestJSONError(t *testing.T) {
	rr := preparation(t, `{`)

	if status := rr.Code; status != http.StatusUnprocessableEntity {
		t.Errorf("TestSuccessCreate returned wrong status code: got %v want %v", status, http.StatusUnprocessableEntity)
	}
}

func preparation(t *testing.T, jsonBody string) *httptest.ResponseRecorder {
	req, err := http.NewRequest("POST", "/transport-solution", bytes.NewBuffer([]byte(jsonBody)))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Create)

	handler.ServeHTTP(rr, req)
	return rr
}
