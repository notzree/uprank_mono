import { CreateUserBody } from "@/types/user";
import { Client, GetJobResponse, RankRequest } from "./client";
import { Job } from "@/types/job";
import { stringify } from "querystring";
export class v1Client implements Client {
    private base_url: string;
    
    constructor(base_url: string | undefined) {
        if (base_url === undefined) {
            throw new Error("base_url is undefined");
        }
        this.base_url = base_url;
    }
    async GetAllJobs(auth_token: string): Promise<GetJobResponse> {
        const url = this.base_url + "/v1/private/jobs";
        const response = await fetch(url, this._build_request(auth_token, "GET"));
        const data = await response.json();
        return {
            jobs: data
        }
    }

    private _build_request(auth_token: string, method: string, body?: any): RequestInit{
        return {
            method: method,
            headers: {
                "Content-Type": "application/json",
                "Authorization": `Bearer ${auth_token}`
            },
            body: JSON.stringify(body),
        }
    }

    async GetPlatformJobById(job_id: string,auth_token: string): Promise<Job> {
        const url = this.base_url + "/v1/private/jobs/" + job_id + "/"
        const response = await fetch(url, this._build_request(auth_token, "GET"));
        const data = await response.json();
        return data;
    }

    async Rank(req: RankRequest, auth_token: string): Promise<void> {
        const url = this.base_url + "/v1/test";
        const response = await fetch(url, this._build_request(auth_token, "POST", req));
        if (!response.ok) {
            throw new Error("Failed to rank jobs");
        }
    }

    async SyncUser(req: CreateUserBody): Promise<Response> {
        const url = this.base_url + "/v1/public/users";
        const response = await fetch(url, this._build_request("", "POST", JSON.stringify(req)));
        return response;
    }
}

