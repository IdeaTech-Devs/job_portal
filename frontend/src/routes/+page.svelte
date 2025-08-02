<script lang="ts">
	import { onMount } from 'svelte';
	import type { Job, JobFilter, PaginatedResponse } from '$lib/api';
	import { api } from '$lib/api';
	import JobCard from '$lib/components/JobCard.svelte';
	import FilterPanel from '$lib/components/FilterPanel.svelte';
	import ApplyModal from '$lib/components/ApplyModal.svelte';
	import Pagination from '$lib/components/Pagination.svelte';
	import { Search, Briefcase, Users, TrendingUp, Filter } from 'lucide-svelte';

	let jobs: Job[] = [];
	let allJobs: Job[] = []; // Store all jobs for search
	let locations: string[] = [];
	let filters: JobFilter = {};
	let searchQuery = '';
	let isLoading = true;
	let error = '';
	let selectedJob: Job | null = null;
	let isModalOpen = false;
	let showFilters = false; // For mobile filter toggle
	
	// Pagination state
	let currentPage = 1;
	let pagination = {
		page: 1,
		limit: 12,
		total: 0,
		total_pages: 0,
		has_next: false,
		has_prev: false
	};

	onMount(async () => {
		await loadData();
	});

	async function loadData() {
		try {
			isLoading = true;
			error = '';
			
			const [jobsData, locationsData] = await Promise.all([
				api.getJobs(filters, currentPage, 12),
				api.getLocations()
			]);
			
			jobs = jobsData.jobs;
			allJobs = jobsData.jobs; // For search functionality
			pagination = jobsData.pagination;
			locations = locationsData;
		} catch (err) {
			error = err instanceof Error ? err.message : 'Gagal memuat data';
		} finally {
			isLoading = false;
		}
	}

	async function handleFiltersChange(newFilters: JobFilter) {
		filters = newFilters;
		currentPage = 1; // Reset to first page when filters change
		await loadData();
		// Re-apply search after filter change
		applySearch();
	}

	function handleSearch() {
		applySearch();
	}

	function applySearch() {
		if (!searchQuery.trim()) {
			jobs = allJobs;
			return;
		}

		const query = searchQuery.toLowerCase().trim();
		jobs = allJobs.filter(job => 
			job.position.toLowerCase().includes(query) ||
			job.company.toLowerCase().includes(query) ||
			job.location.toLowerCase().includes(query)
		);
	}

	function handleApply(job: Job) {
		selectedJob = job;
		isModalOpen = true;
	}

	function closeModal() {
		isModalOpen = false;
		selectedJob = null;
	}

	function toggleFilters() {
		showFilters = !showFilters;
	}

	async function handlePageChange(page: number) {
		currentPage = page;
		await loadData();
		// Scroll to top when page changes
		window.scrollTo({ top: 0, behavior: 'smooth' });
	}
</script>

<svelte:head>
	<title>Job Portal - Lowongan Kerja Terbaru</title>
</svelte:head>

