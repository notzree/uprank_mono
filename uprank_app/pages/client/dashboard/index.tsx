import * as React from "react";
import { InferGetServerSidePropsType } from "next";
import prisma from "@/prisma/client";
import {
    ResizableHandle,
    ResizablePanel,
    ResizablePanelGroup,
} from "@/components/ui/resizable";
import { FreelancerList } from "./components/freelancer-list";
import Navbar from "./components/navbar";
import { serializeDates } from "@/utils/api_utils/serialize";

export default function Dashboard({
    freelancer_props,
}: any) {
    return (
        <div className="grid min-h-screen w-full">
            <div className="flex flex-col">
                <Navbar />
                <main className="flex flex-1 flex-col items-center justify-center gap-4 p-4 md:gap-8 md:p-6">
                    <div className="flex flex-row items-center justify-center w-full text-center">
                        <div className=" flex flex-col flex-1">
                        <h2>
                            Your current job postings
                        </h2>
                        <div className="flex-1">
                        Render the jobs here (Fetched from upwork api)
                        </div>
                        </div>
                        <div className= "flex-1">
                            <h2>
                                Your former Upranks
                            </h2>
                            <div>
                                Render the former upranks here (Fetched from our database, but restricted to only storing certain data as upwork does not allow caching beyond 24 hours.)
                            </div>
                        </div>
                    </div>
                </main>
            </div>
        </div>
    );
}


