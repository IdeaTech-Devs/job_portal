<script lang="ts">
	import '../app.css';
	import { Briefcase, FileText, Search, User, Menu, X } from 'lucide-svelte';

	let mobileMenuOpen = false;

	function toggleMobileMenu() {
		mobileMenuOpen = !mobileMenuOpen;
	}
</script>

<svelte:head>
	<title>Job Portal - Temukan Pekerjaan Impian Anda</title>
	<meta name="description" content="Portal lowongan kerja terbaik dengan berbagai posisi menarik dari perusahaan ternama" />
</svelte:head>

<div class="min-h-screen bg-gray-50">
	<!-- Navigation -->
	<nav class="bg-white border-b border-gray-200 sticky top-0 z-50">
		<div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
			<div class="flex justify-between h-16">
				<!-- Logo -->
				<div class="flex items-center">
					<a href="/" class="flex items-center space-x-3 group">
						<div class="p-2 bg-blue-600 rounded-lg group-hover:scale-110 transition-transform duration-200">
							<Briefcase class="h-5 w-5 text-white" />
						</div>
						<div class="hidden sm:block">
							<span class="text-xl font-bold text-gray-900">
								Job Portal
							</span>
							<p class="text-xs text-gray-500 -mt-1">Temukan Karir Impian</p>
						</div>
						<div class="sm:hidden">
							<span class="text-lg font-bold text-gray-900">
								Job Portal
							</span>
						</div>
					</a>
				</div>

				<!-- Desktop Navigation -->
				<div class="hidden md:flex items-center space-x-6">
					<a href="/" class="flex items-center space-x-2 text-gray-700 hover:text-blue-600 px-4 py-2 rounded-lg hover:bg-blue-50 transition-all duration-200 group">
						<Search class="h-4 w-4" />
						<span class="font-medium">Lowongan</span>
					</a>
					<a href="/applications" class="flex items-center space-x-2 text-gray-700 hover:text-blue-600 px-4 py-2 rounded-lg hover:bg-blue-50 transition-all duration-200 group">
						<FileText class="h-4 w-4" />
						<span class="font-medium">Lamaran Saya</span>
					</a>
					<div class="flex items-center space-x-2 px-4 py-2 bg-blue-600 rounded-lg text-white">
						<User class="h-4 w-4" />
						<span class="text-sm font-medium">User</span>
					</div>
				</div>

				<!-- Mobile menu button -->
				<div class="md:hidden flex items-center">
					<button 
						on:click={toggleMobileMenu}
						class="p-2 rounded-lg text-gray-700 hover:text-blue-600 hover:bg-blue-50 transition-colors"
					>
						{#if mobileMenuOpen}
							<X class="h-6 w-6" />
						{:else}
							<Menu class="h-6 w-6" />
						{/if}
					</button>
				</div>
			</div>

			<!-- Mobile Navigation -->
			{#if mobileMenuOpen}
				<div class="md:hidden border-t border-gray-200 bg-white">
					<div class="px-2 pt-2 pb-3 space-y-1">
						<a 
							href="/" 
							class="flex items-center space-x-3 px-3 py-2 text-gray-700 hover:text-blue-600 hover:bg-blue-50 rounded-lg transition-colors"
							on:click={() => mobileMenuOpen = false}
						>
							<Search class="h-4 w-4" />
							<span class="font-medium">Lowongan</span>
						</a>
						<a 
							href="/applications" 
							class="flex items-center space-x-3 px-3 py-2 text-gray-700 hover:text-blue-600 hover:bg-blue-50 rounded-lg transition-colors"
							on:click={() => mobileMenuOpen = false}
						>
							<FileText class="h-4 w-4" />
							<span class="font-medium">Lamaran Saya</span>
						</a>
						<div class="flex items-center space-x-3 px-3 py-2 bg-blue-600 rounded-lg text-white">
							<User class="h-4 w-4" />
							<span class="font-medium">User</span>
						</div>
					</div>
				</div>
			{/if}
		</div>
	</nav>

	<!-- Main Content -->
	<main>
		<slot />
	</main>

	<!-- Footer -->
	<footer class="bg-white border-t border-gray-200 mt-16">
		<div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
			<div class="text-center">
				<div class="flex justify-center items-center space-x-2 mb-4">
					<Briefcase class="h-5 w-5 text-blue-600" />
					<span class="text-lg font-semibold text-gray-900">Job Portal</span>
				</div>
				<p class="text-gray-500 text-sm mb-6">
					© 2025 Job Portal. Untuk keperluan technical test Roomah.
				</p>
				<div class="flex flex-col sm:flex-row justify-center space-y-2 sm:space-y-0 sm:space-x-8 text-sm text-gray-400">
					<button class="hover:text-blue-600 transition-colors">Tentang Kami</button>
					<button class="hover:text-blue-600 transition-colors">Kebijakan Privasi</button>
					<button class="hover:text-blue-600 transition-colors">Syarat & Ketentuan</button>
					<button class="hover:text-blue-600 transition-colors">Kontak</button>
				</div>
			</div>
		</div>
	</footer>
</div> 