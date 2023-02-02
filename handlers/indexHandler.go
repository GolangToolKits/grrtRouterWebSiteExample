package handlers

import "net/http"

// Index Index
func (h *SiteHandler) Index(w http.ResponseWriter, r *http.Request) {

	h.Templates.ExecuteTemplate(w, indexPage, nil)
}
