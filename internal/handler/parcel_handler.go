package handler

import (
	"encoding/json"
	"net/http"

	"github.com/saad7890/flowship/internal/handler/dto"
	"github.com/saad7890/flowship/internal/usecase"
)

type ParcelHandler struct {
	usecase *usecase.ParcelUsecase
}

func NewParcelHandler(u *usecase.ParcelUsecase) *ParcelHandler {
	return &ParcelHandler{usecase: u}
}

func (h *ParcelHandler) CreateParcel(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateParcelRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	parcel, err := h.usecase.CreateParcel(
		req.TrackingNumber,
		req.SenderName,
		req.SenderPhone,
		req.ReceiverName,
		req.ReceiverPhone,
		req.PickupAddress,
		req.DeliveryAddress,
		req.CODAmount,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp := dto.ParcelResponse{
		ID:             parcel.ID,
		TrackingNumber: parcel.TrackingNumber,
		Status:         string(parcel.Status),
		CODAmount:      parcel.CODAmount,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}


func (h *ParcelHandler) GetParcel(w http.ResponseWriter, r *http.Request) {
	trackingNumber := r.PathValue("trackingNumber")

	parcel, err := h.usecase.GetParcelByTrackingNumber(trackingNumber)
	if err != nil {
		http.Error(w, "parcel not found", http.StatusNotFound)
		return
	}

	resp := dto.ParcelResponse{
		ID:             parcel.ID,
		TrackingNumber: parcel.TrackingNumber,
		Status:         string(parcel.Status),
		CODAmount:      parcel.CODAmount,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
