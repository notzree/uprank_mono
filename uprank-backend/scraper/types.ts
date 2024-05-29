interface ScrapeFreelancerData {
    id: string 
    url: string
}

export interface QueueScrapeFreelancersReqest {
    job_id: string
    freelancers: ScrapeFreelancerData[]

}

export interface SQSEvent {
    Records: SQSRecord[];
}

export interface SQSRecord {
    messageId: string;
    receiptHandle: string;
    body: string;
    attributes: SQSAttributes;
    messageAttributes: { [key: string]: any };
    md5OfBody: string;
    eventSource: string;
    eventSourceARN: string;
    awsRegion: string;
}

interface SQSAttributes {
    ApproximateReceiveCount: string;
    SentTimestamp: string;
    SequenceNumber: string;
    MessageGroupId: string;
    SenderId: string;
    MessageDeduplicationId: string;
    ApproximateFirstReceiveTimestamp: string;
}
