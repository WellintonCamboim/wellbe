CREATE TABLE emotion_log (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES "user"(id) ON DELETE CASCADE,
    emotion emotion_type NOT NULL,
    period day_period NOT NULL,
    logged_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    notes TEXT,
    CONSTRAINT unique_emotion_per_period UNIQUE (user_id, DATE(logged_at), period)
);

CREATE INDEX idx_emotion_log_user ON emotion_log(user_id);