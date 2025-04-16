CREATE TABLE sleep_record (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES "user"(id) ON DELETE CASCADE,
    record_date DATE NOT NULL DEFAULT CURRENT_DATE,
    target_hours DECIMAL(3,1) NOT NULL CHECK (target_hours BETWEEN 4 AND 12),
    actual_hours DECIMAL(3,1) NOT NULL CHECK (actual_hours BETWEEN 0 AND 24),
    quality_rating INTEGER CHECK (quality_rating BETWEEN 1 AND 5),
    was_interrupted BOOLEAN NOT NULL DEFAULT FALSE,
    medication_used BOOLEAN NOT NULL DEFAULT FALSE,
    notes TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT unique_sleep_record UNIQUE (user_id, record_date)
);