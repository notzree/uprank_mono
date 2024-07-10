
export type CreateUserBody = {
    user:{id : string;
    email : string;
    first_name: string
    company_name: string;
    created_at: Date;
    updated_at: Date;
    last_login: Date;
    }
    completed_onboarding: boolean;
}

