-- Modify "upwork_jobs" table
ALTER TABLE "upwork_jobs" DROP COLUMN "user_upworkjob";
-- Create "user_upworkjob" table
CREATE TABLE "user_upworkjob" ("user_id" character varying NOT NULL, "upwork_job_id" character varying NOT NULL, PRIMARY KEY ("user_id", "upwork_job_id"), CONSTRAINT "user_upworkjob_upwork_job_id" FOREIGN KEY ("upwork_job_id") REFERENCES "upwork_jobs" ("id") ON UPDATE NO ACTION ON DELETE CASCADE, CONSTRAINT "user_upworkjob_user_id" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE CASCADE);
