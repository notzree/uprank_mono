


export type CreateUserBody = {
    user: User;
    completed_onboarding: boolean;
}
export type User = {
    id: string;
    first_name: string;
    company_name: string;
    email: string;
    created_at: Date,
    updated_at: Date,
    last_login: Date,
}