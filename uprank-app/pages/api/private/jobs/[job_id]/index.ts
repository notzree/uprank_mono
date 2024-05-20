import { NextApiRequest, NextApiResponse } from "next";
import prisma from "@/prisma/client";
import { getAuth } from "@clerk/nextjs/server";
import { Job } from "@prisma/client";
import enableCors from "@/utils/api_utils/enable_cors"
import { getIdOr401 } from "@/utils/api_utils/auth_utils";


async function handler(
    req: NextApiRequest,
    res: NextApiResponse
) {
    if (req.method === "GET") {
        await GET(req, res);
    } else {
        res.status(405).json({ message: "Method Not allowed" });
    }
}

async function GET(req: NextApiRequest, res: NextApiResponse) {
    try {
        const user_id = getIdOr401(req, res);
        let { job_id } = req.query;
        
        if (job_id instanceof Array || !job_id ) {
            res.status(400).json({ message: "Invalid Job ID" });
            return;
        }
        const result = await prisma.job.findUnique({
            where: {
                id: job_id,
                user_id: user_id as string,
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


