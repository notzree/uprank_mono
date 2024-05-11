export type Job = {
    id: string;
    title: string;
    userId: string;
    createdAt: Date;
    location: string;
    description: string;
    skills: string[];
    experienceLevel: string;
    hourly: boolean;
    fixed: boolean;
    hourlyRate: number[];
    fixedRate: number[];
    numMonths: number[];
    averageUprankScore: number | null;
    maxUprankScore: number | null;
    minUprankScore: number | null;

  };
  