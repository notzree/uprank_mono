import type { UUID } from "crypto";
import type { CreateUpworkFreelancerResponse } from "~types/freelancer";
export type Job = {
    id: string;
    title: string;
    userId: string;
    createdAt: Date | null;
    location: string;
    description: string;
    skills: string[];
    experienceLevel: string;
    hourly: boolean;
    fixed: boolean;
    hourlyRate: number[];
    fixedRate: number;
    averageUprankScore: number | null;
    maxUprankScore: number | null;
    minUprankScore: number | null;

  };

  //Type that gets sent from content -> popup
  export type ScrapedJobData = {
    id: string;
    title: string;
    location: string;
    description: string;
    skills: string[];
    experience_level: string;
    hourly: boolean;
    fixed: boolean;
    hourly_rate: number[];
    fixed_rate: number;

  }



  export type UnstableScrapedJobData = {
    job: ScrapedJobData;
    missingFields: boolean;
  }

  export type GetUpworkJobClientResponse = {
    upwork_job: GetUpworkJobResponse | null
    error_msg: string | null
  };

  type GetUpworkJobResponse = {
    id: string;
    title: string;
    created_at: string;
    location: string;
    description: string;
    skills: string[];
    experience_level: string;
    hourly: boolean;
    hourly_rate: number[];
    edges: EdgeFreelancers | null;
  }

  type EdgeFreelancers = {
    upworkfreelancer: CreateUpworkFreelancerResponse[];
  };

  //Type that gets sent to the service worker
  export type CreateJobProxyRequest = {
    authentication_token: string;
    job: ScrapedJobData;
  }
  //Type that the service worker sends to the backend
  export type CreateJobClientRequest = {
    upwork_job_request: ScrapedJobData | null;
  }

  export type CreateJobClientResponse = {
    job: CreateJobResponse | null;
    error_msg: string | null;
  }

  type CreateJobResponse = {
    id?: UUID;
    originPlatform?: string;
  }


  