CREATE TABLE parcels (
    id BIGSERIAL PRIMARY KEY,

    tracking_number VARCHAR(100) NOT NULL UNIQUE,

    sender_name VARCHAR(255) NOT NULL,
    sender_phone VARCHAR(50) NOT NULL,

    receiver_name VARCHAR(255) NOT NULL,
    receiver_phone VARCHAR(50) NOT NULL,

    pickup_address TEXT NOT NULL,
    delivery_address TEXT NOT NULL,

    cod_amount NUMERIC(10, 2) NOT NULL DEFAULT 0,

    status VARCHAR(50) NOT NULL,

    created_at TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL
);

CREATE INDEX idx_parcels_tracking_number ON parcels(tracking_number);