// Description: This file is responsible for creating a new job in the database.
// Path: api/private/jobs

import { NextApiRequest, NextApiResponse } from "next";
import prisma from "@/prisma/client";
import { getAuth } from "@clerk/nextjs/server";
import { PrismaClientKnownRequestError } from "@prisma/client/runtime/library";
import { Scraped_Job_Data } from "@/types/job";
import enableCors from "@/utils/api_utils/enable_cors"

 async function handler(
    req: NextApiRequest,
    res: NextApiResponse
) {
    console.log("Request method", req.method);
    if (req.method === "POST") {

        await POST(req, res);
    } else {
        res.status(405).json({ message: "Method Not allowed" });
    }
}

async function POST(req: NextApiRequest, res: NextApiResponse) {
    try {
        const { userId } = getAuth(req);
        if (!userId && process.env.NODE_ENV === "production") {
            res.status(401).json({ message: "User not authenticated" });
        }

        const body: Post_Job_Request = req.body;
        const job = body.job;

        const result = await prisma.job.create({
            data: {
                id: job.id,
                title: job.title,
                location: job.location,
                description: job.description,
                skills: job.skills,
                experience_level: job.experience_level,
                hourly: job.hourly,
                fixed: job.fixed,
                hourly_rate: job.hourly_rate,
                fixed_rate: job.fixed_rate,
                user_id: userId as string,
            },
        });
        res.status(200).json({result} as Post_Job_Response);
    } catch (error) {
            console.log(error);
            res.status(500).json({ message: `${error}` });
    }
}

export default enableCors(handler);


export interface Post_Job_Request {
    job: Scraped_Job_Data;
}

export interface Post_Job_Response { //type this shit later 
    result: any

}