package server

import (
	"net/http"
)

func NewServer() (server *http.Server) {
	server = &http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("GET /checklists", getChecklists)
	http.HandleFunc("GET /checklists/{id}", getChecklist)
	http.HandleFunc("POST /checklists", createChecklist)
	http.HandleFunc("PUT /checklists/{id}", updateChecklist)
	http.HandleFunc("DELETE /checklists/{id}", deleteChecklist)

	http.HandleFunc("GET /checklists/{checklistID}/items", getChecklistItems)
	http.HandleFunc("GET /checklists/{checklistID}/items/{id}", getChecklistItem)
	http.HandleFunc("POST /checklists/{checklistID}/items", createChecklistItem)
	http.HandleFunc("PUT /checklists/{checklistID}/items/{id}", updateChecklistItem)
	http.HandleFunc("DELETE /checklists/{checklistID}/items/{id}", deleteChecklistItem)

	return server
}
