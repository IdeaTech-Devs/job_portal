<script lang="ts">
	export let count: number = 6;
</script>

<div class="grid grid-cols-1 lg:grid-cols-2 gap-4 sm:gap-6">
	{#each Array(count) as _, i}
		<div class="bg-white/80 backdrop-blur-md rounded-xl shadow-lg border border-gray-200/50 overflow-hidden">
			<!-- Header -->
			<div class="bg-gradient-to-r from-blue-50 to-indigo-50 p-4 sm:p-6 border-b border-gray-200/50">
				<div class="flex flex-col sm:flex-row sm:items-start sm:justify-between">
					<div class="flex-1 mb-3 sm:mb-0">
						<!-- Title Skeleton -->
						<div class="h-6 bg-gray-200 rounded animate-pulse w-48 mb-2"></div>
						<div class="h-5 bg-blue-200 rounded-full animate-pulse w-32 mb-2"></div>
						<!-- Subtitle Skeleton -->
						<div class="flex items-center space-x-1 mb-1">
							<div class="h-4 w-4 bg-gray-200 rounded animate-pulse"></div>
							<div class="h-4 bg-gray-200 rounded animate-pulse w-24"></div>
						</div>
						<div class="flex items-center space-x-1">
							<div class="h-4 w-4 bg-gray-200 rounded animate-pulse"></div>
							<div class="h-4 bg-gray-200 rounded animate-pulse w-32"></div>
						</div>
					</div>
					<!-- Salary Skeleton -->
					<div class="h-6 bg-green-200 rounded-full animate-pulse w-24"></div>
				</div>
			</div>

			<!-- Content -->
			<div class="p-4 sm:p-6">
				<div class="space-y-3 sm:space-y-4">
					<!-- Company Info -->
					<div class="flex items-center space-x-2 sm:space-x-3">
						<div class="h-4 w-4 bg-gray-200 rounded animate-pulse"></div>
						<div class="h-4 bg-gray-200 rounded animate-pulse w-32"></div>
					</div>
					
					<!-- Location -->
					<div class="flex items-center space-x-2 sm:space-x-3">
						<div class="h-4 w-4 bg-gray-200 rounded animate-pulse"></div>
						<div class="h-4 bg-gray-200 rounded animate-pulse w-24"></div>
					</div>
					
					<!-- Salary Range -->
					<div class="flex items-center space-x-2 sm:space-x-3">
						<div class="h-4 w-4 bg-gray-200 rounded animate-pulse"></div>
						<div class="h-4 bg-gray-200 rounded animate-pulse w-36"></div>
					</div>
				</div>

				<!-- Action Button Skeleton -->
				<div class="mt-6">
					<div class="h-10 bg-blue-200 rounded-lg animate-pulse"></div>
				</div>
			</div>
		</div>
	{/each}
</div>

<style>
	@keyframes pulse {
		0%, 100% {
			opacity: 1;
		}
		50% {
			opacity: 0.5;
		}
	}

	.animate-pulse {
		animation: pulse 2s cubic-bezier(0.4, 0, 0.6, 1) infinite;
	}
</style> 