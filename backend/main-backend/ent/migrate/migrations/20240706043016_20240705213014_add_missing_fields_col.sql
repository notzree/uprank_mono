-- Modify "upwork_freelancers" table
ALTER TABLE "upwork_freelancers" ADD COLUMN "missing_fields" boolean NOT NULL DEFAULT false;
