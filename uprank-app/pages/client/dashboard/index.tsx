import * as React from "react";
import { GetJobResponse } from "@/client/client";
import Navbar from "./components/navbar";
import JobCardList from "./components/JobCardList";
export default function Dashboard({ jobs_props }: {jobs_props: GetJobResponse}) {
    const [jobs, setJobs] = React.useState(jobs_props.jobs);
    return (
        <div className="grid min-h-screen w-full">
            <div className="flex flex-col">
                <Navbar />
                <main className="flex flex-1 flex-col gap-4 p-4 md:gap-8 md:p-6">
                    <div className="flex flex-row items-center justify-center w-full text-center">
                        <div className=" flex flex-col ">
                            <h2>Your current job postings</h2>
                            <JobSearchFilter jobs={jobs_props.jobs} setJob={setJobs} />
                            <div className="flex-1">
                            <JobCardList jobs={jobs} />
                            </div>
                        </div>
                    </div>
                </main>
            </div>
        </div>
    );
}



import { getAuth } from "@clerk/nextjs/server";
import { GetServerSideProps } from "next";
import { v1Client } from "@/client/v1_client";
import { Job } from "@/types/job";
import JobCard from "./components/JobCard";
import JobSearchFilter from "./components/JobSearchFilter";


export const getServerSideProps: GetServerSideProps = async (ctx) => {
    const base_url = process.env.NEXT_PUBLIC_BACKEND_DEV_BASE_URL;
    const { userId, getToken } = getAuth(ctx.req);
    const token = await getToken();
    if (!userId || !token) {
        return {
            redirect: {
                destination: "/",
                permanent: false,
            },
        };
    }
    const client = new v1Client(base_url);
    const jobs = await client.GetAllJobs(token);
    return {
        props: {
            jobs_props: jobs,
        },
    };
}
