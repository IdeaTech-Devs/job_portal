<script lang="ts">
	import { ChevronLeft, ChevronRight } from 'lucide-svelte';
	import type { PaginationInfo } from '$lib/api';

	export let pagination: PaginationInfo;
	export let onPageChange: (page: number) => void;

	function getPageNumbers(): number[] {
		const pages: number[] = [];
		const current = pagination.page;
		const total = pagination.total_pages;
		
		// For mobile or small screens, show fewer pages
		const isMobile = window.innerWidth < 768;
		const maxVisiblePages = isMobile ? 3 : 5;
		
		if (total <= maxVisiblePages) {
			// Show all pages if total is small
			for (let i = 1; i <= total; i++) {
				pages.push(i);
			}
		} else {
			// Always show first page
			pages.push(1);
			
			// Calculate range around current page
			let start = Math.max(2, current - Math.floor(maxVisiblePages / 2));
			let end = Math.min(total - 1, start + maxVisiblePages - 3);
			
			// Adjust start if end is too close to total
			if (end === total - 1) {
				start = Math.max(2, total - maxVisiblePages + 2);
			}
			
			// Add ellipsis if needed
			if (start > 2) {
				pages.push(-1); // Ellipsis
			}
			
			// Add middle pages
			for (let i = start; i <= end; i++) {
				pages.push(i);
			}
			
			// Add ellipsis if needed
			if (end < total - 1) {
				pages.push(-1); // Ellipsis
			}
			
			// Always show last page
			pages.push(total);
		}
		
		return pages;
	}

	// Mobile-friendly pagination
	function getMobilePageNumbers(): number[] {
		const current = pagination.page;
		const total = pagination.total_pages;
		
		if (total <= 5) {
			// Show all pages if 5 or fewer
			const pages: number[] = [];
			for (let i = 1; i <= total; i++) {
				pages.push(i);
			}
			return pages;
		}
		
		// For mobile, show current page and adjacent pages
		const pages: number[] = [];
		
		// Always show first page
		pages.push(1);
		
		// Add ellipsis if current page is far from start
		if (current > 3) {
			pages.push(-1);
		}
		
		// Show current page and adjacent pages
		for (let i = Math.max(2, current - 1); i <= Math.min(total - 1, current + 1); i++) {
			pages.push(i);
		}
		
		// Add ellipsis if current page is far from end
		if (current < total - 2) {
			pages.push(-1);
		}
		
		// Always show last page
		if (total > 1) {
			pages.push(total);
		}
		
		return pages;
	}
</script>

{#if pagination.total_pages > 1}
	<div class="flex flex-col items-center space-y-4 mt-8">
		<!-- Desktop Pagination -->
		<div class="hidden md:flex items-center justify-center space-x-1">
			<!-- Previous Button -->
			<button
				on:click={() => onPageChange(pagination.page - 1)}
				disabled={!pagination.has_prev}
				class="flex items-center space-x-1 px-3 py-2 text-sm font-medium text-gray-500 bg-white border border-gray-300 rounded-lg hover:bg-gray-50 hover:text-gray-700 disabled:opacity-50 disabled:cursor-not-allowed transition-all duration-200"
			>
				<ChevronLeft class="h-4 w-4" />
				<span>Sebelumnya</span>
			</button>

			<!-- Page Numbers -->
			<div class="flex items-center space-x-1">
				{#each getPageNumbers() as pageNum}
					{#if pageNum === -1}
						<span class="px-2 py-2 text-gray-400">...</span>
					{:else}
						<button
							on:click={() => onPageChange(pageNum)}
							class="px-3 py-2 text-sm font-medium rounded-lg transition-all duration-200 {pageNum === pagination.page 
								? 'bg-blue-600 text-white' 
								: 'text-gray-500 bg-white border border-gray-300 hover:bg-gray-50 hover:text-gray-700'}"
						>
							{pageNum}
						</button>
					{/if}
				{/each}
			</div>

			<!-- Next Button -->
			<button
				on:click={() => onPageChange(pagination.page + 1)}
				disabled={!pagination.has_next}
				class="flex items-center space-x-1 px-3 py-2 text-sm font-medium text-gray-500 bg-white border border-gray-300 rounded-lg hover:bg-gray-50 hover:text-gray-700 disabled:opacity-50 disabled:cursor-not-allowed transition-all duration-200"
			>
				<span>Selanjutnya</span>
				<ChevronRight class="h-4 w-4" />
			</button>
		</div>

		<!-- Mobile Pagination -->
		<div class="md:hidden flex items-center justify-center space-x-1">
			<!-- Previous Button -->
			<button
				on:click={() => onPageChange(pagination.page - 1)}
				disabled={!pagination.has_prev}
				class="flex items-center px-3 py-2 text-sm font-medium text-gray-500 bg-white border border-gray-300 rounded-lg hover:bg-gray-50 hover:text-gray-700 disabled:opacity-50 disabled:cursor-not-allowed transition-all duration-200"
			>
				<ChevronLeft class="h-4 w-4" />
			</button>

			<!-- Page Numbers -->
			<div class="flex items-center space-x-1">
				{#each getMobilePageNumbers() as pageNum}
					{#if pageNum === -1}
						<span class="px-2 py-2 text-gray-400 text-xs">...</span>
					{:else}
						<button
							on:click={() => onPageChange(pageNum)}
							class="px-2 py-2 text-xs font-medium rounded-lg transition-all duration-200 {pageNum === pagination.page 
								? 'bg-blue-600 text-white' 
								: 'text-gray-500 bg-white border border-gray-300 hover:bg-gray-50 hover:text-gray-700'}"
						>
							{pageNum}
						</button>
					{/if}
				{/each}
			</div>

			<!-- Next Button -->
			<button
				on:click={() => onPageChange(pagination.page + 1)}
				disabled={!pagination.has_next}
				class="flex items-center px-3 py-2 text-sm font-medium text-gray-500 bg-white border border-gray-300 rounded-lg hover:bg-gray-50 hover:text-gray-700 disabled:opacity-50 disabled:cursor-not-allowed transition-all duration-200"
			>
				<ChevronRight class="h-4 w-4" />
			</button>
		</div>

		<!-- Page Info -->
		<div class="text-center text-sm text-gray-500">
			Menampilkan {((pagination.page - 1) * pagination.limit) + 1} - {Math.min(pagination.page * pagination.limit, pagination.total)} dari {pagination.total} lowongan
		</div>
	</div>
{/if} 