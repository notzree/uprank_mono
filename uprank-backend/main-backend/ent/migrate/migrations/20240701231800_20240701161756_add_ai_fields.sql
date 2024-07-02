-- Modify "upwork_freelancers" table
ALTER TABLE "upwork_freelancers" DROP COLUMN "uprank_score", ADD COLUMN "uprank_specialization_score" double precision NULL DEFAULT 0, ADD COLUMN "uprank_estimated_completion_time" interval NULL;
