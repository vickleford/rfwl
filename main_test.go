package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHzToWavelength(t *testing.T) {
	var testCases = []struct {
		hz         float64
		wavelength float64
	}{
		{144e6, 2.08},
		{54e6, 5.5},
		{28e6, 10.7},
		{3.5e6, 85.6},
		{4e6, 74.9},
		{440e6, .68},
	}

	for _, tt := range testCases {
		r := HzToWavelength(tt.hz)
		if (r - tt.wavelength) > 0.1 {
			t.Fatalf("Incorrect conversion for %f hz: Expected %f, got %f", tt.hz,
				tt.wavelength, r)
		}
	}
}

func TestHealthCheckHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/v1/health", nil)
	if err != nil {
		t.Fatal(err)
	}

	// satisfies http.ResponseWriter
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HealthCheckHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: Expected %v, got %v",
			http.StatusOK, status)
	}

	expected := `{"alive": true}`
	if rr.Body.String() != expected {
		t.Errorf("Handler returned unexpected body: Expected %v, got %v", expected,
			rr.Body.String())
	}
}

func TestWavelengthHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/v1/wavelength?hz=144000000", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(WavelengthHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: Expected %v, got %v",
			http.StatusOK, status)
	}

	expected := `{"wavelength": {"length": 2.08, "unit": "meters"}}`
	if rr.Body.String() != expected {
		t.Errorf("Handler returned unexpected body: Expected %v, got %v", expected,
			rr.Body.String())
	}
}
