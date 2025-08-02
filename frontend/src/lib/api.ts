const API_BASE_URL = import.meta.env.VITE_API_URL || 'https://jobportal-production-fdc2.up.railway.app/api';

export interface Job {
	id: number;
	position: string;
	company: string;
	location: string;
	salary_min: number;
	salary_max: number;
	created_at: string;
}

export interface Application {
	id: number;
	job_id: number;
	name: string;
	email: string;
	cv_filename: string;
	applied_at: string;
	job?: Job;
}

export interface JobFilter {
	location?: string;
	salary_min?: number;
	salary_max?: number;
	search?: string;
}

export interface PaginationInfo {
	page: number;
	limit: number;
	total: number;
	total_pages: number;
	has_next: boolean;
	has_prev: boolean;
}

export interface PaginatedResponse {
	jobs: Job[];
	pagination: PaginationInfo;
}

class ApiClient {
	private async request<T>(endpoint: string, options: RequestInit = {}): Promise<T> {
		const url = `${API_BASE_URL}${endpoint}`;
		const response = await fetch(url, {
			...options,
			headers: {
				'Content-Type': 'application/json',
				...options.headers,
			},
		});

		if (!response.ok) {
			const error = await response.json().catch(() => ({ error: 'Network error' }));
			throw new Error(error.error || `HTTP ${response.status}`);
		}

		return response.json();
	}

	async getJobs(filters: JobFilter = {}, page: number = 1, limit: number = 12): Promise<PaginatedResponse> {
		const params = new URLSearchParams();
		if (filters.location) params.append('location', filters.location);
		if (filters.salary_min) params.append('salary_min', filters.salary_min.toString());
		if (filters.salary_max) params.append('salary_max', filters.salary_max.toString());
		if (filters.search) params.append('search', filters.search);
		params.append('page', page.toString());
		params.append('limit', limit.toString());

		return this.request<PaginatedResponse>(`/jobs?${params.toString()}`);
	}

	async getJob(id: number): Promise<Job> {
		return this.request<Job>(`/jobs/${id}`);
	}

	async getLocations(): Promise<string[]> {
		return this.request<string[]>('/locations');
	}

	async submitApplication(formData: FormData): Promise<{ message: string; application: Application }> {
		const url = `${API_BASE_URL}/applications`;
		const response = await fetch(url, {
			method: 'POST',
			body: formData,
		});

		if (!response.ok) {
			const error = await response.json().catch(() => ({ error: 'Network error' }));
			throw new Error(error.error || `HTTP ${response.status}`);
		}

		return response.json();
	}

	async getApplications(): Promise<Application[]> {
		return this.request<Application[]>('/applications');
	}
}

export const api = new ApiClient(); 
