import { Freelancer } from "@/types/user-types"
export interface Subscription {
    id: number;
    client_id: string;
    freelancer_id: string;
    start_date: string;
    end_date: string | null;
    freelancer_confirmation: boolean;
    freelancer: Freelancer;
}