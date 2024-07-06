import { Job, UpworkJob } from "@/types/job";

//todo: make the return type a union of all job types
export function return_origin_job(job: Job): UpworkJob {
    if (job.origin_platform == "upwork") {
        if (job.edges.upworkjob == null) {
            throw new Error("Upwork job not found");
        }
        console.log(job.origin_platform)
        return job.edges.upworkjob;
    }

    throw new Error("Origin platform not yet supported");
}