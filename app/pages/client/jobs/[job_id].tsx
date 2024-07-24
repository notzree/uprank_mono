import * as React from "react";
import Navbar from "../shared/components/navbar";
import FreelancerTable from "./components/freelancer-table";
import type { UpworkFreelancer } from "@/types/freelancer";

export default function Jobs({ job_prop }: { job_prop: Job }) {
    const [freelancers, setFreelancers] = React.useState<UpworkFreelancer[]>(
        job_prop.edges.upworkjob?.edges.upworkfreelancer || []
    );
    const [showFilters, setShowFilters] = React.useState(true);
    const job = job_prop.edges.upworkjob;
    if (!job) {
        return <div>Job not found</div>;
    }
    return (
        <div className="grid min-h-screen w-screen">
            <div className="flex flex-col">
                <Navbar />
                <main className="flex flex-1 flex-col pt-16">
                    <div className="flex flex-row w-screen px-6">
                       {showFilters && <FreelancerSearchFilter
                        job={job}
                        original_freelancers={job_prop.edges.upworkjob?.edges.upworkfreelancer || []} 
                        visible_freelancers={freelancers}
                        setFreelancers={setFreelancers}
                        />
                       }
                       <div className="flex flex-col overflow-auto">
                       <JobCard job={job_prop} />
                       <FreelancerTable
                            original_freelancers={job_prop.edges.upworkjob?.edges.upworkfreelancer || []}
                            visible_freelancers={freelancers}
                            setFreelancers={setFreelancers}
                            job={job}
                        />
                       </div>
                    </div>
                </main>
            </div>
        </div>
    );
}

import { getAuth } from "@clerk/nextjs/server";
import { GetServerSideProps } from "next";
import { BackendClient } from "@/backend-client/backend-client";
import { Job } from "@/types/job";
import FreelancerSearchFilter from "./components/freelancer-filter";
import JobCard from "./components/detailed-job-card";

export const getServerSideProps: GetServerSideProps = async (ctx) => {
    const { userId, getToken } = getAuth(ctx.req);
    const token = await getToken();
    const { job_id } = ctx.params!;

    if (!userId || !token || typeof job_id !== "string") {
        return {
            redirect: {
                destination: "/",
                permanent: false,
            },
        };
    }

    const client = new BackendClient();
    const job = await client.GetPlatformJobById(job_id, token);

    if (!job) {
        return {
            notFound: true,
        };
    }

    return {
        props: {
            job_prop: job,
        },
    };
};
