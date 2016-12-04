package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

// meters per second
const lightspeed = 299792458

func HzToWavelength(hz float64) float64 {
	wavelength := lightspeed / hz
	return wavelength
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, `{"alive": true}`)
}

func WavelengthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	qs := r.URL.Query()
	frequency, _ := strconv.Atoi(qs.Get("hz"))

	wavelength := HzToWavelength(float64(frequency))
	io.WriteString(
		w, fmt.Sprintf(`{"wavelength": {"length": %.2f, "unit": "meters"}}`,
			wavelength))
}

func main() {
	http.HandleFunc("/api/v1/health", HealthCheckHandler)
	http.HandleFunc("/api/v1/wavelength", WavelengthHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
