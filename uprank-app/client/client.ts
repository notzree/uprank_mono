import {Job} from "@/types/job"


//The client interacts with the server and handles all the transportation logic of fetching / serializing data
export interface Client {
    GetAllJobs(auth_token: string): Promise<GetJobResponse>;
}

export type GetJobResponse = {
    jobs: Job[];
    
}

