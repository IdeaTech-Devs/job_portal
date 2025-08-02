import { writable } from 'svelte/store';
import type { Application } from '$lib/api';

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
		clear: () => {
			set([]);
			if (typeof window !== 'undefined') {
				localStorage.removeItem('applications');
			}
		}
	};
}

export const applications = createApplicationsStore(); 