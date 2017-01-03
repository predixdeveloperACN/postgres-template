package server

import (
	"net/http"
	"encoding/json"
	"fmt"
	"database/sql"

	st "github.com/predixdeveloperACN/postgres-template/storage"
)

func HandleGetAnimals(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	animals, err := st.GetAnimals()
	if err == sql.ErrNoRows {
		http.Error(w, "{}", http.StatusNoContent)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if len(animals) == 0 {
		http.Error(w, "{}", http.StatusNoContent)
		return
	}

	outJSON, err := json.Marshal(animals)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(outJSON))
}

