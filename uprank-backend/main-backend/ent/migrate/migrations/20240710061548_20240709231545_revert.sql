ALTER TABLE "freelancer_inference_data" RENAME COLUMN "upwork_freelancer_id" TO "upwork_freelancer_freelancer_inference_data";
-- Drop index "freelancer_inference_data_upwork_freelancer_id_key" from table: "freelancer_inference_data"
DROP INDEX "freelancer_inference_data_upwork_freelancer_id_key";
-- Create index "freelancer_inference_data_upwork_freelancer_freelancer_inferenc" to table: "freelancer_inference_data"
CREATE UNIQUE INDEX "freelancer_inference_data_upwork_freelancer_freelancer_inferenc" ON "freelancer_inference_data" ("upwork_freelancer_freelancer_inference_data");
-- Rename a column from "upwork_freelancer_id" to "upwork_freelancer_freelancer_inference_data"

