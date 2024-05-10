export interface Freelancer {
    name: string;
    location: string;
    price: string;
    totalEarnings: string;
    cv: string;
    jobs: Jobs[];
}

export interface Jobs {
    title: string;
    startDate: string;
    endDate: string;
    earnings: number;
    hours: number;
    description: string;
    proposals: number;
    interviews: number;
    client : Client;
}

export interface Client {
    location: string;
    totalSpend: number;
}