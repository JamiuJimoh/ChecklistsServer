package server

import (
	"net/http"

	"github.com/JamiuJimoh/checklist/internal/checklist"
)

func NewServer() (server *http.Server) {
	server = &http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("GET /checklists", checklist.GetChecklists)
	http.HandleFunc("GET /checklists/{id}", checklist.GetChecklist)
	http.HandleFunc("POST /checklists", checklist.CreateChecklist)
	http.HandleFunc("PATCH /checklists/{id}", checklist.UpdateChecklist)
	http.HandleFunc("DELETE /checklists/{id}", checklist.DeleteChecklist)

	// http.HandleFunc("GET /checklists/{checklistID}/items", getChecklistItems)
	// http.HandleFunc("GET /checklists/{checklistID}/items/{id}", getChecklistItem)
	// http.HandleFunc("POST /checklists/{checklistID}/items", createChecklistItem)
	// http.HandleFunc("PUT /checklists/{checklistID}/items/{id}", updateChecklistItem)
	// http.HandleFunc("DELETE /checklists/{checklistID}/items/{id}", deleteChecklistItem)

	return server
}
