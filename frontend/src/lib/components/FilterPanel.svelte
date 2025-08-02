<script lang="ts">
	import type { JobFilter } from '$lib/api';
	import { Filter, MapPin, DollarSign, X } from 'lucide-svelte';

	export let filters: JobFilter;
	export let locations: string[] = [];
	export let onFiltersChange: (filters: JobFilter) => void;

	const salaryRanges = [
		{ label: '1-3 Juta', min: 1000000, max: 3000000, color: 'bg-green-100 text-green-700' },
		{ label: '3-5 Juta', min: 3000000, max: 5000000, color: 'bg-blue-100 text-blue-700' },
		{ label: '5-8 Juta', min: 5000000, max: 8000000, color: 'bg-purple-100 text-purple-700' },
		{ label: '8 Juta+', min: 8000000, max: 0, color: 'bg-orange-100 text-orange-700' }
	];

	function updateLocation(event: Event) {
		const target = event.target as HTMLSelectElement;
		onFiltersChange({ ...filters, location: target.value || undefined });
	}

	function updateSalaryRange(min: number, max: number) {
		onFiltersChange({ 
			...filters, 
			salary_min: min,
			salary_max: max === 0 ? undefined : max
		});
	}

	function clearFilters() {
		onFiltersChange({});
	}

	$: activeFiltersCount = (filters.location ? 1 : 0) + (filters.salary_min ? 1 : 0);
</script>

<div class="bg-white rounded-xl border border-gray-200 shadow-sm">
	<!-- Header -->
	<div class="border-b border-gray-100 p-6">
		<div class="flex items-center space-x-3">
			<Filter class="h-5 w-5 text-blue-600" />
			<h3 class="text-lg font-semibold text-gray-900">Filter Lowongan</h3>
			{#if activeFiltersCount > 0}
				<span class="ml-auto bg-blue-100 text-blue-700 px-3 py-1 rounded-full text-sm font-medium">
					{activeFiltersCount} aktif
				</span>
			{/if}
		</div>
	</div>
	
	<div class="p-6 space-y-8">
		<!-- Location Filter -->
		<div class="space-y-3">
			<label for="location" class="flex items-center space-x-2 text-sm font-medium text-gray-700">
				<MapPin class="h-4 w-4 text-blue-500" />
				<span>Lokasi</span>
			</label>
			<select 
				id="location"
				class="w-full px-4 py-3 border border-gray-200 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent transition-all duration-200 bg-gray-50 focus:bg-white"
				value={filters.location || ''}
				on:change={updateLocation}
			>
				<option value="">Semua Lokasi</option>
				{#each locations as location}
					<option value={location}>{location}</option>
				{/each}
			</select>
		</div>

		<!-- Salary Range Filter -->
		<div class="space-y-3">
			<label class="flex items-center space-x-2 text-sm font-medium text-gray-700">
				<DollarSign class="h-4 w-4 text-green-500" />
				<span>Rentang Gaji</span>
			</label>
			<div class="space-y-2">
				{#each salaryRanges as range}
					<label class="flex items-center p-3 rounded-lg border border-gray-200 hover:border-blue-300 hover:bg-blue-50 transition-all duration-200 cursor-pointer group">
						<input 
							type="radio" 
							name="salary-range"
							class="mr-3 h-4 w-4 text-blue-600 focus:ring-blue-500"
							checked={filters.salary_min === range.min && filters.salary_max === range.max}
							on:change={() => updateSalaryRange(range.min, range.max)}
						/>
						<span class="text-sm text-gray-700 group-hover:text-gray-900 flex-1">{range.label}</span>
						<div class="px-3 py-1 rounded-full text-xs font-medium {range.color}">
							{range.label}
						</div>
					</label>
				{/each}
			</div>
		</div>

		<!-- Clear Filters -->
		{#if activeFiltersCount > 0}
			<button 
				on:click={clearFilters}
				class="w-full flex items-center justify-center space-x-2 px-4 py-3 bg-gray-100 hover:bg-gray-200 text-gray-700 font-medium rounded-lg transition-all duration-200 group"
			>
				<X class="h-4 w-4 group-hover:scale-110 transition-transform" />
				<span>Hapus Filter ({activeFiltersCount})</span>
			</button>
		{/if}
	</div>
</div> 