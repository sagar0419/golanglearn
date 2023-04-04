package api

import (
	"fmt"
	"net/http"

	"github.com/practice/new-Code/pkg"
)

// Metrics Handle func
func Metrics(w http.ResponseWriter, r *http.Request) {
	a, _ := pkg.MetricsDb()
	for _, value := range a {
		metricsvalue := fmt.Sprintf(" %v\n", value)
		fmt.Fprintf(w, metricsvalue)
	}
}
