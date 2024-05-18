//creates multiple UpworkFreelancerProposal associated with a job

import { NextApiRequest, NextApiResponse } from "next";
import prisma from "@/prisma/client";
import { getAuth } from "@clerk/nextjs/server";
import { Job, UpworkFreelancerProposal } from "@prisma/client";
import enableCors from "@/utils/api_utils/enable_cors";
import { Scraped_Freelancer_Data, Send_Freelancer_Body } from "@/types/freelancer";
import { Decimal } from "@prisma/client/runtime/library";

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
        let { job_id } = req.query;
        try {
        } catch (error) {
            res.status(500).json({
                message: `Invalid Job ID`,
                error: error,
            });
        }

        if (job_id instanceof Array || !job_id) {
            res.status(400).json({ message: "Invalid Job ID" });
            return;
        }
        const job = await prisma.job.findUnique({
            where: {
                id: job_id,
            },
        });
        const is_hourly = job?.hourly;
        const is_fixed = job?.fixed;
        if (!job) {
            res.status(404).json({ message: "Invalid Job ID" });
            return;
        }
        //first verify if jobId is valid.
        const body: Scraped_Freelancer_Data[] = req.body;
        const upwork_freelancer_proposals: UpworkFreelancerProposal[] = body.map(
            (freelancer) => {
                return {
                    job_id: job_id,
                    url: freelancer.url,
                    name: freelancer.name,
                    title: freelancer.title,
                    description: freelancer.description,
                    city: freelancer.location.city,
                    country: freelancer.location.country,
                    timezone: freelancer.location.timezone,
                    cv: freelancer.cv,
                    ai_reccomended: freelancer.ai_reccomended,
                    fixed_charge_amount: is_fixed ? 
                        parseFloat(freelancer.fixed_charge_amount) : null,
                    fixed_charge_currency: freelancer.fixed_charge_currency,
                    hourly_charge_amount: is_hourly ? parseFloat(
                        freelancer.hourly_charge_amount
                    ) : null,
                    hourly_charge_currency: freelancer.hourly_charge_currency,
                    invited: freelancer.invited,
                    photo_url: freelancer.photo_url,
                    recent_hours: freelancer.recent_hours,
                    total_hours: freelancer.total_hours,
                    total_portfolio_items: freelancer.total_portfolio_items,
                    total_portfolio_v2_items: freelancer.total_portfolio_v2_items,
                    upwork_total_feedback: freelancer.total_feedback,
                    upwork_recent_feedback: freelancer.recent_feedback,
                    upwork_top_rated_status: freelancer.top_rated_status,
                    upwork_top_rated_plus_status: freelancer.top_rated_plus_status,
                    upwork_sponsored: freelancer.sponsored,
                    upwork_job_success_score: freelancer.job_success_score,
                    upwork_reccomended: freelancer.reccomended,
                    skills: freelancer.skills,
                    average_recent_earnings: freelancer.earnings_info.average_recent_earnings,
                    combined_average_recent_earnings: freelancer.earnings_info.combined_average_recent_earnings,
                    combined_recent_earnings: freelancer.earnings_info.combined_recent_earnings,
                    combined_total_earnings: freelancer.earnings_info.combined_total_earnings,
                    combined_total_revenue: freelancer.earnings_info.combined_total_revenue,
                    recent_earnings: freelancer.earnings_info.recent_earnings,
                    total_revenue: freelancer.earnings_info.total_revenue,
                    uprank_score: 0,
                    uprank_updated_at: null,
                    uprank_reccomended: false,
                    uprank_reccomended_reasons: null,
                    uprank_not_enough_data: false,
                };
            }
        );
        

        const result = await prisma.upworkFreelancerProposal.createMany({
            data: upwork_freelancer_proposals,
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
