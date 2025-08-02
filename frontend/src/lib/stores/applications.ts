import { writable } from 'svelte/store';
import type { Application } from '$lib/api';
import { api } from '$lib/api';

// Load applications from localStorage on initialization
function createApplicationsStore() {
	const stored = typeof window !== 'undefined' ? localStorage.getItem('applications') : null;
	const initial = stored ? JSON.parse(stored) : [];
	
	const { subscribe, set, update } = writable<Application[]>(initial);

	return {
		subscribe,
		add: (application: Application) => {
			update(applications => {
				const newApplications = [application, ...applications];
				if (typeof window !== 'undefined') {
					localStorage.setItem('applications', JSON.stringify(newApplications));
				}
				return newApplications;
			});
		},
		loadFromAPI: async () => {
			try {
				const apiApplications = await api.getApplications();
				set(apiApplications);
				// Save to localStorage as backup
				if (typeof window !== 'undefined') {
					localStorage.setItem('applications', JSON.stringify(apiApplications));
				}
			} catch (error) {
				console.error('Failed to load applications from API:', error);
				// Fallback to localStorage if API fails
				const stored = localStorage.getItem('applications');
				if (stored) {
					set(JSON.parse(stored));
				}
			}
		},
		clear: () => {
			set([]);
			if (typeof window !== 'undefined') {
				localStorage.removeItem('applications');
			}
		}
	};
}

export const applications = createApplicationsStore(); 