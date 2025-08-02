<script lang="ts">
	import { onMount } from 'svelte';
	import { api } from '$lib/api';
	import { applications } from '$lib/stores/applications';
	import JobCard from '$lib/components/JobCard.svelte';
	import FilterPanel from '$lib/components/FilterPanel.svelte';
	import Pagination from '$lib/components/Pagination.svelte';
	import JobSkeleton from '$lib/components/JobSkeleton.svelte';
	import ApplyModal from '$lib/components/ApplyModal.svelte';
	import { Search, Briefcase, MapPin, TrendingUp, Filter } from 'lucide-svelte';
	import type { JobFilter } from '$lib/api';

	let jobs: any[] = [];
	let locations: string[] = [];
	let currentPage = 1;
	let totalPages = 1;
	let totalJobs = 0;
	let isLoading = true;
	let error = '';
	let showFilters = false;

	// Modal state
	let selectedJob: any = null;
	let isModalOpen = false;

	// Filter states
	let searchQuery = '';
	let filters: JobFilter = {};

	async function loadData() {
		try {
			isLoading = true;
			error = '';

			const response = await api.getJobs(filters, currentPage, 12);
			jobs = response.jobs;
			totalPages = response.pagination.total_pages;
			totalJobs = response.pagination.total;

			// Load locations if not already loaded
			if (locations?.length === 0) {
				locations = await api.getLocations();
			}
		} catch (err) {
			error = err instanceof Error ? err.message : 'Gagal memuat data lowongan';
		} finally {
			isLoading = false;
		}
	}

	function handleSearch() {
		currentPage = 1;
		loadData();
	}

	function handleFilter(newFilters: JobFilter) {
		filters = newFilters;
		currentPage = 1;
		loadData();
	}

	function handlePageChange(page: number) {
		currentPage = page;
		loadData();
	}

	function toggleFilters() {
		showFilters = !showFilters;
	}

	function handleApply(job: any) {
		selectedJob = job;
		isModalOpen = true;
	}

	function closeModal() {
		isModalOpen = false;
		selectedJob = null;
	}

	onMount(() => {
		loadData();
	});
</script>

<svelte:head>
	<title>Job Portal - Temukan Pekerjaan Impian Anda</title>
</svelte:head>

<!-- Hero Section -->
<div class="bg-gradient-to-r from-blue-600 to-indigo-700 text-white py-12 sm:py-16">
	<div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
		<div class="text-center">
			<div class="flex items-center justify-center space-x-2 mb-4">
				<Briefcase class="h-8 w-8 sm:h-10 sm:w-10" />
				<h1 class="text-3xl sm:text-4xl lg:text-5xl font-bold">Job Portal</h1>
			</div>
			<p class="text-lg sm:text-xl text-blue-100 mb-8 max-w-2xl mx-auto">
				Temukan pekerjaan impian Anda dengan ribuan lowongan dari perusahaan ternama
			</p>

			<!-- Search Bar -->
			<div class="max-w-2xl mx-auto">
				<div class="relative">
					<Search class="absolute left-4 top-1/2 transform -translate-y-1/2 h-5 w-5 text-gray-400" />
					<input
						type="text"
						bind:value={searchQuery}
						placeholder="Cari posisi, perusahaan, atau lokasi..."
						class="w-full pl-12 pr-4 py-3 sm:py-4 text-gray-900 bg-white rounded-lg shadow-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2"
						on:keydown={(e) => e.key === 'Enter' && handleSearch()}
					/>
					<button
						on:click={handleSearch}
						class="absolute right-2 top-1/2 transform -translate-y-1/2 bg-blue-600 hover:bg-blue-700 text-white px-4 py-2 rounded-md transition-colors"
					>
						Cari
					</button>
				</div>
			</div>

			<!-- Stats -->
			<div class="mt-8 flex flex-col sm:flex-row justify-center items-center space-y-4 sm:space-y-0 sm:space-x-8">
				<div class="text-center">
					<div class="text-2xl sm:text-3xl font-bold">{totalJobs}</div>
					<div class="text-sm text-blue-100">Lowongan Tersedia</div>
				</div>
				<div class="text-center">
					<div class="text-2xl sm:text-3xl font-bold">{locations?.length || 0}</div>
					<div class="text-sm text-blue-100">Lokasi</div>
				</div>
				<div class="text-center">
					<div class="text-2xl sm:text-3xl font-bold">{($applications?.length || 0)}</div>
					<div class="text-sm text-blue-100">Lamaran Anda</div>
				</div>
			</div>
		</div>
	</div>