<!-- Hero Section -->
<div class="bg-white border-b border-gray-100 py-16">
	<div class="max-w-6xl mx-auto px-4 sm:px-6 lg:px-8">
		<div class="text-center">
			<h1 class="text-4xl md:text-5xl font-bold text-gray-900 mb-6 leading-tight">
				Temukan Pekerjaan
				<span class="text-blue-600">Impian</span> Anda
			</h1>
			<p class="text-lg md:text-xl text-gray-600 mb-12 max-w-3xl mx-auto leading-relaxed">
				Jelajahi ribuan lowongan kerja dari perusahaan terkemuka di Indonesia. 
				Mulai karir baru Anda hari ini!
			</p>
			
			<!-- Search Bar -->
			<div class="max-w-2xl mx-auto mb-12">
				<div class="relative">
					<div class="absolute inset-y-0 left-0 pl-4 flex items-center pointer-events-none">
						<Search class="h-5 w-5 text-gray-400" />
					</div>
					<input 
						type="text"
						bind:value={searchQuery}
						on:input={handleSearch}
						placeholder="Cari lowongan, perusahaan, atau lokasi..."
						class="w-full pl-12 pr-12 py-4 text-lg border border-gray-200 rounded-2xl focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent transition-all duration-200 bg-gray-50 focus:bg-white"
					/>
					{#if searchQuery}
						<button 
							on:click={() => { searchQuery = ''; applySearch(); }}
							class="absolute inset-y-0 right-0 pr-4 flex items-center text-gray-400 hover:text-gray-600 transition-colors"
						>
							×
						</button>
					{/if}
				</div>
			</div>

			<!-- Stats -->
			<div class="flex flex-col sm:flex-row items-center justify-center space-y-6 sm:space-y-0 sm:space-x-12 text-center">
				<div class="flex flex-col items-center">
					<div class="text-3xl font-bold text-gray-900 mb-2">{pagination.total}</div>
					<div class="text-sm text-gray-600">Lowongan Aktif</div>
				</div>
				<div class="flex flex-col items-center">
					<div class="text-3xl font-bold text-gray-900 mb-2">{locations.length}</div>
					<div class="text-sm text-gray-600">Kota Tersedia</div>
				</div>
				<div class="flex flex-col items-center">
					<div class="text-3xl font-bold text-gray-900 mb-2">100%</div>
					<div class="text-sm text-gray-600">Gratis</div>
				</div>
			</div>
		</div>
	</div>
</div>

<!-- Main Content -->
<div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
	<!-- Mobile Filter Toggle -->
	<div class="lg:hidden mb-6">
		<button 
			on:click={toggleFilters}
			class="w-full flex items-center justify-center space-x-2 px-6 py-4 bg-white border border-gray-200 rounded-xl shadow-sm hover:shadow-md transition-all duration-200"
		>
			<Filter class="h-5 w-5 text-gray-600" />
			<span class="font-medium text-gray-700">Filter Lowongan</span>
		</button>
	</div>

	<!-- Mobile Filter Panel -->
	{#if showFilters}
		<div class="lg:hidden mb-8">
			<FilterPanel 
				{filters}
				{locations}
				onFiltersChange={handleFiltersChange}
			/>
		</div>
	{/if}

	<div class="grid grid-cols-1 lg:grid-cols-4 gap-8">
		<!-- Filter Panel (Desktop) -->
		<div class="hidden lg:block">
			<FilterPanel 
				{filters}
				{locations}
				onFiltersChange={handleFiltersChange}
			/>
		</div>

		<!-- Job Listings -->
		<div class="lg:col-span-3">
			<div class="mb-8">
				<div class="flex flex-col sm:flex-row sm:items-center sm:justify-between mb-6">
					<div class="mb-4 sm:mb-0">
						<h2 class="text-2xl font-bold text-gray-900 mb-2">
							Lowongan Kerja Terbaru
						</h2>
						<p class="text-gray-600">
							Menampilkan {jobs.length} lowongan kerja yang sesuai dengan kriteria Anda
							{#if searchQuery}
								untuk "<span class="font-medium text-blue-600">{searchQuery}</span>"
							{/if}
						</p>
					</div>
				</div>
			</div>

			{#if isLoading}
				<div class="flex justify-center items-center py-16">
					<div class="flex items-center space-x-3">
						<div class="animate-spin rounded-full h-8 w-8 border-b-2 border-blue-600"></div>
						<span class="text-gray-600">Memuat lowongan...</span>
					</div>
				</div>
			{:else if error}
				<div class="bg-red-50 border border-red-200 rounded-xl p-8">
					<div class="text-center">
						<p class="text-red-600 mb-4">{error}</p>
						<button 
							on:click={loadData}
							class="bg-red-600 hover:bg-red-700 text-white font-medium py-3 px-6 rounded-lg transition-colors"
						>
							Coba Lagi
						</button>
					</div>
				</div>
			{:else if jobs.length === 0}
				<div class="bg-gray-50 border border-gray-200 rounded-xl p-12">
					<div class="text-center">
						<Search class="mx-auto h-16 w-16 text-gray-400 mb-6" />
						<h3 class="text-xl font-semibold text-gray-900 mb-3">
							{#if searchQuery}
								Tidak ada lowongan yang ditemukan untuk "{searchQuery}"
							{:else}
								Tidak ada lowongan yang ditemukan
							{/if}
						</h3>
						<p class="text-gray-600 mb-6 max-w-md mx-auto">
							{#if searchQuery}
								Coba ubah kata kunci pencarian atau filter yang berbeda
							{:else}
								Coba ubah filter atau cari dengan kriteria yang berbeda
							{/if}
						</p>
						<div class="flex flex-col sm:flex-row gap-3 justify-center">
							{#if searchQuery}
								<button 
									on:click={() => { searchQuery = ''; applySearch(); }}
									class="bg-blue-600 hover:bg-blue-700 text-white font-medium py-3 px-6 rounded-lg transition-colors"
								>
									Hapus Pencarian
								</button>
							{/if}
							<button 
								on:click={() => handleFiltersChange({})}
								class="bg-gray-600 hover:bg-gray-700 text-white font-medium py-3 px-6 rounded-lg transition-colors"
							>
								Hapus Filter
							</button>
						</div>
					</div>
				</div>
			{:else}
				<div class="grid grid-cols-1 md:grid-cols-2 xl:grid-cols-3 gap-6">
					{#each jobs as job (job.id)}
						<JobCard {job} onApply={handleApply} />
					{/each}
				</div>

				<!-- Pagination -->
				<Pagination {pagination} onPageChange={handlePageChange} />
			{/if}
		</div>
	</div>
</div>

<ApplyModal 
	job={selectedJob}
	isOpen={isModalOpen}
	on:close={closeModal}
/> 