import type { Client } from "./client";
import type { UpdateUpworkFreelancerClientResponse, CreateUpworkFreelancerClientRequest, CreateUpworkFreelancerClientResponse } from "~types/freelancer";
import type { CreateJobClientResponse, CreateJobClientRequest, GetUpworkJobClientResponse } from "~types/job";

export class V1Client implements Client {
    constructor() {
        // Add your constructor logic here
    }
    async getUpworkJob(upwork_job_id: string, auth_token: string): Promise<GetUpworkJobClientResponse>{
        const response = await fetch(
            `${process.env.PLASMO_PUBLIC_BACKEND_URL}/v1/private/jobs/upwork/${upwork_job_id}`,
            {
                method: "GET",
                headers: {
                    "Content-Type": "application/json",
                    "Authorization": `Bearer ${auth_token}`
                },
            }
        );
        const response_json = await response.json();
        if (!response.ok) {
            return {
                upwork_job: null,
                error_msg: response_json.msg
            }
        }
        return {
            upwork_job: response_json,
            error_msg: null
        }
    }

    async createJob(job_data: CreateJobClientRequest, auth_token: string): Promise<CreateJobClientResponse> {
        const response = await fetch(
            `${process.env.PLASMO_PUBLIC_BACKEND_URL}/v1/private/jobs`,
            {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    "Authorization": `Bearer ${auth_token}`
                },
                body: JSON.stringify(job_data),
            }
        );
        const response_json = await response.json();
        if (!response.ok){
            return {
                job: null,
                error_msg: response_json.msg
            }
        }

        return {
            job: response_json,
            error_msg: null
        }
    }

    async createUpworkFreelancers(upwork_job_id: string, freelancer_data: CreateUpworkFreelancerClientRequest,  auth_token:string): Promise<CreateUpworkFreelancerClientResponse>{
        const response = await fetch(
            `${process.env.PLASMO_PUBLIC_BACKEND_URL}/v1/private/jobs/upwork/${upwork_job_id}/freelancers`,
            {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                     "Authorization": `Bearer ${auth_token}`
                },
                body: JSON.stringify(freelancer_data.freelancers),
            }
        );
        const response_json = await response.json();
        if (!response.ok){
            return {
                result: null,
                error_msg: response_json.msg
            }
        }
        return {
            result: response_json,
            error_msg: null
        }
    }

    async updateUpworkFreelancers(upwork_job_id: string, freelancer_data: CreateUpworkFreelancerClientRequest,  auth_token:string): Promise<UpdateUpworkFreelancerClientResponse>{
        const response = await fetch(
            `${process.env.PLASMO_PUBLIC_BACKEND_URL}/v1/private/jobs/upwork/${upwork_job_id}/freelancers/update`,
            {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                     "Authorization": `Bearer ${auth_token}`
                },
                body: JSON.stringify(freelancer_data.freelancers),
            }
        );
        const response_json = await response.json();
        if (!response.ok){
            return {
                result: null,
                error_msg: response_json.msg
            }
        }
        return {
            result: response_json,
            error_msg: null
        }
    }
}