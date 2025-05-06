BEGIN;

DROP INDEX IF EXISTS unique_emotion_per_period;
DROP INDEX IF EXISTS idx_task_user;
DROP INDEX IF EXISTS idx_skill_user;
DROP INDEX IF EXISTS idx_sleep_record_user;
DROP FUNCTION IF EXISTS to_date_immutable;

COMMIT;