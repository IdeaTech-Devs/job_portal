<script lang="ts">
	import { applications } from '$lib/stores/applications';
	import { FileText, Calendar, MapPin, Building, DollarSign, CheckCircle, Clock } from 'lucide-svelte';
	import type { PageProps } from './$types';

	export const data: PageProps['data'] = undefined as any;
	export const params: PageProps['params'] = undefined as any;

	function formatDate(dateString: string): string {
		const date = new Date(dateString);
		return date.toLocaleDateString('id-ID', {
			year: 'numeric',
			month: 'long',
			day: 'numeric',
			hour: '2-digit',
			minute: '2-digit'
		});
	}

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

	function getStatusColor(status: string): string {
		switch (status) {
			case 'Menunggu Review':
				return 'bg-yellow-100 text-yellow-800';
			case 'Diterima':
				return 'bg-green-100 text-green-800';
			case 'Ditolak':
				return 'bg-red-100 text-red-800';
			default:
				return 'bg-gray-100 text-gray-800';
		}
	}
</script>

<svelte:head>
	<title>Lamaran Saya - Job Portal</title>
</svelte:head>

<!-- Header Section -->
<div class="bg-gradient-to-r from-green-600 to-emerald-700 text-white py-6 sm:py-8 mb-6 sm:mb-8 rounded-2xl">
	<div class="max-w-4xl mx-auto px-4 sm:px-6 lg:px-8 text-center">
		<div class="flex items-center justify-center space-x-2 sm:space-x-3 mb-3 sm:mb-4">
			<FileText class="h-6 w-6 sm:h-8 sm:w-8" />
			<h1 class="text-2xl sm:text-3xl font-bold">Lamaran Saya</h1>
		</div>
		<p class="text-green-100 text-base sm:text-lg mb-4 sm:mb-6">
			Kelola semua lamaran kerja yang telah Anda ajukan
		</p>
		<div class="flex flex-col sm:flex-row items-center justify-center space-y-3 sm:space-y-0 sm:space-x-8">
			<div class="text-center">
				<div class="text-xl sm:text-2xl font-bold">{$applications.length}</div>
				<div class="text-xs sm:text-sm text-green-100">Total Lamaran</div>
			</div>
			<div class="text-center">
				<div class="text-xl sm:text-2xl font-bold">
					{$applications.filter(app => app.job).length}
				</div>
				<div class="text-xs sm:text-sm text-green-100">Lowongan Aktif</div>
			</div>
		</div>
	</div>
</div>

