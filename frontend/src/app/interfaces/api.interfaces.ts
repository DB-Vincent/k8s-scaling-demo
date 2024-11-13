export interface Replica {
    name: string;
    current: boolean;
    nodeName: string;
    status: string;
    startTime: string;
    timeSince?: string; // Optional because we add this client-side
}

export interface ApiResponse {
    replicas: Replica[];
}

export interface ApiError {
    error: string;
}