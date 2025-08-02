<script lang="ts">
	import type { Job } from '$lib/api';
	import { api } from '$lib/api';
	import { applications } from '$lib/stores/applications';
	import { X, Upload, CheckCircle, AlertCircle, Building, MapPin, DollarSign } from 'lucide-svelte';
	import { createEventDispatcher } from 'svelte';

	const dispatch = createEventDispatcher();

	export let job: Job | null = null;
	export let isOpen = false;

	let name = '';
	let email = '';
	let cvFile: File | null = null;
	let isLoading = false;
	let error = '';
	let success = false;

	// Generate unique IDs for form fields
	let formId = Math.random().toString(36).substr(2, 9);
	
	// Reactive IDs that update when formId changes
	$: nameId = `name-${formId}`;
	$: emailId = `email-${formId}`;
	$: cvId = `cv-${formId}`;

	function closeModal() {
		isOpen = false;
		resetForm();
		dispatch('close');
	}

	function resetForm() {
		name = '';
		email = '';
		cvFile = null;
		error = '';
		success = false;
		// Generate new form ID when modal is reset
		formId = Math.random().toString(36).substr(2, 9);
	}

	function handleFileChange(event: Event) {
		const target = event.target as HTMLInputElement;
		if (target.files && target.files[0]) {
			cvFile = target.files[0];
		}
	}

	async function handleSubmit() {
		if (!job || !name || !email || !cvFile) {
			error = 'Semua field harus diisi';
			return;
		}

		if (!email.includes('@')) {
			error = 'Format email tidak valid';
			return;
		}

		if (!cvFile.name.toLowerCase().endsWith('.pdf')) {
			error = 'Hanya file PDF yang diperbolehkan';
			return;
		}

		if (cvFile.size > 5 * 1024 * 1024) {
			error = 'Ukuran file maksimal 5MB';
			return;
		}

		isLoading = true;
		error = '';

		try {
			const formData = new FormData();
			formData.append('name', name);
			formData.append('email', email);
			formData.append('job_id', job.id.toString());
			formData.append('cv', cvFile);

			const result = await api.submitApplication(formData);
			
			// Add to local store
			applications.add(result.application);
			
			success = true;
			setTimeout(() => {
				closeModal();
			}, 2000);
		} catch (err) {
			error = err instanceof Error ? err.message : 'Terjadi kesalahan';
		} finally {
			isLoading = false;
		}
	}

	// Reset form when modal opens
	$: if (isOpen && !success) {
		resetForm();
	}
</script>

{#if isOpen}
	<div class="fixed inset-0 bg-black/50 backdrop-blur-sm flex items-center justify-center p-4 z-50">
		<div class="bg-white rounded-xl max-w-md w-full max-h-[90vh] overflow-y-auto shadow-xl">
			<!-- Header -->
			<div class="border-b border-gray-100 p-6">
				<div class="flex justify-between items-center">
					<h3 class="text-xl font-semibold text-gray-900">Lamar Pekerjaan</h3>
					<button 
						on:click={closeModal}
						class="text-gray-400 hover:text-gray-600 transition-colors"
					>
						<X class="h-6 w-6" />
					</button>
				</div>
			</div>

			<div class="p-6">
				{#if job}
					<!-- Job Info -->
					<div class="bg-gray-50 p-4 rounded-lg mb-6">
						<h4 class="font-semibold text-gray-900 mb-3">{job.position}</h4>
						<div class="space-y-2 text-sm text-gray-600">
							<div class="flex items-center space-x-2">
								<Building class="h-4 w-4 text-blue-500" />
								<span>{job.company}</span>
							</div>
							<div class="flex items-center space-x-2">
								<MapPin class="h-4 w-4 text-gray-400" />
								<span>{job.location}</span>
							</div>
							<div class="flex items-center space-x-2">
								<DollarSign class="h-4 w-4 text-green-500" />
								<span>Rp {job.salary_min.toLocaleString()} - Rp {job.salary_max.toLocaleString()}</span>
							</div>
						</div>
					</div>
				{/if}

				{#if success}
					<!-- Success State -->
					<div class="text-center py-8">
						<div class="flex justify-center mb-4">
							<div class="p-3 bg-green-100 rounded-full">
								<CheckCircle class="h-8 w-8 text-green-600" />
							</div>
						</div>
						<h4 class="text-lg font-semibold text-gray-900 mb-2">Lamaran Berhasil Dikirim!</h4>
						<p class="text-gray-600">Tim HR akan menghubungi Anda dalam 1-3 hari kerja.</p>
					</div>
				{:else}
					<!-- Form -->
					<form on:submit|preventDefault={handleSubmit} class="space-y-6">
						{#if error}
							<div class="flex items-center space-x-2 p-3 bg-red-50 border border-red-200 rounded-lg">
								<AlertCircle class="h-5 w-5 text-red-500" />
								<p class="text-red-800 text-sm">{error}</p>
							</div>
						{/if}

						<div class="space-y-4">
							<div>
								<label for={nameId} class="block text-sm font-medium text-gray-700 mb-2">
									Nama Lengkap *
								</label>
								<input 
									id={nameId}
									type="text"
									bind:value={name}
									class="w-full px-4 py-3 border border-gray-200 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent transition-all duration-200 bg-gray-50 focus:bg-white"
									placeholder="Masukkan nama lengkap"
									required
								/>
							</div>

							<div>
								<label for={emailId} class="block text-sm font-medium text-gray-700 mb-2">
									Email *
								</label>
								<input 
									id={emailId}
									type="email"
									bind:value={email}
									class="w-full px-4 py-3 border border-gray-200 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent transition-all duration-200 bg-gray-50 focus:bg-white"
									placeholder="contoh@email.com"
									required
								/>
							</div>

							<div>
								<label for={cvId} class="block text-sm font-medium text-gray-700 mb-2">
									CV (PDF) *
								</label>
								<div class="relative">
									<input 
										id={cvId}
										type="file"
										accept=".pdf"
										on:change={handleFileChange}
										class="w-full px-4 py-3 border border-gray-200 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent transition-all duration-200 file:mr-4 file:py-2 file:px-4 file:rounded-lg file:border-0 file:text-sm file:font-medium file:bg-blue-50 file:text-blue-700 hover:file:bg-blue-100 bg-gray-50 focus:bg-white"
										required
									/>
									{#if cvFile}
										<div class="mt-2 flex items-center space-x-2 text-sm text-green-600">
											<CheckCircle class="h-4 w-4" />
											<span>{cvFile.name}</span>
										</div>
									{/if}
								</div>
								<p class="text-xs text-gray-500 mt-1">
									Maksimal 5MB, format PDF
								</p>
							</div>
						</div>

						<div class="flex space-x-3 pt-4">
							<button 
								type="button"
								on:click={closeModal}
								class="flex-1 px-4 py-3 bg-gray-100 hover:bg-gray-200 text-gray-700 font-medium rounded-lg transition-all duration-200"
								disabled={isLoading}
							>
								Batal
							</button>
							<button 
								type="submit"
								class="flex-1 px-4 py-3 bg-blue-600 hover:bg-blue-700 text-white font-medium rounded-lg transition-all duration-200 flex items-center justify-center space-x-2"
								disabled={isLoading}
							>
								{#if isLoading}
									<div class="animate-spin rounded-full h-4 w-4 border-b-2 border-white"></div>
									<span>Mengirim...</span>
								{:else}
									<Upload class="h-4 w-4" />
									<span>Kirim Lamaran</span>
								{/if}
							</button>
						</div>
					</form>
				{/if}
			</div>
		</div>
	</div>
{/if} 