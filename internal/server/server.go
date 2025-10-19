package server

import (
	"fmt"
	"net/http"
	"os/exec"
)

func Start() {
	http.HandleFunc("/", refreshHouses)
	http.ListenAndServe(":8080", nil)
}

func refreshHouses(w http.ResponseWriter, r *http.Request) {
	postCode := r.URL.Query().Get("postCode")
	if postCode == "" {
		http.Error(w, "Postcode Missing from request", http.StatusBadRequest)
		return
	}

	cmd := exec.Command("bash", "/root/priceProwler/getHousePrices.sh", postCode)

	out, err := cmd.CombinedOutput()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error in execution of script: %v", err.Error()), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "Updated db with info for %s\nOutput: %s", postCode, string(out))
}
