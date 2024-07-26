import * as React from "react";
import Navbar from "../shared/components/navbar";
import FreelancerTable from "./components/freelancer-table";
import DetailedJobCard from "./components/detailed-job-card";
import type { UpworkFreelancer } from "@/types/freelancer";
import { Button } from "@/components/ui/button";

export default function Jobs({ job_prop }: { job_prop: Job }) {
    
    const [freelancers, setFreelancers] = React.useState<UpworkFreelancer[]>(
        job_prop.edges.upworkjob?.edges.upworkfreelancer || []
    );
    const [showFilters, setShowFilters] = React.useState(true);
    const job = job_prop.edges.upworkjob;
    const original_freelancers = job_prop.edges.upworkjob?.edges.upworkfreelancer || [];
        // START temporary fix for MVP, should be computed on the backend in the future
        let average_specialization_score = 0;
        let num_specialization_scores = 0;
        let average_budget_adherence_percentage = 0;
        let num_budget_adherence_percentage = 0;
        let average_budget_overrun_percentage = 0;
        let num_budget_overrun_percentage = 0;
        for (let freelancer of original_freelancers){
            if (freelancer.edges && freelancer.edges.freelancer_inference_data){
                if (freelancer.edges.freelancer_inference_data.finalized_rating_score && freelancer.edges.freelancer_inference_data.finalized_rating_score > 0){
                    num_specialization_scores += 1;
                    average_specialization_score += freelancer.edges.freelancer_inference_data.finalized_rating_score;
                }
                if (freelancer.edges.freelancer_inference_data.budget_adherence_percentage){
                    num_budget_adherence_percentage += 1;
                    average_budget_adherence_percentage += freelancer.edges.freelancer_inference_data.budget_adherence_percentage;
                }
                if (freelancer.edges.freelancer_inference_data.budget_overrun_percentage){
                    num_budget_overrun_percentage += 1;
                    average_budget_overrun_percentage += freelancer.edges.freelancer_inference_data.budget_overrun_percentage;
                }
            }
        }

        average_specialization_score /= num_specialization_scores
        average_budget_adherence_percentage /= num_budget_adherence_percentage
        average_budget_overrun_percentage /= num_budget_overrun_percentage

        // END temporary fix for MVP

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
                        original_freelancers={original_freelancers} 
                        visible_freelancers={freelancers}
                        setFreelancers={setFreelancers}
                        />
                       }
                       <div className="flex flex-col overflow-auto">
                       <DetailedJobCard job={job_prop}
                        average_specialization_score={average_specialization_score}
                        average_budget_adherence_percentage={average_budget_adherence_percentage}
                        average_budget_overrun_percentage={average_budget_overrun_percentage}
                       />
                       <FreelancerTable
                            original_freelancers={original_freelancers}
                            visible_freelancers={freelancers}
                            setFreelancers={setFreelancers}
                            job={job}
                            average_specialization_score={average_specialization_score}
                            average_budget_adherence_percentage={average_budget_adherence_percentage}
                            average_budget_overrun_percentage={average_budget_overrun_percentage}
                            setShowFilters={setShowFilters}
                            showFilters={showFilters}

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
