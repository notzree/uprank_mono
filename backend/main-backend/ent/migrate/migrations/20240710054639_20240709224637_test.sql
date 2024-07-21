-- Modify "freelancer_inference_data" table
ALTER TABLE "freelancer_inference_data" DROP COLUMN "upwork_freelancer_freelancer_inference_data";
-- Modify "upwork_freelancers" table
ALTER TABLE "upwork_freelancers" ADD COLUMN "upwork_freelancer_freelancer_inference_data" bigint NULL, ADD CONSTRAINT "upwork_freelancers_freelancer_inference_data_freelancer_inferen" FOREIGN KEY ("upwork_freelancer_freelancer_inference_data") REFERENCES "freelancer_inference_data" ("id") ON UPDATE NO ACTION ON DELETE SET NULL;
