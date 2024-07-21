import { Job } from "@/types/job";
import { Input } from "@/components/ui/input";
import {
    Accordion,
    AccordionItem,
    AccordionTrigger,
    AccordionContent,
} from "@/components/ui/accordion";
import { Label } from "@/components/ui/label";
import { Checkbox } from "@/components/ui/checkbox";
import * as React from "react";
import Fuse from "fuse.js";
import { return_origin_job } from "@/utils/job_utils";
export default function JobSearchFilter({
    jobs,
    setJob,
}: {
    jobs: Job[];
    setJob: React.Dispatch<React.SetStateAction<Job[]>>;
}) {
    const [searchTerm, setSearchTerm] = React.useState("");
    const origin_title_map = new Map<string, string>();

    const fuse = new Fuse(jobs, {
        keys: [
            {
                name: "title",
                getFn: (job) => return_origin_job(job).title || "undefined",
            },
        ],
        includeScore: true,
        threshold: 0.4, // Adjust the threshold for more or less fuzzy matching
    });

    const filterJobs = (term: string) => {
        if (term.trim() === "") {
            setJob(jobs); // Show all jobs if search term is empty
        } else {
            const result = fuse.search(term);
            const filteredJobs = result.map((res) => res.item);
            setJob(filteredJobs);
        }
    };

    const handleInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        const term = e.target.value;
        setSearchTerm(term);
        filterJobs(term);
    };

    return (
        <div className="flex flex-col h-full bg-background border-r border-l border-t ">
            <div className="p-4 border-b">
                <h2 className="text-lg font-semibold">Filters</h2>
            </div>
            <div className="p-4">
                <Input
                    type="text"
                    placeholder="Job Title"
                    value={searchTerm}
                    onChange={handleInputChange}
                />
                <Accordion type="single" collapsible>
                    <AccordionItem value="categories">
                        <AccordionTrigger className="text-base font-medium">
                            Categories
                        </AccordionTrigger>
                        <AccordionContent>
                            <div className="grid gap-2">
                                <Label className="flex items-center gap-2 font-normal">
                                    <Checkbox id="category-electronics" />{" "}
                                    Electronics
                                </Label>
                                <Label className="flex items-center gap-2 font-normal">
                                    <Checkbox id="category-clothing" /> Clothing
                                </Label>
                                <Label className="flex items-center gap-2 font-normal">
                                    <Checkbox id="category-home" /> Home
                                </Label>
                                <Label className="flex items-center gap-2 font-normal">
                                    <Checkbox id="category-beauty" /> Beauty
                                </Label>
                            </div>
                        </AccordionContent>
                    </AccordionItem>
                    <AccordionItem value="tags">
                        <AccordionTrigger className="text-base font-medium">
                            Tags
                        </AccordionTrigger>
                        <AccordionContent>
                            <div className="grid gap-2">
                                <Label className="flex items-center gap-2 font-normal">
                                    <Checkbox id="tag-new" /> New
                                </Label>
                                <Label className="flex items-center gap-2 font-normal">
                                    <Checkbox id="tag-sale" /> Sale
                                </Label>
                                <Label className="flex items-center gap-2 font-normal">
                                    <Checkbox id="tag-bestseller" /> Bestseller
                                </Label>
                                <Label className="flex items-center gap-2 font-normal">
                                    <Checkbox id="tag-featured" /> Featured
                                </Label>
                            </div>
                        </AccordionContent>
                    </AccordionItem>
                    <AccordionItem value="price">
                        <AccordionTrigger className="text-base font-medium">
                            Price Range
                        </AccordionTrigger>
                        <AccordionContent>
                            <div className="grid gap-2">
                                <Label className="flex items-center gap-2 font-normal">
                                    <Checkbox id="price-under50" /> Under $50
                                </Label>
                                <Label className="flex items-center gap-2 font-normal">
                                    <Checkbox id="price-50to100" /> $50 - $100
                                </Label>
                                <Label className="flex items-center gap-2 font-normal">
                                    <Checkbox id="price-100to200" /> $100 - $200
                                </Label>
                                <Label className="flex items-center gap-2 font-normal">
                                    <Checkbox id="price-over200" /> Over $200
                                </Label>
                            </div>
                        </AccordionContent>
                    </AccordionItem>
                </Accordion>
            </div>
        </div>
    );
}
