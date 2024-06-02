package item

import (
	"context"
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

var ctx = context.Background()

// create
func CreateChecklistItem(w http.ResponseWriter, r *http.Request) {
	strChID := r.PathValue("checklistID")
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
			http.Error(w, "an error occured", http.StatusInternalServerError)
			return
		}
	}

	// save to db
	err = AddChecklistItem(checklistID, &ci)
	if err != nil {
		http.Error(w, "error occured while saving checklistItem", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&ci)
}

// get item by id
func GetChecklistItem(w http.ResponseWriter, r *http.Request) {
	itemID := r.PathValue("id")
	id, err := strconv.Atoi(itemID)
	if err != nil {
		http.Error(w, fmt.Sprintf("invalid path value: %v", r.URL.Path), http.StatusNotFound)
		return
	}

	// retrieve from db
	ci, err := retrieveById(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("could not find item with id: %v", id), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&ci)
}

// get all items by checklist id
func GetChecklistItems(w http.ResponseWriter, r *http.Request) {
	strChID := r.PathValue("checklistID")
	checklistID, err := strconv.Atoi(strChID)
	if err != nil {
		http.Error(w, fmt.Sprintf("invalid path value: %v", r.URL.Path), http.StatusNotFound)
		return
	}

	// retrieve from db
	items, err := RetrieveByChecklistID(checklistID)

	if err != nil {
		http.Error(w, fmt.Sprintf("could not find items belonging to checklist with id: %v", checklistID), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&items)
}

// update
func UpdateChecklistItem(w http.ResponseWriter, r *http.Request) {
	itemID := r.PathValue("id")
	id, err := strconv.Atoi(itemID)
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
	ci.Id = id
	err = update(&ci)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&ci)
}

// delete
func DeleteChecklistItem(w http.ResponseWriter, r *http.Request) {
	itemID := r.PathValue("id")
	id, err := strconv.Atoi(itemID)
	if err != nil {
		http.Error(w, fmt.Sprintf("invalid path value: %v", r.URL.Path), http.StatusNotFound)
		return
	}

	// delete on db
	err = deleteItem(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	message, _ := json.Marshal(`{"deleted": true}`)
	fmt.Fprintln(w, string(message))
}
