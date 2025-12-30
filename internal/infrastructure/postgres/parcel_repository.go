package postgres

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/saad7890/flowship/internal/domain"
	"github.com/saad7890/flowship/internal/repository"
)

type ParcelRepository struct {
	db *sql.DB
}

func NewParcelRepository(db *sql.DB) repository.ParcelRepository {
	return &ParcelRepository{db: db}
}

func (r *ParcelRepository) Create(parcel *domain.Parcel) error {
	query := `
		INSERT INTO parcels (
			tracking_number,
			sender_name, sender_phone,
			receiver_name, receiver_phone,
			pickup_address, delivery_address,
			cod_amount, status,
			created_at, updated_at
		)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)
		RETURNING id
	`

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return r.db.QueryRowContext(
		ctx,
		query,
		parcel.TrackingNumber,
		parcel.SenderName,
		parcel.SenderPhone,
		parcel.ReceiverName,
		parcel.ReceiverPhone,
		parcel.PickupAddress,
		parcel.DeliveryAddress,
		parcel.CODAmount,
		parcel.Status,
		parcel.CreatedAt,
		parcel.UpdatedAt,
	).Scan(&parcel.ID)
}


func (r *ParcelRepository) FindByTrackingNumber(trackingNumber string) (*domain.Parcel, error) {
	query := `
		SELECT id, tracking_number,
		       sender_name, sender_phone,
		       receiver_name, receiver_phone,
		       pickup_address, delivery_address,
		       cod_amount, status,
		       created_at, updated_at
		FROM parcels
		WHERE tracking_number = $1
	`

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	row := r.db.QueryRowContext(ctx, query, trackingNumber)

	var p domain.Parcel
	err := row.Scan(
		&p.ID,
		&p.TrackingNumber,
		&p.SenderName,
		&p.SenderPhone,
		&p.ReceiverName,
		&p.ReceiverPhone,
		&p.PickupAddress,
		&p.DeliveryAddress,
		&p.CODAmount,
		&p.Status,
		&p.CreatedAt,
		&p.UpdatedAt,
	)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, errors.New("parcel not found")
	}

	return &p, err
}


func (r *ParcelRepository) Update(parcel *domain.Parcel) error {
	query := `
		UPDATE parcels
		SET status = $1,
		    updated_at = $2
		WHERE id = $3
	`

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := r.db.ExecContext(
		ctx,
		query,
		parcel.Status,
		time.Now(),
		parcel.ID,
	)

	return err
}
func (r *ParcelRepository) FindByID(id int64) (*domain.Parcel, error) {
	query := `
		SELECT id, tracking_number,
		       sender_name, sender_phone,
		       receiver_name, receiver_phone,
		       pickup_address, delivery_address,
		       cod_amount, status,
		       created_at, updated_at
		FROM parcels
		WHERE id = $1
	`

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	row := r.db.QueryRowContext(ctx, query, id)

	var p domain.Parcel
	err := row.Scan(
		&p.ID,
		&p.TrackingNumber,
		&p.SenderName,
		&p.SenderPhone,
		&p.ReceiverName,
		&p.ReceiverPhone,
		&p.PickupAddress,
		&p.DeliveryAddress,
		&p.CODAmount,
		&p.Status,
		&p.CreatedAt,
		&p.UpdatedAt,
	)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, errors.New("parcel not found")
	}

	return &p, err
}
