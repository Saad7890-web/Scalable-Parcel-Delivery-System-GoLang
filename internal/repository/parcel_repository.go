package repository

import "github.com/saad7890/flowship/internal/domain"

type ParcelRepository interface {
	Create(parcel *domain.Parcel) error
	FindByID(id int64) (*domain.Parcel, error)
	FindByTrackingNumber(trackingNumber string) (*domain.Parcel, error)
	Update(parcel *domain.Parcel) error
}