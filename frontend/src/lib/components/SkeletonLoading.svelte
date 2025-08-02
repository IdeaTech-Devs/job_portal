<script lang="ts">
	export let count: number = 6;
	export let type: 'job' | 'application' = 'job';
</script>

<div class="space-y-4 sm:space-y-6">
	{#each Array(count) as _, i}
		<div class="bg-white/80 backdrop-blur-md rounded-xl shadow-lg border border-gray-200/50 overflow-hidden">
			<!-- Header -->
			<div class="bg-gradient-to-r from-gray-50 to-gray-100 p-4 sm:p-6 border-b border-gray-200/50">
				<div class="flex flex-col sm:flex-row sm:items-start sm:justify-between">
					<div class="flex-1 mb-3 sm:mb-0">
						<!-- Title Skeleton -->
						<div class="flex flex-col sm:flex-row sm:items-center space-y-2 sm:space-y-0 sm:space-x-3 mb-2">
							<div class="h-6 bg-gray-200 rounded animate-pulse w-48"></div>
							<div class="h-5 bg-gray-200 rounded-full animate-pulse w-32"></div>
						</div>
						<!-- Subtitle Skeleton -->
						<div class="flex flex-col sm:flex-row sm:items-center space-y-1 sm:space-y-0 sm:space-x-4">
							<div class="flex items-center space-x-1">
								<div class="h-4 w-4 bg-gray-200 rounded animate-pulse"></div>
								<div class="h-4 bg-gray-200 rounded animate-pulse w-24"></div>
							</div>
							<div class="flex items-center space-x-1">
								<div class="h-4 w-4 bg-gray-200 rounded animate-pulse"></div>
								<div class="h-4 bg-gray-200 rounded animate-pulse w-32"></div>
							</div>
						</div>
					</div>
					<!-- Status Skeleton -->
					<div class="h-6 bg-gray-200 rounded-full animate-pulse w-24"></div>
				</div>
			</div>

			<!-- Content -->
			<div class="p-4 sm:p-6">
				<div class="grid grid-cols-1 lg:grid-cols-2 gap-4 sm:gap-6">
					<!-- Left Column -->
					<div class="space-y-3 sm:space-y-4">
						<div class="h-5 bg-gray-200 rounded animate-pulse w-32"></div>
						<div class="space-y-2 sm:space-y-3">
							{#each Array(3) as _, j}
								<div class="flex items-center space-x-2 sm:space-x-3">
									<div class="h-4 w-4 bg-gray-200 rounded animate-pulse"></div>
									<div>
										<div class="h-3 bg-gray-200 rounded animate-pulse w-16 mb-1"></div>
										<div class="h-4 bg-gray-200 rounded animate-pulse w-32"></div>
									</div>
								</div>
							{/each}
						</div>
					</div>

					<!-- Right Column -->
					<div class="space-y-3 sm:space-y-4">
						<div class="h-5 bg-gray-200 rounded animate-pulse w-32"></div>
						<div class="space-y-2 sm:space-y-3">
							{#each Array(3) as _, j}
								<div>
									<div class="h-3 bg-gray-200 rounded animate-pulse w-16 mb-1"></div>
									<div class="h-4 bg-gray-200 rounded animate-pulse w-40"></div>
								</div>
							{/each}
						</div>
					</div>
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