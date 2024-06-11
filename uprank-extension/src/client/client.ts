import type { CreateUpworkFreelancerClientResponse, CreateUpworkFreelancerClientRequest,UpdateUpworkFreelancerClientResponse  } from "~types/freelancer";
import type { CreateJobClientResponse, CreateJobClientRequest, GetUpworkJobClientResponse } from "~types/job";

//The client interacts with the server and handles all the transportation logic of fetching / serializing data
export interface Client {
    getUpworkJob(upwork_job_id: string, auth_token:string): Promise<GetUpworkJobClientResponse>;
    createJob(job_data: CreateJobClientRequest,auth_token:string): Promise<CreateJobClientResponse>;
    createUpworkFreelancers (upwork_job_id: string, freelancer_data: CreateUpworkFreelancerClientRequest, auth_token:string): Promise<CreateUpworkFreelancerClientResponse>;
    updateUpworkFreelancers(upwork_job_id: string, freelancer_data: CreateUpworkFreelancerClientRequest, auth_token:string): Promise<UpdateUpworkFreelancerClientResponse>;
}
