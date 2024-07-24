import * as React from "react";


import { Job } from "@/types/job";

import dayjs from "dayjs";
import utc from "dayjs/plugin/utc";

import localizedFormat from "dayjs/plugin/localizedFormat";
import {
    Card,
    CardContent,
} from "@/components/ui/card";
import { return_origin_job } from "@/utils/job_utils";
import { Badge } from "@/components/ui/badge"
import { Collapsible, CollapsibleTrigger, CollapsibleContent } from "@/components/ui/collapsible"
dayjs.extend(localizedFormat);
dayjs.extend(utc);

export default function JobCard({ job }: { job: Job }) {
    const [showFullDisplay, setShowFullDisplay] = React.useState(true);
    if (job == null){
        return (
            <div>
                No jobs found.
            </div>
        )
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
        <CardContent className="grid grid-cols-[1fr_auto] gap-6 p-4 sm:p-6">
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
                    {showFullDisplay ? full_description : truncated_description}
            </p>
            <div className="flex items-center gap-4 text-sm">
            </div>
            <Collapsible
                defaultOpen={true}
            >
              <CollapsibleContent>
                <div className="grid gap-4">
                  <div>
                    <h4 className="text-sm font-semibold">Skills</h4>
                    {og.skills?.map((skill, index) => (
                            <Badge key={skill} variant="outline" className="px-2 py-1 text-xs m-1">
                                {skill}
                            </Badge>
                        )
                    )}
                  </div>
                  <div>
                    <h4 className="text-sm font-semibold">Budget</h4>
                    <p className="text-lg font-semibold">{og.hourly ? "$" + og.hourly_rate + "/hr" : "$" + og.fixed_rate }</p>
                  </div>
                </div>
              </CollapsibleContent>
              <CollapsibleTrigger 
                onClick={() => setShowFullDisplay(!showFullDisplay)}
                className="font-semibold flex items-center gap-1 [&[data-state=open]>svg]:-rotate-90">
                    Expand
                <ChevronRightIcon className="w-4 h-4 transition-all translate-y-px" />
              </CollapsibleTrigger>
            </Collapsible>
          </div>
          <div className="flex flex-col items-end gap-4">
            {/* <Button variant="secondary" size="sm">
                Add Details
            </Button> */}
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


function BriefcaseIcon(props) {
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
      <path d="M16 20V4a2 2 0 0 0-2-2h-4a2 2 0 0 0-2 2v16" />
      <rect width="20" height="14" x="2" y="6" rx="2" />
    </svg>
  )
}


function CalendarIcon(props) {
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
  )
}


function ChevronRightIcon(props) {
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
  )
}


function LocateIcon(props) {
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
  )
}


function XIcon(props) {
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
      <path d="M18 6 6 18" />
      <path d="m6 6 12 12" />
    </svg>
  )
}