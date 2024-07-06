import { AnyARecord } from "dns";
import { Client, GetJobResponse } from "./client";

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
}