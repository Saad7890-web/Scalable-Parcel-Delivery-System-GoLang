package usecase

import (
	"errors"

	"github.com/saad7890/flowship/internal/domain"
	"github.com/saad7890/flowship/internal/repository"
)

type ParcelUsecase struct {
	parcelRepo repository.ParcelRepository
}

func NewParcelUsecase(parcelRepo repository.ParcelRepository) *ParcelUsecase {
	return &ParcelUsecase{
		parcelRepo: parcelRepo,
	}
}

func (u *ParcelUsecase) GetParcelByTrackingNumber(trackingNumber string) (*domain.Parcel, error) {
	return u.parcelRepo.FindByTrackingNumber(trackingNumber)
}


func (u *ParcelUsecase) CreateParcel(
	trackingNumber string,
	senderName string,
	senderPhone string,
	receiverName string,
	receiverPhone string,
	pickupAddress string,
	deliveryAddress string,
	codAmount float64,
) (*domain.Parcel, error) {

	
	if trackingNumber == "" {
		return nil, errors.New("tracking number is required")
	}

	parcel := domain.NewParcel(
		trackingNumber,
		senderName,
		senderPhone,
		receiverName,
		receiverPhone,
		pickupAddress,
		deliveryAddress,
		codAmount,
	)

	if err := u.parcelRepo.Create(parcel); err != nil {
		return nil, err
	}

	return parcel, nil
}

func (u *ParcelUsecase) UpdateParcelStatus(
	trackingNumber string,
	status domain.ParcelStatus,
) error {

	parcel, err := u.parcelRepo.FindByTrackingNumber(trackingNumber)
	if err != nil {
		return err
	}

	return u.parcelRepo.Update(parcel)
}
