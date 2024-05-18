import { NextApiRequest, NextApiResponse } from "next";
import prisma from "@/prisma/client";
import { getAuth } from "@clerk/nextjs/server";
import { Job } from "@prisma/client";
import enableCors from "@/utils/api_utils/enable_cors"


async function handler(
    req: NextApiRequest,
    res: NextApiResponse
) {
    if (req.method === "GET") {
        await handleGet(req, res);
    } else {
        res.status(405).json({ message: "Method Not allowed" });
    }
}

async function handleGet(req: NextApiRequest, res: NextApiResponse) {
    try {
        const { userId } = getAuth(req)
        if (!userId && process.env.NODE_ENV === "production") {
            console.log("User id" ,userId);
            res.status(401).json({ message: "User not authenticated" });
        }
        let { job_id } = req.query;
        try {
        } catch (error) {
            res.status(500).json({
                message: `Invalid Job ID`,
                error: error,
            });
        }
        
        if (job_id instanceof Array) {
            res.status(400).json({ message: "Invalid Job ID" });
            return;
        }
        const result = await prisma.job.findUnique({
            where: {
                id: job_id,
                user_id: userId as string,
            },
            include:{
                _count:{
                    select:{
                        Freelancers: true
                    }
                }
            }
        });
        

        const responseBody: GetJobResult = {
            exists: result ? true : false,
            job: result,
        };
        
        res.status(200).json(responseBody);
    } catch (error) {
        console.log(error);
        res.status(500).json({ message: `something went wrong`, error: error });
    }
}


export default enableCors(handler);

export interface GetJobResult {
    exists: boolean,
    job: Job | null

}


