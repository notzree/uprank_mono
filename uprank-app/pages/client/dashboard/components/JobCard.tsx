import { Job, UpworkJob } from "@/types/job";
import {
    Card,
    CardContent,
    CardDescription,
    CardFooter,
    CardHeader,
    CardTitle,
} from "@/components/ui/card";
import { return_origin_job } from "@/utils/job_utils";
export default function JobCard({job}: {job: Job}) {
    const og = return_origin_job(job);
    let display_title = og.title || "No title found";

    let display_description = og.description || "No description found";
    if (og.description) {
        display_description = truncate_description(og.description);
    }

    return (
        <Card>
            <CardHeader>
                <CardTitle>{display_title}</CardTitle>
                <CardDescription>created on: {job.origin_platform}</CardDescription>
            </CardHeader>
            <CardContent>
                <p>{display_description}</p>
            </CardContent>
            <CardFooter>
                <p>idk what to put here heehee</p>
            </CardFooter>
        </Card>
    );
}



function truncate_description(description: string): string {
    return description.slice(0, 100).concat("...");
}
