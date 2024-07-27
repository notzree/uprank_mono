import { Inter } from "next/font/google";
const inter = Inter({ subsets: ["latin"] });
import Link from "next/link";
import { Button } from "@/components/ui/button";
import { CardHeader, CardContent, Card } from "@/components/ui/card";
import { buttonVariants } from "@/components/ui/button";
import {
    sample_average_budget_overrun_percentage,
    sample_average_specialization_score,
    sample_averagebudget_adherence_percentage,
    sample_upwork_job,
    sample_job,
    sample_original_data,
} from "@/public/sample-table-data";
import type { UpworkFreelancer } from "@/types/freelancer";
import * as React from "react";
import { useUser } from "@clerk/nextjs";
import FreelancerTable from "./client/jobs/components/freelancer-table";
import { Job, UpworkJob } from "@/types/job";
import DetailedJobCard from "./client/jobs/components/detailed-job-card";
import FreelancerSearchFilter from "@/pages/client/jobs/components/freelancer-filter";

export default function Component() {
    const [freelancers, setFreelancers] = React.useState<UpworkFreelancer[]>(
        sample_original_data as unknown as UpworkFreelancer[]
    );
    const [showFilters, setShowFilters] = React.useState<boolean>(true);
    const { user } = useUser();
    return (
        <div key="1" className="flex flex-col min-h-screen">
            <header className="flex items-center justify-between px-8 py-6 border-b">
                <div className="flex items-center">
                    <Link className="text-lg font-bold" href="/">
                        Uprank
                    </Link>
                </div>
                <nav className="flex gap-4">
                    {user?.id ? (
                        <Link
                            className="text-sm font-medium hover:underline"
                            href="/client/dashboard"
                        >
                            Dashboard
                        </Link>
                    ) : (
                        <>
                            {/* <Link
                                className="text-sm font-medium hover:underline"
                                href="/sign-up"
                            >
                                Sign up
                            </Link>
                            <p>or</p> */}
                            <Link
                                className="text-sm font-medium hover:underline"
                                href="/sign-in"
                            >
                                Sign in
                            </Link>
                        </>
                    )}
                </nav>
            </header>
            <main className="flex flex-1 flex-col overflow-hidden">

            <div className="flex flex-row w-screen px-6">
                  {showFilters && (
                                <FreelancerSearchFilter
                                    job={
                                        sample_upwork_job as unknown as UpworkJob
                                    }
                                    original_freelancers={
                                        sample_original_data as unknown as UpworkFreelancer[]
                                    }
                                    visible_freelancers={freelancers}
                                    setFreelancers={setFreelancers}
                                />
                            )}
                            <div className="flex flex-col overflow-auto">
                            <DetailedJobCard
                                job={sample_job as unknown as Job}
                                average_budget_adherence_percentage={
                                    sample_averagebudget_adherence_percentage
                                }
                                average_specialization_score={
                                    sample_average_specialization_score
                                }
                                average_budget_overrun_percentage={
                                    sample_average_budget_overrun_percentage
                                }
                                defaultOpen={false}
                            /> 
                            <FreelancerTable
                            original_freelancers={
                                sample_original_data as unknown as UpworkFreelancer[]
                            }
                            average_budget_adherence_percentage={
                                sample_averagebudget_adherence_percentage
                            }
                            average_specialization_score={
                                sample_average_specialization_score
                            }
                            average_budget_overrun_percentage={
                                sample_average_budget_overrun_percentage
                            }
                            visible_freelancers={freelancers}
                            setFreelancers={setFreelancers}
                            showFilters={showFilters}
                            setShowFilters={setShowFilters}
                            job={sample_upwork_job as unknown as UpworkJob}
                        />
                            </div>
            </div>

                <section className="py-24 px-10 text-center">
                    <Link
                        className={buttonVariants({ variant: "outline" })}
                        href="https://forms.gle/cJTuPrewaEYWa61s9"
                    >
                        Register for early access
                    </Link>
                </section>
            </main>
            <footer className="py-6 px-10 text-center text-sm text-gray-600">
                © Uprank Design. All Rights Reserved.
            </footer>
        </div>
    );
}

function computeDemoUpworkStats(freelancers: UpworkFreelancer[]) {

}