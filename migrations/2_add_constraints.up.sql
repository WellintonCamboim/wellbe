BEGIN;

-- Função imutável para conversão de timestamp para date
CREATE OR REPLACE FUNCTION to_date_immutable(timestamp WITH TIME ZONE) 
RETURNS date
LANGUAGE sql
IMMUTABLE
AS $$
    SELECT $1::date;
$$;

-- Índice para substituir a constraint de unicidade
CREATE UNIQUE INDEX IF NOT EXISTS unique_emotion_per_period 
ON emotion_log (user_id, to_date_immutable(logged_at), period);

-- Adicione índices para melhor performance
CREATE INDEX IF NOT EXISTS idx_task_user ON task(user_id);
CREATE INDEX IF NOT EXISTS idx_skill_user ON skill(user_id);
CREATE INDEX IF NOT EXISTS idx_sleep_record_user ON sleep_record(user_id);

COMMIT;