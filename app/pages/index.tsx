import { Inter } from "next/font/google";
const inter = Inter({ subsets: ["latin"] });
import Image from "next/image";
import Link from "next/link";
import {
    Accordion,
    AccordionContent,
    AccordionItem,
    AccordionTrigger,
} from "@/components/ui/accordion";
import { buttonVariants } from "@/components/ui/button";
import type { UpworkFreelancer } from "@/types/freelancer";
import * as React from "react";
import { useUser } from "@clerk/nextjs";
type CardData = {
    title: string;
    content: string;
};

export default function Component() {
    const { user } = useUser();
    return (
        <div key="1" className="flex flex-col min-h-screen">
            {/* <head>
            <meta name="viewport" content="width=device-width, initial-scale=1.0"></meta>
            </head> */}
            <header className="flex items-center justify-between px-8 py-6 border-b">
                <div className="flex items-center">
                    <Link className="text-lg font-bold" href="#">
                        Uprank
                    </Link>
                </div>
                <nav className="flex gap-4">
                <Link href="/demo"
                className="text-sm font-medium hover:underline"
                >
                    Demo
                    </Link>
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
            <main className="flex-1">
                <section className="py-16 px-10 text-center flex flex-col justify-center items-center m-auto">
                    <h1 className="text-2xl md:text-5xl font-bold text-left ">
                        The Freelancer Analytics Platform.
                    </h1>
                    <div className="flex justify-center items-center ">
                        <p className="mt-4 text-lg text-foreground w-[20rem] md:w-[40rem] text-center">
                            Replace hours of manual research with comprehensive
                            freelancer metrics like specialization scores,
                            budget adherence rates, and more.
                        </p>
                    </div>
                    <div className="flex space-x-4 justify-center items-center pt-2">
                        <Link
                            className={buttonVariants({ variant: "outline" })}
                            href="/demo"
                        >
                            <span
                            className=" text-xl font-semibold decoration-blue-400 underline decoration-2"
                            >
                                Try the demo
                            </span>
                        </Link>
                        <Link
                            className={buttonVariants({ variant: "outline" })}
                            href="https://forms.gle/cJTuPrewaEYWa61s9"
                        >
                            <span
                            className=" text-xl font-semibold decoration-blue-400 underline decoration-2"
                            >
                                Signup for early access
                            </span>
                        </Link>
                    </div>
                </section>
                <section className="py-2 px-10 bg-background " id="features">
                    <div className="flex flex-col md:flex-row justify-center items-center  px-2 md:px-12 md: space-x-12">
                        <div className="flex flex-col justify-center items-center w-[20rem] md:w-[40rem] h-[12rem]">
                            <h3 className="text-2xl text-center font-semibold">
                                Use AI powered metrics
                            </h3>
                            <p>
                                A second opinion never hurts. Uprank leverages
                                AI models to score each freelancer for your
                                specific project needs, helping you make
                                well-informed hiring decisions quickly and
                                confidently.
                            </p>
                        </div>
                        <div className="flex flex-col justify-center items-center w-[20rem] md:w-[40rem] h-[12rem]">
                            <h3 className="text-2xl text-left font-semibold">
                                {/* Automate your hiring process */}
                                Make Data driven decisions
                            </h3>
                            <p>
                                Uprank automatically collects all freelancer
                                data -{" "}
                                <span className="underline underline-offset-2 decoration-blue-400 decoration-2 ">
                                    including their previously completed jobs
                                </span>
                                , so you can focus on making the right hiring
                                decisions. Rather than navigating through a UI
                                built for showcasing freelancers, see all the
                                data you need in one place.
                            </p>
                        </div>
                    </div>
                </section>
                <div className="relative overflow-hidden">
                    <Image
                        src="/product-hero-image.png"

                        alt="Product Image"
                        width={600}
                        height={400}
                        className="w-full h-[350px] object-cover"
                    />
            <div className="absolute inset-0 bg-gradient-to-b from-transparent to-black/70 flex flex-col justify-end p-6">
                <h3 className="text-2xl font-bold text-white mb-2"></h3>
                <p className="text-white/80 text-sm"></p>
            </div>
            </div>
                <section className="py-24 px-10 flex flex-col justify-center items-center">
                    <h2 className="text-3xl font-bold">FAQ</h2>
                    <Accordion type="single" collapsible className="w-[20rem] md:w-[40rem]">
                        <AccordionItem value="item-1">
                        <AccordionTrigger>
                                Why do I need Uprank?
                            </AccordionTrigger>
                            <AccordionContent>
                                Existing freelancer platforms profit whenever you hire a freelancer, so to make all their freelancers hireable, the provided rating metrics are biased.
                                Every freelancer has a 4.5 star rating, a 95+ job success score, and a 100% completion rate, making it hard to differentiate between real talent and scam artists.
                                <br/>
                                <br/>
                                Uprank scans all of the previous job history of each freelancer and provides unbiased metrics to help you make the best hiring decisions. You wouldn&apos;t hire a full-timer 
                                without a background check, so why hire a freelancer without one?
                                <br/>
                                <br/>
                                See an example with real data (names redacted): <Link href="/demo" className="underline decoration-blue-400 decoration-2 font-semibold">Demo</Link>
                            </AccordionContent>
                        </AccordionItem>
                        <AccordionItem value="item-2">
                            <AccordionTrigger>
                                How do I add a job?
                            </AccordionTrigger>
                            <AccordionContent>
                                We have a chrome extension to scrape data from your job postings on Upwork. You are in control of this process, and can request for your data to be deleted at anytime.
                            </AccordionContent>
                        </AccordionItem>
                        <AccordionItem value="item-3">
                            <AccordionTrigger>What platforms do you support?</AccordionTrigger>
                            <AccordionContent>
                                Currently, we only support Upwork. We have plans to expand to other popular platforms including Fiverr and Contra.
                            </AccordionContent>
                        </AccordionItem>
                        <AccordionItem value="item-4">
                            <AccordionTrigger>I have a feature request or suggestion.</AccordionTrigger>
                            <AccordionContent>
                                <Link className={buttonVariants({"variant": "default"})}
                                href="https://insigh.to/b/uprank">Leave feedback</Link>
                            </AccordionContent>
                        </AccordionItem>
                        <AccordionItem value="item-5">
                            <AccordionTrigger>What features are coming next?</AccordionTrigger>
                            <AccordionContent>
                                Uprank is currently being developed by&nbsp;
                                <Link href="https://www.linkedin.com/in/rz2004/" target="_blank" className="underline decoration-blue-400 decoration-2 font-semibold">me!</Link>&nbsp;
                                Leaving feature requests are the best way to get the features you want.
                                Currently, these are the items on the roadmap:
                                <br/>
                                <br/>
                                <ul>
                                    <li>0. UI/UX overhaul</li>
                                    <li>1. Support for more platforms (Fiverr, Contra)</li>
                                    <li>2. More detailed metrics (Time-to-completion estimates using AI)</li>
                                    <li>3. Automatic listing cross-posting (Have upwork job listings be automatically posted on other platforms)</li>
                                </ul>
                            </AccordionContent>
                        </AccordionItem>
                    </Accordion>
                </section>
            </main>
            <footer className="py-6 px-10 text-center text-sm text-gray-600">
                © Uprank Design. All Rights Reserved.
            </footer>
        </div>
    );
}

function ReplyIcon(props: any) {
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
            <polyline points="9 17 4 12 9 7" />
            <path d="M20 18v-2a4 4 0 0 0-4-4H4" />
        </svg>
    );
}
