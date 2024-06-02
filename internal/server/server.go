package server

import (
	"net/http"

	"github.com/JamiuJimoh/checklist/internal/checklist"
	"github.com/JamiuJimoh/checklist/internal/item"
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

	http.HandleFunc("POST /checklists/{checklistID}/items", item.CreateChecklistItem)
	http.HandleFunc("GET /checklists/{checklistID}/items", item.GetChecklistItems)
	http.HandleFunc("GET /items/{id}", item.GetChecklistItem)
	http.HandleFunc("PATCH /items/{id}", item.UpdateChecklistItem)
	http.HandleFunc("DELETE /items/{id}", item.DeleteChecklistItem)

	return server
}
