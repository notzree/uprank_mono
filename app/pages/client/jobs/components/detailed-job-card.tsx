import * as React from "react";

import { Job } from "@/types/job";
import {
    TooltipProvider,
    Tooltip,
    TooltipTrigger,
    TooltipContent,
} from "@/components/ui/tooltip";
import dayjs from "dayjs";
import utc from "dayjs/plugin/utc";
import { Button } from "@/components/ui/button";
import localizedFormat from "dayjs/plugin/localizedFormat";
import { Card, CardContent } from "@/components/ui/card";
import { return_origin_job } from "@/utils/job_utils";
import { Badge } from "@/components/ui/badge";
import {
    Collapsible,
    CollapsibleTrigger,
    CollapsibleContent,
} from "@/components/ui/collapsible";
dayjs.extend(localizedFormat);
dayjs.extend(utc);

export default function DetailedJobCard({
    job,
    average_specialization_score,
    average_budget_adherence_percentage,
    average_budget_overrun_percentage,
    defaultOpen = true,
}: {
    job: Job;
    average_specialization_score: number;
    average_budget_adherence_percentage: number;
    average_budget_overrun_percentage: number;
    defaultOpen?: boolean;
}) {
    const [showFullDisplay, setShowFullDisplay] = React.useState(defaultOpen);
    if (job == null) {
        return <div>No jobs found.</div>;
    }
    const og = return_origin_job(job);
    let display_title = og.title || "No title found";
    let full_description = og.description || "No description found";
    let truncated_description = "No description found";
    if (og.description) {
        truncated_description = truncate_description(og.description);
    }

    let created_at = og.created_at || "No date found";
    if (og.created_at) {
        created_at = formate_date(og.created_at);
    }

    return (
        <Card className="w-full max-w-3xl rounded-lg overflow-hidden">
            <CardContent className="grid grid-cols-2 gap-6 p-4 sm:p-6">
                <div className="grid gap-2">
                    <div className="flex items-center gap-2">
                        <h3 className="text-lg font-semibold">
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
                        </h3>
                    </div>
                    <p className="text-sm text-muted-foreground">
                        {showFullDisplay
                            ? full_description
                            : truncated_description}
                    </p>
                    <div className="flex items-center gap-4 text-sm"></div>
                    <Collapsible defaultOpen={defaultOpen}>
                        <CollapsibleContent>
                            <div className="grid gap-4">
                                <div>
                                    <h4 className="text-sm font-semibold">
                                        Skills
                                    </h4>
                                    {og.skills?.map((skill, index) => (
                                        <Badge
                                            key={skill}
                                            variant="outline"
                                            className="px-2 py-1 text-xs m-1"
                                        >
                                            {skill}
                                        </Badge>
                                    ))}
                                </div>
                                <div>
                                    <h4 className="text-sm font-semibold">
                                        Budget
                                    </h4>
                                    <p className="text-lg font-semibold">
                                        {og.hourly
                                            ? "$" + og.hourly_rate + "/hr"
                                            : "$" + og.fixed_rate}
                                    </p>
                                </div>
                            </div>
                        </CollapsibleContent>
                        <CollapsibleTrigger
                            onClick={() => setShowFullDisplay(!showFullDisplay)}
                            className="font-semibold flex items-center gap-1 [&[data-state=open]>svg]:-rotate-90"
                        >
                            Expand
                            <ChevronRightIcon className="w-4 h-4 transition-all translate-y-px" />
                        </CollapsibleTrigger>
                    </Collapsible>
                </div>
                <div className="flex flex-col items-end gap-4">
                <div className="flex flex-col justify-between space-y-4">
                <h3 className="text-lg font-semibold">
                    Uprank Metrics
                </h3>
                        <div className="flex justify-between items-center">
                            <h4 className=" text-md text-left">
                                Average Specialization Score:
                            </h4>
                            <div className="flex flex-row items-center">
                                <p className="">
                                {average_specialization_score.toFixed(2)}
                            </p>
                            <TooltipProvider>
                                <Tooltip delayDuration={300}>
                                    <TooltipTrigger asChild>
                                        <Button variant="ghost" size="icon">
                                            <svg
                                                xmlns="http://www.w3.org/2000/svg"
                                                width="16"
                                                height="16"
                                                viewBox="0 0 24 24"
                                                fill="none"
                                                stroke="currentColor"
                                                strokeWidth="2"
                                                strokeLinecap="round"
                                                strokeLinejoin="round"
                                                className="h-4 w-4"
                                            >
                                                <circle
                                                    cx="12"
                                                    cy="12"
                                                    r="10"
                                                />
                                                <path d="M9.09 9a3 3 0 0 1 5.83 1c0 2-3 3-3 3" />
                                                <line
                                                    x1="12"
                                                    y1="17"
                                                    x2="12.01"
                                                    y2="17"
                                                />
                                            </svg>
                                        </Button>
                                    </TooltipTrigger>
                                    <TooltipContent className="max-w-md mb-2">
                                        <div>
                                            <div className="font-bold">
                                                What is this?
                                            </div>
                                        </div>
                                        <div className="max-w-md break-words overflow-hidden text-md">
                                            The specialization score shows how
                                            well a freelancer&apos;s past work
                                            matches the job you&apos;re
                                            offering.
                                            <span className="font-semibold">
                                                {" "}
                                                A higher score means they have
                                                more relevant experience and
                                                skills for your project.{" "}
                                            </span>
                                        </div>
                                    </TooltipContent>
                                </Tooltip>
                            </TooltipProvider>
                            </div>
                            
                        </div>
                        <div className="flex justify-between items-center">
                            <h4 className="text-md">
                                Average Budget Adherence Percentage:
                            </h4>
                            <p>
                            {average_budget_adherence_percentage.toFixed(0) +
                                "%"}
                            </p>
                            <TooltipProvider>
                                <Tooltip delayDuration={300}>
                                    <TooltipTrigger asChild>
                                        <Button variant="ghost" size="icon">
                                            <svg
                                                xmlns="http://www.w3.org/2000/svg"
                                                width="16"
                                                height="16"
                                                viewBox="0 0 24 24"
                                                fill="none"
                                                stroke="currentColor"
                                                strokeWidth="2"
                                                strokeLinecap="round"
                                                strokeLinejoin="round"
                                                className="h-4 w-4"
                                            >
                                                <circle
                                                    cx="12"
                                                    cy="12"
                                                    r="10"
                                                />
                                                <path d="M9.09 9a3 3 0 0 1 5.83 1c0 2-3 3-3 3" />
                                                <line
                                                    x1="12"
                                                    y1="17"
                                                    x2="12.01"
                                                    y2="17"
                                                />
                                            </svg>
                                        </Button>
                                    </TooltipTrigger>
                                    <TooltipContent className="max-w-md mb-2">
                                        <div>
                                            <div className="font-bold">
                                                What is this?
                                            </div>
                                        </div>
                                        <div className="max-w-md break-words overflow-hidden text-md">
                                            The budget adherence percentage
                                            indicates how often a freelancer
                                            stays within the agreed budget for
                                            their projects.{" "}
                                            <span className=" font-semibold">
                                                A higher percentage means they
                                                are more likely to keep costs
                                                under control
                                            </span>
                                        </div>
                                    </TooltipContent>
                                </Tooltip>
                            </TooltipProvider>
                        </div>
                        <div className="flex justify-start items-center">
                            <h4 className=" text-md">
                                Average Budget Overrun Percentage:
                            </h4>
                            {average_budget_overrun_percentage.toFixed(0) + "%"}
                            <TooltipProvider>
                                <Tooltip delayDuration={300}>
                                    <TooltipTrigger asChild>
                                        <Button variant="ghost" size="icon">
                                            <svg
                                                xmlns="http://www.w3.org/2000/svg"
                                                width="16"
                                                height="16"
                                                viewBox="0 0 24 24"
                                                fill="none"
                                                stroke="currentColor"
                                                strokeWidth="2"
                                                strokeLinecap="round"
                                                strokeLinejoin="round"
                                                className="h-4 w-4"
                                            >
                                                <circle
                                                    cx="12"
                                                    cy="12"
                                                    r="10"
                                                />
                                                <path d="M9.09 9a3 3 0 0 1 5.83 1c0 2-3 3-3 3" />
                                                <line
                                                    x1="12"
                                                    y1="17"
                                                    x2="12.01"
                                                    y2="17"
                                                />
                                            </svg>
                                        </Button>
                                    </TooltipTrigger>
                                    <TooltipContent className="max-w-md mb-2">
                                        <div>
                                            <div className="font-bold">
                                                What is this?
                                            </div>
                                        </div>
                                        <div className="max-w-md break-words overflow-hidden text-md">
                                            The budget overrun percentage shows
                                            how much a freelancer typically
                                            exceeds the agreed budget. A higher
                                            percentage means they are more
                                            likely to go over the budge by a
                                            larger amount.
                                            <span className="font-semibold">
                                                {" "}
                                                If this value is high, Expect to
                                                pay more than you agreed to.
                                            </span>
                                        </div>
                                    </TooltipContent>
                                </Tooltip>
                            </TooltipProvider>
                        </div>
                    </div>
                    <div className="flex items-center gap-1 text-sm text-muted-foreground">
                        <CalendarIcon className="w-4 h-4" />
                        <span>Created: {created_at.toLocaleString()}</span>
                    </div>
                    <div className="flex items-center gap-1 text-sm text-muted-foreground">
                        <LocateIcon className="w-4 h-4 text-muted-foreground" />
                        <span>Original Platform: {job?.origin_platform}</span>
                    </div>
                </div>
            </CardContent>
        </Card>
    );
}