<div class="max-w-4xl mx-auto px-4 sm:px-6 lg:px-8">
	{#if $applications.length === 0}
		<div class="bg-white/80 backdrop-blur-md rounded-xl shadow-lg border border-gray-200/50 text-center py-8 sm:py-12">
			<div class="text-gray-400 mb-4">
				<FileText class="mx-auto h-12 w-12 sm:h-16 sm:w-16" />
			</div>
			<h3 class="text-lg sm:text-xl font-semibold text-gray-900 mb-2">Belum ada lamaran</h3>
			<p class="text-gray-500 mb-6 max-w-md mx-auto text-sm sm:text-base">
				Anda belum mengajukan lamaran kerja. Mulai cari lowongan yang sesuai dengan keahlian dan minat Anda!
			</p>
			<a href="/" class="inline-flex items-center space-x-2 bg-gradient-to-r from-blue-600 to-indigo-600 hover:from-blue-700 hover:to-indigo-700 text-white font-medium py-2 sm:py-3 px-4 sm:px-6 rounded-lg transition-all duration-200">
				<FileText class="h-4 w-4" />
				<span>Cari Lowongan</span>
			</a>
		</div>
	{:else}
		<div class="space-y-4 sm:space-y-6">
			{#each $applications as application (application.id)}
				<div class="bg-white/80 backdrop-blur-md rounded-xl shadow-lg border border-gray-200/50 overflow-hidden hover:shadow-xl transition-all duration-300">
					<!-- Header -->
					<div class="bg-gradient-to-r from-blue-50 to-indigo-50 p-4 sm:p-6 border-b border-gray-200/50">
						<div class="flex flex-col sm:flex-row sm:items-start sm:justify-between">
							<div class="flex-1 mb-3 sm:mb-0">
								<div class="flex flex-col sm:flex-row sm:items-center space-y-2 sm:space-y-0 sm:space-x-3 mb-2">
									<h3 class="text-lg sm:text-xl font-bold text-gray-900">
										{application.job?.position || 'Posisi tidak tersedia'}
									</h3>
									<span class="px-2 sm:px-3 py-1 bg-blue-100 text-blue-700 text-xs font-medium rounded-full self-start sm:self-auto">
										{application.job?.company || 'Perusahaan tidak tersedia'}
									</span>
								</div>
								<div class="flex flex-col sm:flex-row sm:items-center space-y-1 sm:space-y-0 sm:space-x-4 text-sm text-gray-600">
									<div class="flex items-center space-x-1">
										<MapPin class="h-4 w-4" />
										<span>{application.job?.location || 'Lokasi tidak tersedia'}</span>
									</div>
									<div class="flex items-center space-x-1">
										<Calendar class="h-4 w-4" />
										<span>{formatDate(application.applied_at)}</span>
									</div>
								</div>
							</div>
							<div class="flex items-center space-x-2">
								<span class="px-2 sm:px-3 py-1 rounded-full text-xs font-medium {getStatusColor('Menunggu Review')}">
									<Clock class="h-3 w-3 inline mr-1" />
									Menunggu Review
								</span>
							</div>
						</div>
					</div>

					<!-- Content -->
					<div class="p-4 sm:p-6">
						<div class="grid grid-cols-1 lg:grid-cols-2 gap-4 sm:gap-6">
							<!-- Job Details -->
							<div class="space-y-3 sm:space-y-4">
								<h4 class="font-semibold text-gray-900 mb-2 sm:mb-3">Detail Lowongan</h4>
								<div class="space-y-2 sm:space-y-3">
									<div class="flex items-center space-x-2 sm:space-x-3">
										<Building class="h-4 w-4 text-blue-500" />
										<div>
											<div class="text-xs sm:text-sm text-gray-500">Perusahaan</div>
											<div class="font-medium text-sm sm:text-base">{application.job?.company || 'Tidak tersedia'}</div>
										</div>
									</div>
									<div class="flex items-center space-x-2 sm:space-x-3">
										<MapPin class="h-4 w-4 text-red-500" />
										<div>
											<div class="text-xs sm:text-sm text-gray-500">Lokasi</div>
											<div class="font-medium text-sm sm:text-base">{application.job?.location || 'Tidak tersedia'}</div>
										</div>
									</div>
									{#if application.job}
										<div class="flex items-center space-x-2 sm:space-x-3">
											<DollarSign class="h-4 w-4 text-green-500" />
											<div>
												<div class="text-xs sm:text-sm text-gray-500">Gaji</div>
												<div class="font-medium text-sm sm:text-base">{formatSalary(application.job.salary_min, application.job.salary_max)}</div>
											</div>
										</div>
									{/if}
								</div>
							</div>

							<!-- Application Details -->
							<div class="space-y-3 sm:space-y-4">
								<h4 class="font-semibold text-gray-900 mb-2 sm:mb-3">Detail Lamaran</h4>
								<div class="space-y-2 sm:space-y-3">
									<div>
										<div class="text-xs sm:text-sm text-gray-500">Nama</div>
										<div class="font-medium text-sm sm:text-base">{application.name}</div>
									</div>
									<div>
										<div class="text-xs sm:text-sm text-gray-500">Email</div>
										<div class="font-medium text-sm sm:text-base">{application.email}</div>
									</div>
									<div>
										<div class="text-xs sm:text-sm text-gray-500">CV</div>
										<div class="font-medium text-sm sm:text-base text-blue-600">{application.cv_filename}</div>
									</div>
								</div>
							</div>
						</div>
					</div>
				</div>
			{/each}
		</div>

		<!-- Summary -->
		<div class="mt-6 sm:mt-8 bg-white/80 backdrop-blur-md rounded-xl shadow-lg border border-gray-200/50 p-4 sm:p-6">
			<div class="text-center">
				<div class="flex items-center justify-center space-x-2 mb-2">
					<CheckCircle class="h-4 w-4 sm:h-5 sm:w-5 text-green-600" />
					<span class="text-base sm:text-lg font-semibold text-gray-900">Ringkasan Lamaran</span>
				</div>
				<p class="text-gray-600 mb-3 sm:mb-4 text-sm sm:text-base">
					Total lamaran: <span class="font-semibold text-blue-600">{$applications.length}</span>
				</p>
				<div class="flex justify-center space-x-6 text-xs sm:text-sm text-gray-500">
					<span>Semua lamaran akan ditinjau dalam 1-3 hari kerja</span>
				</div>
			</div>
		</div>
	{/if}
</div> 