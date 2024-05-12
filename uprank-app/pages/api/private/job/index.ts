import { NextApiRequest, NextApiResponse } from "next";
import prisma from "@/prisma/client";
import { getAuth } from "@clerk/nextjs/server";
import { PrismaClientKnownRequestError } from "@prisma/client/runtime/library";
import { ScrapedJobData } from "@/types/job";


export default async function handler(
    req: NextApiRequest,
    res: NextApiResponse
) {
    if (req.method === "POST") {
        await handlePost(req, res);
    } else {
        res.status(500).json({ message: "Invalid request type" });
    }
}

async function handlePost(req: NextApiRequest, res: NextApiResponse) {
    try {
        const { userId } = getAuth(req);

        if (!userId && process.env.NODE_ENV === "production") {
            res.status(401).json({ message: "User not authenticated" });
        }
        const job = req.body as ScrapedJobData;

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
                num_months: job.num_months,
                user_id: userId as string,
            },
        });
        res.status(200).json({ message: result });
    } catch (error) {
        if (error instanceof PrismaClientKnownRequestError)
            res.status(500).json({ message: "Job already exsits" });
        else {
            res.status(500).json({ message: `${error}` });
        }
    }
}
