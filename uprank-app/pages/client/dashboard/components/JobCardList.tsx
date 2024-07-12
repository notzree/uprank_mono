import { Job } from "@/types/job";
import JobCard from "./JobCard";


export default function JobCardList({jobs}: {jobs: Job[]}){
    return (
        <div>
            {jobs?.map((job) => {
                return (
                    <JobCard key={job.id} job={job} />
                );
            })}
        </div>
    )
}