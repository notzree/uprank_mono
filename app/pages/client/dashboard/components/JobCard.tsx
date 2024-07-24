import { Job } from "@/types/job";
import Link from "next/link";
import dayjs from "dayjs";
import utc from "dayjs/plugin/utc";
import { Button } from "@/components/ui/button";
import localizedFormat from "dayjs/plugin/localizedFormat";
import {
    Card,
    CardContent,
    CardDescription,
    CardFooter,
    CardHeader,
    CardTitle,
} from "@/components/ui/card";
import { return_origin_job } from "@/utils/job_utils";
import { BackendClient } from "@/backend-client/backend-client";
import { useUser } from "@clerk/nextjs";
import { useAuth } from "@clerk/nextjs";
dayjs.extend(localizedFormat);
dayjs.extend(utc);

export default function JobCard({ job }: { job: Job }) {
    const { user } = useUser();
    const {getToken} = useAuth();
    if (job == null){
        return (
            <div>
                No jobs found.
            </div>
        )
    }
    const client = new BackendClient();
    const og = return_origin_job(job);
    let display_title = og.title || "No title found";
    let display_description = og.description || "No description found";
    if (og.description) {
        display_description = truncate_description(og.description);
    }
    let created_at = og.created_at || "No date found";
    if (og.created_at) {
        created_at = formate_date(og.created_at);
    }
    let action_text = "";
    let action_link = "";
    if (og.ranked_at != null) {
        action_text = "View Ranking";
        action_link = "jobs/" + job.id;
    } else {
        action_text = "Rank";
        action_link = "";
    }


    const Rank = async ()=>{
        const token = await getToken();
        if (!user){
            return
        }
        if (job.id == null){
            return
        }
        if (job.edges.upworkjob?.id == null){
            return
        }
        if (job?.origin_platform == null){
            return
        }
        if (token == null){
            return
        }
        await client.Rank({
            job_id: job.id,
            user_id: user?.id,
            platform_id: job.edges.upworkjob?.id,
            platform: job?.origin_platform,
    }, token)}
    
    
    return (
        <Card>
            <CardHeader>
                <CardTitle>
                    <a
                        href={
                            "https://www.upwork.com/ab/applicants/" +
                            og.id +
                            "/applicants"
                        }
                        target="_blank"
                        className=" underline"
                    >
                        {display_title}
                    </a>
                </CardTitle>
                <CardDescription className="flex flex-row gap-x-2">
                    <span> Original Platform: {job?.origin_platform}</span> |
                    <span>Created : {created_at.toLocaleString()}</span>
                </CardDescription>
            </CardHeader>
            <CardContent>
                <p>{display_description}</p>
            </CardContent>
            <CardFooter>
            {og.ranked_at != null ? (
            <Link href={action_link}>{action_text}</Link>
        ) : (
            <Button onClick={()=>Rank()}>{action_text}</Button>
        )}
            </CardFooter>
        </Card>
    );
}

function truncate_description(description: string): string {
    return description.slice(0, 100).concat("...");
}

function formate_date(date: Date): string {
    return dayjs.utc(date).local().format("dddd, MMMM D, YYYY");
}
