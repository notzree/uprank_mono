//Scrapes Freelancer data givnen a job_id in the body
//post request because it is saving data in the db.

import { NextApiRequest, NextApiResponse } from "next";
import prisma from "@/prisma/client";
import { getIdOr401 } from "@/utils/api_utils/auth_utils";


export default async function handler(request: any) {
    console.log(request);


}

async function POST(req:NextApiRequest, res: NextApiResponse) {
    const user_id = getIdOr401(req, res);
    let { job_id } = req.query;
        if (job_id instanceof Array || !job_id ) {
            res.status(400).json({ message: "Invalid Job ID" });
            return;
        }
        try {
            const job = await prisma.job.findUnique({
                where:{
                    id: job_id,
                    user_id: user_id as string
                },
                include:{
                    Freelancers: true
                }
            })
            console.log(job);

        }
         catch (error){
            console.log(error);
            res.status(500).json({ message: `something went wrong`, error: error });
         }

    

}
