-- Rename a column from "upwork_freelancer_freelancer_inference_data" to "upwork_freelancer_id"
ALTER TABLE "freelancer_inference_data" RENAME COLUMN "upwork_freelancer_freelancer_inference_data" TO "upwork_freelancer_id";

-- Drop index "freelancer_inference_data_upwork_freelancer_freelancer_inferenc" from table: "freelancer_inference_data"
DROP INDEX "freelancer_inference_data_upwork_freelancer_freelancer_inferenc";
-- Create index "freelancer_inference_data_upwork_freelancer_id_key" to table: "freelancer_inference_data"
CREATE UNIQUE INDEX "freelancer_inference_data_upwork_freelancer_id_key" ON "freelancer_inference_data" ("upwork_freelancer_id");

