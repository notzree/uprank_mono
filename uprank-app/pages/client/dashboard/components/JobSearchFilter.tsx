import { Job } from "@/types/job";
import { Input } from "@/components/ui/input";
import * as React from "react";
import Fuse from 'fuse.js';
import { return_origin_job } from "@/utils/job_utils";
export default function JobSearchFilter({ jobs, setJob }: { jobs: Job[], setJob: React.Dispatch<React.SetStateAction<Job[]>> }) {
    const [searchTerm, setSearchTerm] = React.useState('');
    const origin_title_map = new Map<string, string>();


    const fuse = new Fuse(jobs, {
        keys: [
            { name: 'title', getFn: (job) => return_origin_job(job).title || "undefined" },
        ],
        includeScore: true,
        threshold: 0.4, // Adjust the threshold for more or less fuzzy matching
    });

    const filterJobs = (term: string) => {
        const result = fuse.search(term);
        const filteredJobs = result.map(res => res.item);
        setJob(filteredJobs);
    };

    const handleInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        const term = e.target.value;
        setSearchTerm(term);
        filterJobs(term);
    };

    return (
        <div className="flex w-full max-w-sm items-center space-x-2">
            <Input 
                type="text" 
                placeholder="Job Title" 
                value={searchTerm} 
                onChange={handleInputChange} 
            />
        </div>
    );
}
