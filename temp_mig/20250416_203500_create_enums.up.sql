-- Criação dos ENUMs primeiro
CREATE TYPE emotion_type AS ENUM (
    'happy', 'sad', 'neutral', 'calm',
    'anxious', 'stressed', 'excited', 'tired'
);

CREATE TYPE day_period AS ENUM (
    'morning', 'afternoon', 'evening'
);