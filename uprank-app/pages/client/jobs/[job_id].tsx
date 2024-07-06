//data table here with information about the job (just do upwork for now)


export default function Jobs({job_prop}:{job_prop: Job}){

    console.log(job_prop)
    return (
        <div>
            <h1>Jobs</h1>
        </div>
    )
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
}