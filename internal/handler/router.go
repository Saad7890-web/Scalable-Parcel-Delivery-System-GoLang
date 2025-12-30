package handler

import (
	"net/http"

	"github.com/saad7890/flowship/internal/usecase"
)

func NewRouter(parcelUsecase *usecase.ParcelUsecase) *http.ServeMux {
	mux := http.NewServeMux()

	parcelHandler := NewParcelHandler(parcelUsecase)

	mux.HandleFunc("POST /parcels", parcelHandler.CreateParcel)
	mux.HandleFunc("GET /parcels/{trackingNumber}", parcelHandler.GetParcel)
	mux.HandleFunc("/health", HealthCheck)

	return mux
}
