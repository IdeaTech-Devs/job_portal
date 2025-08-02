import { w as writable } from "./index.js";
const API_BASE_URL = "https://jobportal-production-fdc2.up.railway.app/api";
class ApiClient {
  async request(endpoint, options = {}) {
    const url = `${API_BASE_URL}${endpoint}`;
    const response = await fetch(url, {
      ...options,
      headers: {
        "Content-Type": "application/json",
        ...options.headers
      }
    });
    if (!response.ok) {
      const error = await response.json().catch(() => ({ error: "Network error" }));
      throw new Error(error.error || `HTTP ${response.status}`);
    }
    return response.json();
  }
  async getJobs(filters = {}, page = 1, limit = 12) {
    const params = new URLSearchParams();
    if (filters.location) params.append("location", filters.location);
    if (filters.salary_min) params.append("salary_min", filters.salary_min.toString());
    if (filters.salary_max) params.append("salary_max", filters.salary_max.toString());
    params.append("page", page.toString());
    params.append("limit", limit.toString());
    return this.request(`/jobs?${params.toString()}`);
  }
  async getJob(id) {
    return this.request(`/jobs/${id}`);
  }
  async getLocations() {
    return this.request("/locations");
  }
  async submitApplication(formData) {
    const url = `${API_BASE_URL}/applications`;
    const response = await fetch(url, {
      method: "POST",
      body: formData
    });
    if (!response.ok) {
      const error = await response.json().catch(() => ({ error: "Network error" }));
      throw new Error(error.error || `HTTP ${response.status}`);
    }
    return response.json();
  }
  async getApplications() {
    return this.request("/applications");
  }
}
const api = new ApiClient();
function createApplicationsStore() {
  const stored = typeof window !== "undefined" ? localStorage.getItem("applications") : null;
  const initial = stored ? JSON.parse(stored) : [];
  const { subscribe, set, update } = writable(initial);
  return {
    subscribe,
    add: (application) => {
      update((applications2) => {
        const newApplications = [application, ...applications2];
        if (typeof window !== "undefined") {
          localStorage.setItem("applications", JSON.stringify(newApplications));
        }
        return newApplications;
      });
    },
    loadFromAPI: async () => {
      try {
        const apiApplications = await api.getApplications();
        set(apiApplications);
        if (typeof window !== "undefined") {
          localStorage.setItem("applications", JSON.stringify(apiApplications));
        }
      } catch (error) {
        console.error("Failed to load applications from API:", error);
        const stored2 = localStorage.getItem("applications");
        if (stored2) {
          set(JSON.parse(stored2));
        }
      }
    },
    clear: () => {
      set([]);
      if (typeof window !== "undefined") {
        localStorage.removeItem("applications");
      }
    }
  };
}
const applications = createApplicationsStore();
export {
  applications as a
};
