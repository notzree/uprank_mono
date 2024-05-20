// Description: creates multiple UpworkFreelancerProposal associated with a job
// Path: api/private/jobs/[job_id]/add_freelancers
import { NextApiRequest, NextApiResponse } from "next";
import prisma from "@/prisma/client";
import { Prisma, UpworkFreelancerProposal } from "@prisma/client";
import enableCors from "@/utils/api_utils/enable_cors";
import { Resource } from "sst";
import { SQSClient, SendMessageCommand } from "@aws-sdk/client-sqs";

import { Scraped_Freelancer_Data } from "@/types/freelancer";
import { Decimal } from "@prisma/client/runtime/library";
import { getIdOr401 } from "@/utils/api_utils/auth_utils";

const sqs = new SQSClient({});
async function handler(req: NextApiRequest, res: NextApiResponse) {
    if (req.method === "GET") {
        await GET(req, res);
    } else {
        res.status(405).json({ message: "Method Not allowed" });
    }
}

async function GET(req: NextApiRequest, res: NextApiResponse) {
    const user_id = getIdOr401(req, res);
    let { job_id } = req.query;
    await sqs.send(
        new SendMessageCommand({
            QueueUrl: Resource.ScrapeRequestQueue.url,
            MessageBody: JSON.stringify({
                job_id: job_id,
                user_id: user_id,
            }),
        })
    );
    res.status(200).json({ message: "success" });
}

export default enableCors(handler);

export interface Add_Freelancers_Result {
    result: Prisma.BatchPayload;
}
export interface Add_Freelancers_Response {
    freelancer_data: Scraped_Freelancer_Data[];
}
