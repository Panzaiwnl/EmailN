package main

import (
	"emailn/internal/domain/campaign"
	"emailn/internal/endpoints"
	"emailn/internal/infrastructure/database"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	campaignService := campaign.Service{
		Repository: &database.CampaignRepository{},
	}
	handler := endpoints.Handler{
		CampaignService: campaignService,
	}
	r.Post("/campaigns", endpoints.HandlerError(handler.CampaignsPost))
	r.Get("/campaigns", endpoints.HandlerError(handler.CampaignsGet))

	r.Get("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Funcionando"))
	})

	http.ListenAndServe(":3000", r)
}
