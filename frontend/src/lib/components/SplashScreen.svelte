<script lang="ts">
	import { onMount } from 'svelte';
	import { fade, fly } from 'svelte/transition';
	import { Briefcase, Search, Users, TrendingUp } from 'lucide-svelte';

	let showSplash = true;
	let currentStep = 0;
	const steps = [
		{ icon: Briefcase, text: 'Temukan Pekerjaan Impian', color: 'text-blue-600' },
		{ icon: Search, text: 'Cari dengan Mudah', color: 'text-green-600' },
		{ icon: Users, text: 'Terhubung dengan Perusahaan', color: 'text-purple-600' },
		{ icon: TrendingUp, text: 'Karir yang Menjanjikan', color: 'text-orange-600' }
	];

	onMount(() => {
		// Animate steps
		const stepInterval = setInterval(() => {
			currentStep = (currentStep + 1) % steps.length;
		}, 500);

		// Hide splash after 2 seconds
		setTimeout(() => {
			showSplash = false;
			clearInterval(stepInterval);
		}, 2000);
	});
</script>

{#if showSplash}
	<div 
		class="fixed inset-0 z-50 flex items-center justify-center bg-gradient-to-br from-blue-600 via-purple-600 to-indigo-700"
		transition:fade={{ duration: 300 }}
	>
		<div class="text-center text-white">
			<!-- Logo/Title -->
			<div class="mb-8" transition:fly={{ y: -50, duration: 600 }}>
				<h1 class="text-4xl sm:text-6xl font-bold mb-4 tracking-tight">
					Job Portal
				</h1>
				<p class="text-lg sm:text-xl text-blue-100 opacity-90">
					Platform Lowongan Kerja Terbaik
				</p>
			</div>

			<!-- Animated Steps -->
			<div class="space-y-4" transition:fly={{ y: 50, duration: 600, delay: 300 }}>
				{#each steps as step, index}
					{#if index === currentStep}
						<div 
							class="flex items-center justify-center space-x-3"
							transition:fly={{ y: 20, duration: 400 }}
						>
							<svelte:component this={step.icon} class="h-6 w-6 {step.color}" />
							<span class="text-lg font-medium">{step.text}</span>
						</div>
					{/if}
				{/each}
			</div>

			<!-- Loading Dots -->
			<div class="mt-8 flex justify-center space-x-2">
				{#each Array(3) as _, i}
					<div 
						class="w-2 h-2 bg-white rounded-full animate-pulse"
						style="animation-delay: {i * 0.2}s;"
					></div>
				{/each}
			</div>
		</div>

		<!-- Background Animation -->
		<div class="absolute inset-0 overflow-hidden pointer-events-none">
			<div class="absolute -top-40 -right-40 w-80 h-80 bg-white opacity-10 rounded-full animate-pulse"></div>
			<div class="absolute -bottom-40 -left-40 w-80 h-80 bg-white opacity-10 rounded-full animate-pulse" style="animation-delay: 1s;"></div>
		</div>
	</div>
{/if} 