import Navbar from "../shared/components/navbar";
import { JobDataTable } from "./components/job_table";

export default function Jobs({ job_prop }: { job_prop: Job }) {
    const job = job_prop.edges.upworkjob;
    if (!job) {
        return <div>Job not found</div>;
    }
    return (
        <div className="grid min-h-screen w-full">
            <div className="flex flex-col">
                <Navbar />
                <main className="flex flex-1 flex-col pt-16">
                    <div className="flex flex-row w-full  px-6">
                        <JobDataTable
                            freelancers={
                                job_prop.edges.upworkjob?.edges
                                    .upworkfreelancer || []
                            }
                            job={job}
                        />
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

export const getServerSideProps: GetServerSideProps = async (ctx) => {
    const base_url = process.env.NEXT_PUBLIC_BACKEND_DEV_BASE_URL;
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

    const client = new v1Client(base_url);
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
