-- Modify "upwork_freelancers" table
ALTER TABLE "upwork_freelancers" ADD COLUMN "updated_at" timestamptz NULL, ADD COLUMN "embedded_at" timestamptz NULL;

-- Set default value for "updated_at" column
UPDATE "upwork_freelancers" SET "updated_at" = NOW() WHERE "updated_at" IS NULL;

-- Alter the column to add the NOT NULL constraint
ALTER TABLE "upwork_freelancers" ALTER COLUMN "updated_at" SET NOT NULL;

-- Rename a column from "uprank_updated_at" to "created_at"
ALTER TABLE "upwork_freelancers" RENAME COLUMN "uprank_updated_at" TO "created_at";

-- Modify "upwork_jobs" table (split into steps to handle the NOT NULL constraint)
ALTER TABLE "upwork_jobs" ADD COLUMN "updated_at" timestamptz NULL, ADD COLUMN "embedded_at" timestamptz NULL, ADD COLUMN "ranked_at" timestamptz NULL;

-- Set default value for "updated_at" column in "upwork_jobs"
UPDATE "upwork_jobs" SET "updated_at" = NOW() WHERE "updated_at" IS NULL;

-- Alter the column to add the NOT NULL constraint
ALTER TABLE "upwork_jobs" ALTER COLUMN "updated_at" SET NOT NULL;

-- Modify "work_histories" table
ALTER TABLE "work_histories" ADD COLUMN "embedded_at" timestamptz NULL, ADD COLUMN "created_at" timestamptz NULL, ADD COLUMN "updated_at" timestamptz NULL;

UPDATE "work_histories" SET "updated_at" = NOW() WHERE "updated_at" IS NULL;

UPDATE "work_histories" SET "created_at" = NOW() WHERE "created_at" IS NULL;

ALTER TABLE "work_histories" ALTER COLUMN "updated_at" SET NOT NULL, ALTER COLUMN "created_at" SET NOT NULL;
