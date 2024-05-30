package server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

type ChecklistItem struct {
	Id          int       `json:"id"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	ExpireAt    time.Time `json:"expire_at"`
}

// create
func createChecklistItem(w http.ResponseWriter, r *http.Request) {
	r.Header.Set("Accept", "application/json")
	decoder := json.NewDecoder(r.Body)
	var ci ChecklistItem
	for {
		err := decoder.Decode(&ci)
		if err != nil {
			if err == io.EOF {
				break
			}
			http.Error(w, "an error occured", http.StatusInternalServerError)
			return
		}
	}

	// save to db
	// ci = add()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&ci)
}

// get item by id
func getChecklistItem(w http.ResponseWriter, r *http.Request) {
	strChID := r.PathValue("checklistID")
	itemID := r.PathValue("id")
	checklistID, err := strconv.Atoi(strChID)
	checklistItemID, err := strconv.Atoi(itemID)
	if err != nil {
		http.Error(w, fmt.Sprintf("invalid path value: %v", r.URL.Path), http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "%d %d", checklistID, checklistItemID)

	// retrieve from db
	// ci := retrieve(checklistID, checklistItemID)
	// w.Header().Set("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(&ci)
}

// get all items by checklist id
func getChecklistItems(w http.ResponseWriter, r *http.Request) {
	strChID := r.PathValue("checklistID")
	checklistID, err := strconv.Atoi(strChID)
	if err != nil {
		http.Error(w, fmt.Sprintf("invalid path value: %v", r.URL.Path), http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "%d", checklistID)
	// retrieve from db
	// items := retrieve(id)
	// w.Header().Set("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(&items)
}

// update
func updateChecklistItem(w http.ResponseWriter, r *http.Request) {
	strChID := r.PathValue("checklistID")
	itemID := r.PathValue("id")
	checklistItemID, err := strconv.Atoi(itemID)
	checklistID, err := strconv.Atoi(strChID)
	if err != nil {
		http.Error(w, fmt.Sprintf("invalid path value: %v", r.URL.Path), http.StatusNotFound)
		return
	}

	decoder := json.NewDecoder(r.Body)
	var ci ChecklistItem
	for {
		err := decoder.Decode(&ci)
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
	// items := ci.update(checklistID, checklistItemID)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "%d %d\n", checklistID, checklistItemID)
	json.NewEncoder(w).Encode(&ci)
}

// delete
func deleteChecklistItem(w http.ResponseWriter, r *http.Request) {
	strChID := r.PathValue("checklistID")
	itemID := r.PathValue("id")
	checklistItemID, err := strconv.Atoi(itemID)
	checklistID, err := strconv.Atoi(strChID)
	if err != nil {
		http.Error(w, fmt.Sprintf("invalid path value: %v", r.URL.Path), http.StatusNotFound)
		return
	}

	// delete on db
	// err := delete(checklistID, checklistItemID)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "%d %d\n", checklistID, checklistItemID)
}
