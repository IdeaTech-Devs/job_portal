<script lang="ts">
	import type { Job } from '$lib/api';
	import { MapPin, Building, DollarSign, Calendar, ArrowRight } from 'lucide-svelte';

	export let job: Job;
	export let onApply: (job: Job) => void;

	function formatSalary(min: number, max: number): string {
		const formatCurrency = (amount: number) => {
			return new Intl.NumberFormat('id-ID', {
				style: 'currency',
				currency: 'IDR',
				minimumFractionDigits: 0,
				maximumFractionDigits: 0,
			}).format(amount);
		};

		return `${formatCurrency(min)} - ${formatCurrency(max)}`;
	}

	function formatDate(dateString: string): string {
		const date = new Date(dateString);
		return date.toLocaleDateString('id-ID', {
			month: 'short',
			day: 'numeric',
			year: 'numeric'
		});
	}
</script>

<div class="group bg-white rounded-xl border border-gray-200 hover:border-blue-300 hover:shadow-lg transition-all duration-300 overflow-hidden">
	<div class="p-6">
		<!-- Header -->
		<div class="flex justify-between items-start mb-4">
			<div class="flex-1">
				<h3 class="text-lg font-semibold text-gray-900 mb-3 group-hover:text-blue-600 transition-colors duration-200 line-clamp-2">
					{job.position}
				</h3>
				<div class="flex items-center text-gray-600 mb-2">
					<Building class="h-4 w-4 mr-2 text-blue-500 flex-shrink-0" />
					<span class="text-sm font-medium truncate">{job.company}</span>
				</div>
				<div class="flex items-center text-gray-600 mb-4">
					<MapPin class="h-4 w-4 mr-2 text-gray-400 flex-shrink-0" />
					<span class="text-sm truncate">{job.location}</span>
				</div>
			</div>
			<div class="flex items-center space-x-1 text-xs text-gray-400 ml-4">
				<Calendar class="h-3 w-3" />
				<span>{formatDate(job.created_at)}</span>
			</div>
		</div>

		<!-- Salary -->
		<div class="flex items-center justify-between mb-6">
			<div class="flex items-center text-green-600 font-semibold">
				<DollarSign class="h-4 w-4 mr-1" />
				<span class="text-sm">{formatSalary(job.salary_min, job.salary_max)}</span>
			</div>
			<div class="px-3 py-1 bg-green-100 text-green-700 text-xs font-medium rounded-full">
				Aktif
			</div>
		</div>

		<!-- Apply Button -->
		<button 
			on:click={() => onApply(job)}
			class="w-full bg-blue-600 hover:bg-blue-700 text-white font-medium py-3 px-4 rounded-lg transition-all duration-200 flex items-center justify-center space-x-2 group-hover:shadow-md"
		>
			<span>Lamar Sekarang</span>
			<ArrowRight class="h-4 w-4 group-hover:translate-x-1 transition-transform" />
		</button>
	</div>
</div> 