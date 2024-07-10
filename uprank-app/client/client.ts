import {Job} from "@/types/job"
import { CreateUserBody } from "@/types/user";


//The client interacts with the server and handles all the transportation logic of fetching / serializing data
export interface Client {
    GetAllJobs(auth_token: string): Promise<GetJobResponse>;
    GetPlatformJobById(job_id: string, auth_token: string): Promise<Job>;
    Rank(req: RankRequest, auth_token: string): Promise<void>;
    SyncUser(req: CreateUserBody): Promise<Response>;
}

export type GetJobResponse = {
    jobs: Job[];
    
}


export type RankRequest = {
    job_id: string;
    user_id: string;
    platform_id: string;
    platform: string;
}