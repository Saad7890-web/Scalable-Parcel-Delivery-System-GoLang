package dto

type CreateParcelRequest struct {
	TrackingNumber  string  `json:"tracking_number"`
	SenderName      string  `json:"sender_name"`
	SenderPhone     string  `json:"sender_phone"`
	ReceiverName    string  `json:"receiver_name"`
	ReceiverPhone   string  `json:"receiver_phone"`
	PickupAddress   string  `json:"pickup_address"`
	DeliveryAddress string  `json:"delivery_address"`
	CODAmount       float64 `json:"cod_amount"`
}

type ParcelResponse struct {
	ID             int64   `json:"id"`
	TrackingNumber string  `json:"tracking_number"`
	Status         string  `json:"status"`
	CODAmount      float64 `json:"cod_amount"`
}
