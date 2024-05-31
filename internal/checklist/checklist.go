package checklist

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/JamiuJimoh/checklist/internal/item"
)

type Checklist struct {
	Id        int                  `json:"id"`
	Title     string               `json:"title"`
	CreatedAt time.Time            `json:"created_at"`
	Items     []item.ChecklistItem `json:"items"`
}

var ctx = context.Background()

// create
func CreateChecklist(w http.ResponseWriter, r *http.Request) {
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
	err := insert(&ch)
	if err != nil {
		http.Error(w, "error occured while saving checklist", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&ch)
}

// get item by id
func GetChecklist(w http.ResponseWriter, r *http.Request) {
	checklistID := r.PathValue("id")
	id, err := strconv.Atoi(checklistID)
	if err != nil {
		http.Error(w, fmt.Sprintf("invalid path value: %v", r.URL.Path), http.StatusNotFound)
		return
	}

	// retrieve from db
	ch, err := retrieveById(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("could not find checklist with id: %v", id), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&ch)
}

// get all items by checklist id
func GetChecklists(w http.ResponseWriter, _ *http.Request) {
	// retrieve from db
	items, err := retrieveAll(10, 0)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "an error occured", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&items)
}

// update
func UpdateChecklist(w http.ResponseWriter, r *http.Request) {
	pathID := r.PathValue("id")
	id, err := strconv.Atoi(pathID)
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
	ch.Id = id
	err = update(&ch)
	if err != nil {
		http.Error(w, "error occured while updating checklist", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	message, _ := json.Marshal(`{"updated": true}`)
	w.Write(message)
}

// delete
func DeleteChecklist(w http.ResponseWriter, r *http.Request) {
	pathID := r.PathValue("id")
	id, err := strconv.Atoi(pathID)
	if err != nil {
		http.Error(w, fmt.Sprintf("invalid path value: %v", r.URL.Path), http.StatusNotFound)
		return
	}

	// delete on db
	err = delete(id)
	if err != nil {
		http.Error(w, "error occured while deleting checklist", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	message, _ := json.Marshal(`{"deleted": true}`)
	fmt.Fprintln(w, string(message))
}
