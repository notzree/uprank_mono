import * as React from "react";
import { GetJobResponse } from "@/backend-client/backend-client";
import Navbar from "../shared/components/navbar";
import JobCardList from "./components/JobCardList";
import JobSearchFilter from "./components/JobSearchFilter";
export default function Dashboard({ job_props }: {job_props: Job[]}) {
    const [jobs, setJobs] = React.useState(job_props);
    return (
        <div className="grid min-h-screen w-full">
            <div className="flex flex-col">
                <Navbar />
                <main className="flex flex-1 flex-col pt-16">
                    <div className="flex flex-row w-full ">
                        <div className=" flex w-96 px-4">
                        <JobSearchFilter jobs={job_props} setJob={setJobs} />
                        </div>
                        <div className=" flex flex-col items-center justify-center w-full text-left">
                            <h1 className=" text-2xl font-bold pb-4">Your current job postings</h1>
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
import { BackendClient } from "@/backend-client/backend-client";
import { Job } from "@/types/job";

export const getServerSideProps: GetServerSideProps = async (ctx) => {
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
    const client = new BackendClient();
    const jobs = await client.GetAllJobs(token);
    return {
        props: {
            job_props: jobs,
        },
    };
}
