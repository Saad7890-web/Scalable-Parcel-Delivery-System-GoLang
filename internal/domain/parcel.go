package domain

import "time"

type Parcel struct {
	ID               int64
	TrackingNumber   string
	SenderName       string
	SenderPhone      string
	ReceiverName     string
	ReceiverPhone    string
	PickupAddress    string
	DeliveryAddress  string
	CODAmount        float64
	Status           ParcelStatus
	CreatedAt        time.Time
	UpdatedAt        time.Time
}


func NewParcel(
	trackingNumber string,
	senderName string,
	senderPhone string,
	receiverName string,
	receiverPhone string,
	pickupAddress string,
	deliveryAddress string,
	codAmount float64,
) *Parcel {
	now := time.Now()

	return &Parcel{
		TrackingNumber:  trackingNumber,
		SenderName:      senderName,
		SenderPhone:     senderPhone,
		ReceiverName:    receiverName,
		ReceiverPhone:   receiverPhone,
		PickupAddress:   pickupAddress,
		DeliveryAddress: deliveryAddress,
		CODAmount:       codAmount,
		Status:          ParcelCreated,
		CreatedAt:       now,
		UpdatedAt:       now,
	}
}


func (p *Parcel) UpdateStatus(status ParcelStatus) error {
	p.Status = status
	p.UpdatedAt = time.Now()
	return nil
}
