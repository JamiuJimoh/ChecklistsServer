package server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

type Checklist struct {
	Id         int             `json:"id"`
	Title      string          `json:"title"`
	CreatedAt  time.Time       `json:"created_at"`
	Checklists []ChecklistItem `json:"checklists"`
}

// create
func createChecklist(w http.ResponseWriter, r *http.Request) {
	r.Header.Set("Accept", "application/json")
	decoder := json.NewDecoder(r.Body)
	var ch Checklist
	for {
		err := decoder.Decode(&ch)
		if err != nil {
			if err == io.EOF {
				break
			}
			http.Error(w, "an error occured", http.StatusInternalServerError)
			return
		}
	}

	// save to db
	// ch = add()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&ch)
}

// get item by id
func getChecklist(w http.ResponseWriter, r *http.Request) {
	checklistID := r.PathValue("id")
	id, err := strconv.Atoi(checklistID)
	if err != nil {
		http.Error(w, fmt.Sprintf("invalid path value: %v", r.URL.Path), http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "%d\n", id)

	// retrieve from db
	// ch := retrieve(id)
	// w.Header().Set("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(&ch)
}

// get all items by checklist id
func getChecklists(w http.ResponseWriter, r *http.Request) {
	// retrieve from db
	// items := retrieve(id)
	w.Header().Set("Content-Type", "application/json")
	// fmt.Fprintf(w, "%d", checklistID)
	json.NewEncoder(w).Encode("[]")
}

// update
func updateChecklist(w http.ResponseWriter, r *http.Request) {
	checklistID := r.PathValue("id")
	id, err := strconv.Atoi(checklistID)
	if err != nil {
		http.Error(w, fmt.Sprintf("invalid path value: %v", r.URL.Path), http.StatusNotFound)
		return
	}

	decoder := json.NewDecoder(r.Body)
	var ch Checklist
	for {
		err := decoder.Decode(&ch)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err)
			http.Error(w, "an error occured", http.StatusInternalServerError)
			return
		}
	}

	// update on db
	// items := ch.update(id)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "%d\n", id)
	json.NewEncoder(w).Encode(&ch)
}

// delete
func deleteChecklist(w http.ResponseWriter, r *http.Request) {
	checklistID := r.PathValue("id")
	id, err := strconv.Atoi(checklistID)
	if err != nil {
		http.Error(w, fmt.Sprintf("invalid path value: %v", r.URL.Path), http.StatusNotFound)
		return
	}

	// delete on db
	// err := delete(id)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "%d\n", id)
}
