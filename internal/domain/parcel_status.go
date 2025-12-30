package domain

type ParcelStatus string

const (
	ParcelCreated ParcelStatus = "CREATED"
	ParcelPickedUp ParcelStatus = "PICKED_UP"
	ParcelInTransit ParcelStatus = "IN_TRANSIT"
	ParcelDelivered ParcelStatus = "DELIVERED"
	ParcelFailed ParcelStatus = "FAILED"
)

