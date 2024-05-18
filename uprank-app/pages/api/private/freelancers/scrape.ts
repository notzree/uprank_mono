//Scrapes Freelancer data givnen a job_id in the body
//post request because it is saving data in the db.

import { NextApiRequest, NextApiResponse } from "next";
import prisma from "@/prisma/client";
import { getIdOr401 } from "@/utils/api_utils/auth_utils";


export default async function handler(req:NextApiRequest, res: NextApiResponse) {
    if (req.method !== "POST"){
        res.setHeader('Allow', ['POST']);
        res.status(405).end(`Method ${req.method} Not Allowed`);
    }
    else {
       await handlePost(req, res);
    }

}

async function handlePost(req:NextApiRequest, res: NextApiResponse) {
    const user_id = getIdOr401(req, res);
    const job_id = req.body.job_id;


}