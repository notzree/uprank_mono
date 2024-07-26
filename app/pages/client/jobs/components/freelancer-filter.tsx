import {UpworkJob } from "@/types/job";
import { Input } from "@/components/ui/input";
import {
    Accordion,
    AccordionItem,
    AccordionTrigger,
    AccordionContent,
} from "@/components/ui/accordion";

import * as React from "react";
import Fuse from "fuse.js";
import type { UpworkFreelancer } from "@/types/freelancer";

//Filter requirements:
// - All columns in the freelancer-table
// - Option to filter > or < the 50th percentile ?

export default function FreelancerSearchFilter({
    job,
    original_freelancers,
    visible_freelancers,
    setFreelancers,
}: {
    job: UpworkJob;
    original_freelancers: UpworkFreelancer[];
    visible_freelancers: UpworkFreelancer[];
    setFreelancers: React.Dispatch<React.SetStateAction<UpworkFreelancer[]>>;
}) {
    const [searchTerm, setSearchTerm] = React.useState("");
    const [filterData, setFilterData] = React.useState({
        min_proposed_rate: -1,
        max_proposed_rate: -1,
        min_recent_hours: -1,
        max_recent_hours: -1,
        min_total_hours: -1,
        max_total_hours: -1,
        min_average_recent_earnings: -1,
        max_average_recent_earnings: -1,
        min_total_earnings: -1,
        max_total_earnings: -1,
        min_specialization_score: -1,
        max_specialization_score: -1,
        min_budget_adherence_percentage: -1,
        max_budget_adherence_percentage: -1,
        min_budget_overrun_percentage: -1,
        max_budget_overrun_percentage: -1,
    });
    const fuse = new Fuse(visible_freelancers, {
        keys: [
            {
                name: "title",
                getFn: (freelancer) => freelancer.name,
            },
            {
                name: "skills",
                getFn: (freelancer) => freelancer.skills,
            },
            {
                name: "region",
                getFn: (freelancer) =>
                    freelancer.country + " " + freelancer.city,
            },
        ],
        includeScore: true,
        threshold: 0.4, // Adjust the threshold for more or less fuzzy matching
    });

    const filterFreelancers = (term: string) => {
        if (term.trim() === "") {
            setFreelancers(original_freelancers);
        } else {
            const result = fuse.search(term);
            const filteredFreelancers = result.map((res) => res.item);
            setFreelancers(filteredFreelancers);
        }
    };

    const handleSearchInputChange = (
        e: React.ChangeEvent<HTMLInputElement>
    ) => {
        const term = e.target.value;
        setSearchTerm(term);
        filterFreelancers(term);
    };

    React.useEffect(() => {
        const filterFreelancers = () => {
            const filtered = original_freelancers.filter((freelancer) => {
                const rate =
                    freelancer.fixed_charge_amount ??
                    freelancer.hourly_charge_amount;
                return (
                    (filterData.min_proposed_rate === -1 ||
                        rate >= filterData.min_proposed_rate) &&
                    (filterData.max_proposed_rate === -1 ||
                        rate <= filterData.max_proposed_rate) &&
                    (filterData.min_recent_hours === -1 ||
                        freelancer.recent_hours >=
                            filterData.min_recent_hours) &&
                    (filterData.max_recent_hours === -1 ||
                        freelancer.recent_hours <=
                            filterData.max_recent_hours) &&
                    (filterData.min_total_hours === -1 ||
                        freelancer.total_hours >= filterData.min_total_hours) &&
                    (filterData.max_total_hours === -1 ||
                        freelancer.total_hours <= filterData.max_total_hours) &&
                    (filterData.min_average_recent_earnings === -1 ||
                        freelancer.average_recent_earnings >=
                            filterData.min_average_recent_earnings) &&
                    (filterData.max_average_recent_earnings === -1 ||
                        freelancer.average_recent_earnings <=
                            filterData.max_average_recent_earnings) &&
                    (filterData.min_total_earnings === -1 ||
                        freelancer.combined_total_earnings >=
                            filterData.min_total_earnings) &&
                    (filterData.max_total_earnings === -1 ||
                        freelancer.combined_total_earnings <=
                            filterData.max_total_earnings) &&
                    (filterData.min_specialization_score === -1 ||
                        freelancer.edges.freelancer_inference_data
                            .finalized_rating_score >=
                            filterData.min_specialization_score) &&
                    (filterData.max_specialization_score === -1 ||
                        freelancer.edges.freelancer_inference_data
                            .finalized_rating_score <=
                            filterData.max_specialization_score) &&
                    (filterData.min_budget_adherence_percentage === -1 ||
                        freelancer.edges.freelancer_inference_data
                            .budget_adherence_percentage >=
                            filterData.min_budget_adherence_percentage) &&
                    (filterData.max_budget_adherence_percentage === -1 ||
                        freelancer.edges.freelancer_inference_data
                            .budget_adherence_percentage <=
                            filterData.max_budget_adherence_percentage) &&
                    (filterData.min_budget_overrun_percentage === -1 ||
                        freelancer.edges.freelancer_inference_data
                            .budget_overrun_percentage >=
                            filterData.min_budget_overrun_percentage) &&
                    (filterData.max_budget_overrun_percentage === -1 ||
                        freelancer.edges.freelancer_inference_data
                            .budget_overrun_percentage <=
                            filterData.max_budget_overrun_percentage)
                );
            });
            setFreelancers(filtered);
        };

        filterFreelancers();
    // eslint-disable-next-line react-hooks/exhaustive-deps
    }, [filterData]);

    return (
        <div className="flex flex-col h-full bg-background min-w-80 border-t">
            <div className="p-4 border-b">
                <h2 className="text-lg font-semibold">Filters</h2>
            </div>
            <div className="p-4">
                <Input
                    type="text"
                    placeholder="Search for anything..."
                    value={searchTerm}
                    onChange={handleSearchInputChange}
                />
                <Accordion type="single" collapsible>
                    <AccordionItem value="proposed_rate">
                        <AccordionTrigger className="text-base font-medium">
                            Proposed Rate
                        </AccordionTrigger>
                        <AccordionContent>
                            <div className="grid gap-2">
                                <Input
                                    type="number"
                                    placeholder="Min"
                                    onChange={(e) =>
                                        setFilterData({
                                            ...filterData,
                                            min_proposed_rate:
                                                e.target.value === ""
                                                    ? -1
                                                    : parseFloat(
                                                          e.target.value
                                                      ),
                                        })
                                    }
                                />
                                <Input
                                    type="number"
                                    placeholder="Max"
                                    onChange={(e) =>
                                        setFilterData({
                                            ...filterData,
                                            max_proposed_rate:
                                                e.target.value === ""
                                                    ? -1
                                                    : parseFloat(
                                                          e.target.value
                                                      ),
                                        })
                                    }
                                />
                            </div>
                        </AccordionContent>
                    </AccordionItem>
                    <AccordionItem value="recent_hours">
                        <AccordionTrigger className="text-base font-medium">
                            Recent hours
                        </AccordionTrigger>
                        <AccordionContent>
                            <div className="grid gap-2">
                                <Input
                                    type="number"
                                    placeholder="Min"
                                    onChange={(e) =>
                                        setFilterData({
                                            ...filterData,
                                            min_recent_hours:
                                                e.target.value === ""
                                                    ? -1
                                                    : parseFloat(
                                                          e.target.value
                                                      ),
                                        })
                                    }
                                />
                                <Input
                                    type="number"
                                    placeholder="Max"
                                    onChange={(e) =>
                                        setFilterData({
                                            ...filterData,
                                            max_recent_hours:
                                                e.target.value === ""
                                                    ? -1
                                                    : parseFloat(
                                                          e.target.value
                                                      ),
                                        })
                                    }
                                />
                            </div>
                        </AccordionContent>
                    </AccordionItem>
                    <AccordionItem value="total_hours">
                        <AccordionTrigger className="text-base font-medium">
                            Total hours
                        </AccordionTrigger>
                        <AccordionContent>
                            <div className="grid gap-2">
                                <Input
                                    type="number"
                                    placeholder="Min"
                                    onChange={(e) =>
                                        setFilterData({
                                            ...filterData,
                                            min_total_hours:
                                                e.target.value === ""
                                                    ? -1
                                                    : parseFloat(
                                                          e.target.value
                                                      ),
                                        })
                                    }
                                />
                                <Input
                                    type="number"
                                    placeholder="Max"
                                    onChange={(e) =>
                                        setFilterData({
                                            ...filterData,
                                            max_total_hours:
                                                e.target.value === ""
                                                    ? -1
                                                    : parseFloat(
                                                          e.target.value
                                                      ),
                                        })
                                    }
                                />
                            </div>
                        </AccordionContent>
                    </AccordionItem>
                    <AccordionItem value="average_recent_earnings">
                        <AccordionTrigger className="text-base font-medium">
                            Average recent earnings / job
                        </AccordionTrigger>
                        <AccordionContent>
                            <div className="grid gap-2">
                                <Input
                                    type="number"
                                    placeholder="Min"
                                    onChange={(e) =>
                                        setFilterData({
                                            ...filterData,
                                            min_average_recent_earnings:
                                                e.target.value === ""
                                                    ? -1
                                                    : parseFloat(
                                                          e.target.value
                                                      ),
                                        })
                                    }
                                />
                                <Input
                                    type="number"
                                    placeholder="Max"
                                    onChange={(e) =>
                                        setFilterData({
                                            ...filterData,
                                            max_average_recent_earnings:
                                                e.target.value === ""
                                                    ? -1
                                                    : parseFloat(
                                                          e.target.value
                                                      ),
                                        })
                                    }
                                />
                            </div>
                        </AccordionContent>
                    </AccordionItem>
                    <AccordionItem value="total_earnings">
                        <AccordionTrigger className="text-base font-medium">
                            Total earnings
                        </AccordionTrigger>
                        <AccordionContent>
                            <div className="grid gap-2">
                                <Input
                                    type="number"
                                    placeholder="Min"
                                    onChange={(e) =>
                                        setFilterData({
                                            ...filterData,
                                            min_total_earnings:
                                                e.target.value === ""
                                                    ? -1
                                                    : parseFloat(
                                                          e.target.value
                                                      ),
                                        })
                                    }
                                />
                                <Input
                                    type="number"
                                    placeholder="Max"
                                    onChange={(e) =>
                                        setFilterData({
                                            ...filterData,
                                            max_total_earnings:
                                                e.target.value === ""
                                                    ? -1
                                                    : parseFloat(
                                                          e.target.value
                                                      ),
                                        })
                                    }
                                />
                            </div>
                        </AccordionContent>
                    </AccordionItem>
                    <AccordionItem value="specialization_score">
                        <AccordionTrigger className="text-base font-medium">
                            Specialization Score
                        </AccordionTrigger>
                        <AccordionContent>
                            <div className="grid gap-2">
                                <Input
                                    type="number"
                                    placeholder="Min"
                                    onChange={(e) =>
                                        setFilterData({
                                            ...filterData,
                                            min_specialization_score:
                                                e.target.value === ""
                                                    ? -1
                                                    : parseFloat(
                                                          e.target.value
                                                      ),
                                        })
                                    }
                                />
                                <Input
                                    type="number"
                                    placeholder="Max"
                                    onChange={(e) =>
                                        setFilterData({
                                            ...filterData,
                                            max_specialization_score:
                                                e.target.value === ""
                                                    ? -1
                                                    : parseFloat(
                                                          e.target.value
                                                      ),
                                        })
                                    }
                                />
                            </div>
                        </AccordionContent>
                    </AccordionItem>
                    <AccordionItem value="budget_adherence_percentage">
                        <AccordionTrigger className="text-base font-medium">
                            Budget Adherence Percentage
                        </AccordionTrigger>
                        <AccordionContent>
                            <div className="grid gap-2">
                                <Input
                                    type="number"
                                    placeholder="Min"
                                    onChange={(e) =>
                                        setFilterData({
                                            ...filterData,
                                            min_budget_adherence_percentage:
                                                e.target.value === ""
                                                    ? -1
                                                    : parseFloat(
                                                          e.target.value
                                                      ),
                                        })
                                    }
                                />
                                <Input
                                    type="number"
                                    placeholder="Max"
                                    onChange={(e) =>
                                        setFilterData({
                                            ...filterData,
                                            max_budget_adherence_percentage:
                                                e.target.value === ""
                                                    ? -1
                                                    : parseFloat(
                                                          e.target.value
                                                      ),
                                        })
                                    }
                                />
                            </div>
                        </AccordionContent>
                    </AccordionItem>
                    <AccordionItem value="budget_overrun_percentage">
                        <AccordionTrigger className="text-base font-medium">
                            Budget overrun percentage
                        </AccordionTrigger>
                        <AccordionContent>
                            <div className="grid gap-2">
                                <Input
                                    type="number"
                                    placeholder="Min"
                                    onChange={(e) =>
                                        setFilterData({
                                            ...filterData,
                                            min_budget_overrun_percentage:
                                                e.target.value === ""
                                                    ? -1
                                                    : parseFloat(
                                                          e.target.value
                                                      ),
                                        })
                                    }
                                />
                                <Input
                                    type="number"
                                    placeholder="Max"
                                    onChange={(e) =>
                                        setFilterData({
                                            ...filterData,
                                            max_budget_overrun_percentage:
                                                e.target.value === ""
                                                    ? -1
                                                    : parseFloat(
                                                          e.target.value
                                                      ),
                                        })
                                    }
                                />
                            </div>
                        </AccordionContent>
                    </AccordionItem>
                </Accordion>
            </div>
        </div>
    );
}
