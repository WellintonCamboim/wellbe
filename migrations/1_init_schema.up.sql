BEGIN;

-- Habilita extensão UUID
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Criação dos tipos ENUM
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'emotion_type') THEN
        CREATE TYPE emotion_type AS ENUM (
            'happy', 'sad', 'neutral', 'calm',
            'anxious', 'stressed', 'excited', 'tired'
        );
    END IF;
    
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'day_period') THEN
        CREATE TYPE day_period AS ENUM (
            'morning', 'afternoon', 'evening'
        );
    END IF;
END$$;

-- Tabelas básicas
CREATE TABLE IF NOT EXISTS "user" (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    email VARCHAR(255) UNIQUE NOT NULL,
    birth_date DATE NOT NULL,
    phone VARCHAR(20),
    profession VARCHAR(100),
    education_level VARCHAR(100),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS user_address (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL REFERENCES "user"(id) ON DELETE CASCADE,
    street VARCHAR(255),
    number VARCHAR(20),
    complement VARCHAR(100),
    neighborhood VARCHAR(100),
    city VARCHAR(100) NOT NULL,
    state VARCHAR(100),
    country VARCHAR(100) NOT NULL DEFAULT 'Brasil',
    postal_code VARCHAR(20),
    is_primary BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT unique_user_address UNIQUE (user_id)
);

CREATE TABLE IF NOT EXISTS emotion_log (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL REFERENCES "user"(id) ON DELETE CASCADE,
    emotion emotion_type NOT NULL,
    period day_period NOT NULL,
    logged_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    notes TEXT
);

-- Demais tabelas (task, skill, sleep_record)...
CREATE TABLE IF NOT EXISTS task (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL REFERENCES "user"(id) ON DELETE CASCADE,
    title VARCHAR(100) NOT NULL,
    description TEXT,
    due_date DATE,
    is_completed BOOLEAN NOT NULL DEFAULT FALSE,
    completed_at TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS skill (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL REFERENCES "user"(id) ON DELETE CASCADE,
    name VARCHAR(100) NOT NULL,
    proficiency_level INTEGER NOT NULL CHECK (proficiency_level BETWEEN 1 AND 10),
    category VARCHAR(50),
    last_practiced DATE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS sleep_record (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL REFERENCES "user"(id) ON DELETE CASCADE,
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

COMMIT;