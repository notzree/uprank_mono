import { CreateUserBody } from "@/types/user";
import { Job } from "@/types/job";

export class BackendClient {
    private base_url: string;

    constructor(base_url = "") {
        if (base_url != "") {
            console.log("Default url over-ridden, using: " + base_url);
            this.base_url = base_url;
            return
        } 
            if (
                process.env.NODE_ENV=="development" &&
                process.env.NEXT_PUBLIC_BACKEND_DEV_BASE_URL
            ) {
                this.base_url = process.env.NEXT_PUBLIC_BACKEND_DEV_BASE_URL;
            } else if (
                process.env.NODE_ENV === "production" &&
                process.env.NEXT_PUBLIC_BACKEND_PROD_BASE_URL
            ) {
                this.base_url = process.env.NEXT_PUBLIC_BACKEND_PROD_BASE_URL;
            }
            else {
                throw new Error("No backend url found");
            }
    }

    async GetAllJobs(auth_token: string): Promise<Job[]> {
        const url = this.base_url + "/v1/private/jobs";
        const response = await fetch(
            url,
            this._build_request(auth_token, "GET")
        );
        const data = await response.json();
        return data;
    }

    private _build_request(
        auth_token: string,
        method: string,
        body?: any
    ): RequestInit {
        return {
            method: method,
            headers: {
                "Content-Type": "application/json",
                Authorization: `Bearer ${auth_token}`,
            },
            body: JSON.stringify(body),
        };
    }

    async GetPlatformJobById(job_id: string, auth_token: string): Promise<Job> {
        const url = this.base_url + "/v1/private/jobs/" + job_id + "/";
        const response = await fetch(
            url,
            this._build_request(auth_token, "GET")
        );
        const data = await response.json();
        return data;
    }

    async Rank(req: RankRequest, auth_token: string): Promise<void> {
        const url = this.base_url + "/v1/test";
        const response = await fetch(
            url,
            this._build_request(auth_token, "POST", req)
        );
        if (!response.ok) {
            throw new Error("Failed to rank jobs");
        }
    }

    async SyncUser(req: CreateUserBody): Promise<Response> {
        const url = this.base_url + "/v1/public/users";
        const response = await fetch(
            url,
            this._build_request("", "POST", JSON.stringify(req))
        );
        return response;
    }
}

export type GetJobResponse = {
    jobs: Job[];
};

export type RankRequest = {
    job_id: string;
    user_id: string;
    platform_id: string;
    platform: string;
};
