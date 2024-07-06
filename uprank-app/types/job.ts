// UUID type definition for simplicity
type UUID = string;

// Schema platform type definition
type Platform = string;


// JobEdges type definition
interface JobEdges {
    upworkjob?: UpworkJob;
    loadedTypes: boolean[];
}

// UpworkJob type definition
export interface UpworkJob {
    id?: string;
    title?: string;
    created_at?: Date;
    updated_at?: Date;
    embedded_at?: Date;
    ranked_at?: Date;
    location?: string;
    description?: string;
    skills?: string[];
    experience_level?: string;
    hourly?: boolean;
    fixed?: boolean;
    hourly_rate?: number[];
    fixed_rate?: number;
    average_uprank_score?: number;
    max_uprank_score?: number;
    min_uprank_score?: number;
    edges: UpworkJobEdges;
    job_upworkjob?: UUID;
}

// UpworkJobEdges type definition
interface UpworkJobEdges {
    // Define any necessary fields here
}

// Job type definition
export interface Job {
    id?: UUID;
    origin_platform?: Platform;
    edges: JobEdges;
    user_job?: string;
}
