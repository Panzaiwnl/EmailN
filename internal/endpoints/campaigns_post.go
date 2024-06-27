package endpoints

import (
	"emailn/internal/contract"
	"github.com/go-chi/render"
	"net/http"
)

func (h *Handler) CampaignsPost(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	var request contract.NewCampaignDTO
	render.DecodeJSON(r.Body, &request)
	id, err := h.CampaignService.Create(request)
	return map[string]string{"id": id}, 201, err

}
