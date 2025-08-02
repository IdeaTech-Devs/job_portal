# Job Portal Frontend

Frontend untuk aplikasi Job Portal menggunakan SvelteKit dengan TypeScript dan Tailwind CSS.

## Setup

### Prerequisites
- Node.js 18+
- npm atau yarn

### Install Dependencies
```bash
npm install
```

### Development
```bash
npm run dev
```

Aplikasi akan berjalan di `http://localhost:5173`

### Build
```bash
npm run build
```

### Preview
```bash
npm run preview
```

## Fitur

### 1. Halaman Utama (Job Listings)
- Menampilkan daftar lowongan dalam grid responsive
- Filter berdasarkan lokasi dan rentang gaji
- Real-time filtering tanpa reload halaman
- Loading states dan error handling

### 2. Filter System
- Filter lokasi (dropdown)
- Filter rentang gaji (radio buttons)
- Clear filters functionality

### 3. Apply Modal
- Form lamaran dengan validasi
- Upload CV (PDF only, max 5MB)
- Success feedback
- Error handling

### 4. Halaman Applications
- Daftar lamaran yang telah diajukan
- Informasi detail lowongan
- Status lamaran
- Persistensi data di localStorage

## Struktur Komponen

- `JobCard.svelte` - Kartu lowongan individual
- `FilterPanel.svelte` - Panel filter lowongan
- `ApplyModal.svelte` - Modal form lamaran
- `+page.svelte` - Halaman utama
- `applications/+page.svelte` - Halaman lamaran

## API Integration

- `api.ts` - Client untuk berkomunikasi dengan backend
- `applications.ts` - Store untuk state management
- Error handling dan loading states

## Styling

- Tailwind CSS untuk styling
- Responsive design (mobile-first)
- Custom components dengan utility classes
- Modern UI dengan smooth transitions 