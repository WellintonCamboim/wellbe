BEGIN;

DROP INDEX IF EXISTS unique_emotion_per_period;
DROP FUNCTION IF EXISTS to_date_immutable;

COMMIT;