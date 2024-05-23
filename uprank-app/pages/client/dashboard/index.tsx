import * as React from "react";
import { InferGetServerSidePropsType } from "next";
import {
    ResizableHandle,
    ResizablePanel,
    ResizablePanelGroup,
} from "@/components/ui/resizable";
import Navbar from "./components/navbar";
import { serializeDates } from "@/utils/api_utils/serialize";
export default function Dashboard({ jobs }: any) {
    
    return (
        <div className="grid min-h-screen w-full">
            <div className="flex flex-col">
                <Navbar />
                <main className="flex flex-1 flex-col items-center justify-center gap-4 p-4 md:gap-8 md:p-6">
                    <div className="flex flex-row items-center justify-center w-full text-center">
                        <div className=" flex flex-col flex-1">
                            <h2>Your current job postings</h2>
                            <div className="flex-1">
                            {/* {
                                jobs.map((job: any) => {
                                    return (
                                        <div key={job.id} className="flex flex-row justify-between">
                                            <div>
                                                <h3>{job.title}</h3>
                                                <p>{job.created_at}</p>
                                            </div>
                                            <div>
                                                <button>Edit</button>
                                                <button>Delete</button>
                                            </div>
                                        </div>
                                    );
                                })
                            } */}
                            </div>
                        </div>
                        <div className="flex-1">
                            <h2>Your former Upranks</h2>
                            <div>
                                
                            </div>
                        </div>
                    </div>
                </main>
            </div>
        </div>
    );
}



// import { getAuth } from "@clerk/nextjs/server";
// import { GetServerSideProps } from "next";
// export const getServerSideProps: GetServerSideProps = async (ctx) => {

//     const { userId } = getAuth(ctx.req);
//     if (!userId) {
//         return {
//             redirect: {
//                 destination: "/",
//                 permanent: false,
//             },
//         };
//     }

//     const jobs = await prisma.job.findMany({
//         where: {
//             user_id: userId,
//         },
//     });
//     jobs.forEach((job) => {
//         job = serializeDates(job);
//     });

//     return {
//         props: {
//             jobs: jobs,
//         },
//     };
// }
