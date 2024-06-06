-- Modify "work_histories" table
ALTER TABLE "work_histories" DROP COLUMN "fixed_charge_amount", DROP COLUMN "fixed_charge_currency", DROP COLUMN "hourly_charge_amount", DROP COLUMN "hourly_charge_currency", ADD COLUMN "is_hourly" boolean NOT NULL, ADD COLUMN "freelancer_earnings" numeric NOT NULL;
