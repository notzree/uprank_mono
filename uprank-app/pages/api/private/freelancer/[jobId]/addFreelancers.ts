//creates multiple UpworkFreelancerProposal associated with a job

import { NextApiRequest, NextApiResponse } from "next";
import prisma from "@/prisma/client";
import { getAuth } from "@clerk/nextjs/server";
import { Job, Prisma, UpworkFreelancerProposal } from "@prisma/client";
import enableCors from "@/utils/api_utils/enable_cors";
import { ScrapedFreelancerData, SendFreelancerBody } from "@/types/freelancer";

async function handler(req: NextApiRequest, res: NextApiResponse) {
    if (req.method === "POST") {
        await handlePost(req, res);
    } else {
        res.status(405).json({ message: "Method Not allowed" });
    }
}

async function handlePost(req: NextApiRequest, res: NextApiResponse) {
    try {
        const { userId } = getAuth(req);
        if (!userId && process.env.NODE_ENV === "production") {
            console.log("User id", userId);
            res.status(401).json({ message: "User not authenticated" });
        }
        let { jobId } = req.query;
        try {
        } catch (error) {
            res.status(500).json({
                message: `Invalid Job ID`,
                error: error,
            });
        }

        if (jobId instanceof Array || !jobId) {
            res.status(400).json({ message: "Invalid Job ID" });
            return;
        }
        const job = await prisma.job.findUnique({
            where: {
                id: jobId,
            },
        });
        const isHourly = job?.hourly;
        const isFixed = job?.fixed;
        if (!job) {
            res.status(404).json({ message: "Invalid Job ID" });
            return;
        }
        //first verify if jobId is valid.
        const body: ScrapedFreelancerData[] = req.body;
        const UpworkFreelancerProposals: UpworkFreelancerProposal[] = body.map(
            (freelancer) => {
                return {
                    jobId: jobId,
                    url: freelancer.url,
                    name: freelancer.name,
                    title: freelancer.title,
                    description: freelancer.description,
                    city: freelancer.location.city,
                    country: freelancer.location.country,
                    timezone: freelancer.location.timezone,
                    cv: freelancer.cv,
                    aiReccomended: freelancer.aiReccomended,
                    fixedChargeAmount: isFixed ? 
                    parseFloat(freelancer.fixedChargeAmount) : null,
                    fixedChargeCurrency: freelancer.fixedChargeCurrency,
                    hourlyChargeAmount: isHourly ? parseFloat(
                        freelancer.hourlyChargeAmount
                    ) : null,
                    hourlyChargeCurrency: freelancer.hourlyChargeCurrency,
                    invited: freelancer.invited,
                    photoUrl: freelancer.photoUrl,
                    recentHours: freelancer.recentHours,
                    totalHours: freelancer.totalHours,
                    totalPortfolioItems: freelancer.totalPortfolioItems,
                    totalPortfolioV2Items: freelancer.totalPortfolioV2Items,
                    upwork_totalFeedback: freelancer.totalFeedback,
                    upwork_recentFeedback: freelancer.recentFeedback,
                    upwork_topRatedStatus: freelancer.topRatedStatus,
                    upwork_topRatedPlusStatus: freelancer.topRatedPlusStatus,
                    upwork_sponsored: freelancer.sponsored,
                    upwork_jobSuccessScore: freelancer.jobSuccessScore,
                    upwork_reccomended: freelancer.reccomended,
                    skills: freelancer.skills,
                    averageRecentEarnings:
                        freelancer.earningsInfo.averageRecentEarnings,
                    combinedAverageRecentEarnings:
                        freelancer.earningsInfo.combinedAverageRecentEarnings,
                    combinedRecentEarnings:
                        freelancer.earningsInfo.combinedRecentEarnings,
                    combinedTotalEarnings:
                        freelancer.earningsInfo.combinedTotalEarnings,
                    combinedTotalRevenue:
                        freelancer.earningsInfo.combinedTotalRevenue,
                    recentEarnings: freelancer.earningsInfo.recentEarnings,
                    totalRevenue: freelancer.earningsInfo.totalRevenue,
                    uprank_score: 0,
                    uprank_updated_at: null,
                    uprank_reccomended: false,
                    uprank_reccomended_reasons: null,
                    uprank_not_enough_data: false,
                };
            }
        );

        const result = await prisma.upworkFreelancerProposal.createMany({
            data: UpworkFreelancerProposals,
            skipDuplicates: true, //so this can be called multiple times without creating duplicates
        },
);
        if (result.count ==body.length){
            console.log("All freelancers created successfully")
            res.status(200).json({ message: "All freelancers created successfully" });
        }

        res.status(200).json({ message: result});
    } catch (error) {
        console.log(error);
        res.status(500).json({ message: `something went wrong`, error: error });
    }
}

export default enableCors(handler);

export interface CreateManyFreelancersResult {
    exists: boolean;
    job: Job | null;
}