function truncate_description(description: string): string {
    return description.slice(0, 100).concat("...");
}

function formate_date(date: Date): string {
    return dayjs.utc(date).local().format("dddd, MMMM D, YYYY");
}

function CalendarIcon(props: any) {
    return (
        <svg
            {...props}
            xmlns="http://www.w3.org/2000/svg"
            width="24"
            height="24"
            viewBox="0 0 24 24"
            fill="none"
            stroke="currentColor"
            strokeWidth="2"
            strokeLinecap="round"
            strokeLinejoin="round"
        >
            <path d="M8 2v4" />
            <path d="M16 2v4" />
            <rect width="18" height="18" x="3" y="4" rx="2" />
            <path d="M3 10h18" />
        </svg>
    );
}

function ChevronRightIcon(props: any) {
    return (
        <svg
            {...props}
            xmlns="http://www.w3.org/2000/svg"
            width="24"
            height="24"
            viewBox="0 0 24 24"
            fill="none"
            stroke="currentColor"
            strokeWidth="2"
            strokeLinecap="round"
            strokeLinejoin="round"
        >
            <path d="m9 18 6-6-6-6" />
        </svg>
    );
}

function LocateIcon(props: any) {
    return (
        <svg
            {...props}
            xmlns="http://www.w3.org/2000/svg"
            width="24"
            height="24"
            viewBox="0 0 24 24"
            fill="none"
            stroke="currentColor"
            strokeWidth="2"
            strokeLinecap="round"
            strokeLinejoin="round"
        >
            <line x1="2" x2="5" y1="12" y2="12" />
            <line x1="19" x2="22" y1="12" y2="12" />
            <line x1="12" x2="12" y1="2" y2="5" />
            <line x1="12" x2="12" y1="19" y2="22" />
            <circle cx="12" cy="12" r="7" />
        </svg>
    );
}