</div>

<div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
	<!-- Filter Toggle -->
	<div class="flex justify-between items-center mb-6">
		<div class="flex items-center space-x-2">
			<TrendingUp class="h-5 w-5 text-gray-600" />
			<h2 class="text-xl font-semibold text-gray-900">Lowongan Terbaru</h2>
		</div>
		<button
			on:click={toggleFilters}
			class="flex items-center space-x-2 bg-white border border-gray-300 rounded-lg px-4 py-2 text-gray-700 hover:bg-gray-50 transition-colors"
		>
			<Filter class="h-4 w-4" />
			<span class="hidden sm:inline">Filter</span>
		</button>
	</div>

	<!-- Filter Panel -->
	{#if showFilters}
		<FilterPanel
			{filters}
			{locations}
			onFiltersChange={handleFilter}
		/>
	{/if}

	<!-- Loading State -->
	{#if isLoading}
		<JobSkeleton count={6} />
	{:else if error}
		<div class="bg-white/80 backdrop-blur-md rounded-xl shadow-lg border border-gray-200/50 text-center py-8 sm:py-12">
			<div class="text-red-400 mb-4">
				<Briefcase class="mx-auto h-12 w-12 sm:h-16 sm:w-16" />
			</div>
			<h3 class="text-lg sm:text-xl font-semibold text-gray-900 mb-2">Gagal memuat data</h3>
			<p class="text-gray-500 mb-6 max-w-md mx-auto text-sm sm:text-base">
				{error}
			</p>
			<button 
				on:click={loadData}
				class="inline-flex items-center space-x-2 bg-gradient-to-r from-blue-600 to-indigo-600 hover:from-blue-700 hover:to-indigo-700 text-white font-medium py-2 sm:py-3 px-4 sm:px-6 rounded-lg transition-all duration-200"
			>
				<Briefcase class="h-4 w-4" />
				<span>Coba Lagi</span>
			</button>
		</div>
	{:else if jobs?.length === 0}
		<div class="bg-white/80 backdrop-blur-md rounded-xl shadow-lg border border-gray-200/50 text-center py-8 sm:py-12">
			<div class="text-gray-400 mb-4">
				<Briefcase class="mx-auto h-12 w-12 sm:h-16 sm:w-16" />
			</div>
			<h3 class="text-lg sm:text-xl font-semibold text-gray-900 mb-2">Tidak ada lowongan</h3>
			<p class="text-gray-500 mb-6 max-w-md mx-auto text-sm sm:text-base">
				Tidak ada lowongan yang sesuai dengan kriteria pencarian Anda. Coba ubah filter atau kata kunci pencarian.
			</p>
		</div>
	{:else}
		<!-- Job Listings -->
		<div class="grid grid-cols-1 lg:grid-cols-2 gap-4 sm:gap-6">
			{#each jobs as job (job.id)}
				<JobCard {job} onApply={() => handleApply(job)} />
			{/each}
		</div>

		<!-- Pagination -->
		{#if totalPages > 1}
			<div class="mt-8">
				<Pagination
					pagination={{
						page: currentPage,
						limit: 12,
						total: totalJobs,
						total_pages: totalPages,
						has_next: currentPage < totalPages,
						has_prev: currentPage > 1
					}}
					onPageChange={handlePageChange}
				/>
			</div>
		{/if}
	{/if}
</div>

<!-- Apply Modal -->
<ApplyModal 
	job={selectedJob}
	isOpen={isModalOpen}
	on:close={closeModal}
/> 