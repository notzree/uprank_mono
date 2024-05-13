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

  export type sendJobBody = {
    authentication_token: string;
    job: ScrapedJobData;
  }

  export type UnstableScrapedJobData = {
    job: ScrapedJobData;
    missingFields: boolean;
  }
  