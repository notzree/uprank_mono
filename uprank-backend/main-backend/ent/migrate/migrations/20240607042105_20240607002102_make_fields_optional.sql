-- Modify "work_histories" table
ALTER TABLE "work_histories" ALTER COLUMN "overall_rating" DROP NOT NULL, ALTER COLUMN "freelancer_earnings" DROP NOT NULL, ALTER COLUMN "budget" DROP NOT NULL, ALTER COLUMN "client_rating" DROP NOT NULL, ALTER COLUMN "client_total_jobs_posted" DROP NOT NULL, ALTER COLUMN "client_total_spend" DROP NOT NULL;
