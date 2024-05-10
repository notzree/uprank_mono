export interface Freelancer {
    id: string,
    first_name: string,
    last_name: string,
    email: string,
    profile_image_url: string | null,
    createdAt: string,
    updatedAt: string,
    current_monthly_rate: number,
    about_me: string | null,
    portfolio_url: string | null,
    cummulative_number_of_clients: number,
    current_number_of_clients: number,
    lead_time: number,
    languages: string[],
    skills: string[]
}

export interface Client {
    id: string,
    first_name: string | null,
    last_name: string | null
    company: string,
    email: string,
    profile_image_url: string | null,
    createdAt: string,
    updatedAt: string,
    months_subscribed: number,
    just_browsing: boolean

}