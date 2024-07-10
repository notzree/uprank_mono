-- Modify "upwork_freelancers" table
ALTER TABLE "upwork_freelancers" DROP COLUMN "upwork_freelancer_freelancer_inference_data";
-- Modify "freelancer_inference_data" table
ALTER TABLE "freelancer_inference_data" ADD COLUMN "upwork_freelancer_freelancer_inference_data" character varying NOT NULL, ADD CONSTRAINT "freelancer_inference_data_upwork_freelancers_freelancer_inferen" FOREIGN KEY ("upwork_freelancer_freelancer_inference_data") REFERENCES "upwork_freelancers" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
-- Create index "freelancer_inference_data_upwork_freelancer_freelancer_inferenc" to table: "freelancer_inference_data"
CREATE UNIQUE INDEX "freelancer_inference_data_upwork_freelancer_freelancer_inferenc" ON "freelancer_inference_data" ("upwork_freelancer_freelancer_inference_data");
