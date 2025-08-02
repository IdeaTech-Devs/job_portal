import { c as create_ssr_component, v as validate_component, e as each, a as escape, m as missing_component } from "../../chunks/ssr.js";
import { B as Briefcase, S as Search, T as Trending_up } from "../../chunks/trending-up.js";
import { I as Icon } from "../../chunks/Icon.js";
import { F as File_text } from "../../chunks/file-text.js";
const Menu = create_ssr_component(($$result, $$props, $$bindings, slots) => {
  const iconNode = [
    [
      "line",
      {
        "x1": "4",
        "x2": "20",
        "y1": "12",
        "y2": "12"
      }
    ],
    [
      "line",
      {
        "x1": "4",
        "x2": "20",
        "y1": "6",
        "y2": "6"
      }
    ],
    [
      "line",
      {
        "x1": "4",
        "x2": "20",
        "y1": "18",
        "y2": "18"
      }
    ]
  ];
  return `${validate_component(Icon, "Icon").$$render($$result, Object.assign({}, { name: "menu" }, $$props, { iconNode }), {}, {
    default: () => {
      return `${slots.default ? slots.default({}) : ``}`;
    }
  })}`;
});
const User = create_ssr_component(($$result, $$props, $$bindings, slots) => {
  const iconNode = [
    [
      "path",
      {
        "d": "M19 21v-2a4 4 0 0 0-4-4H9a4 4 0 0 0-4 4v2"
      }
    ],
    ["circle", { "cx": "12", "cy": "7", "r": "4" }]
  ];
  return `${validate_component(Icon, "Icon").$$render($$result, Object.assign({}, { name: "user" }, $$props, { iconNode }), {}, {
    default: () => {
      return `${slots.default ? slots.default({}) : ``}`;
    }
  })}`;
});
const Users = create_ssr_component(($$result, $$props, $$bindings, slots) => {
  const iconNode = [
    [
      "path",
      {
        "d": "M16 21v-2a4 4 0 0 0-4-4H6a4 4 0 0 0-4 4v2"
      }
    ],
    ["circle", { "cx": "9", "cy": "7", "r": "4" }],
    ["path", { "d": "M22 21v-2a4 4 0 0 0-3-3.87" }],
    ["path", { "d": "M16 3.13a4 4 0 0 1 0 7.75" }]
  ];
  return `${validate_component(Icon, "Icon").$$render($$result, Object.assign({}, { name: "users" }, $$props, { iconNode }), {}, {
    default: () => {
      return `${slots.default ? slots.default({}) : ``}`;
    }
  })}`;
});
const SplashScreen = create_ssr_component(($$result, $$props, $$bindings, slots) => {
  let currentStep = 0;
  const steps = [
    {
      icon: Briefcase,
      text: "Temukan Pekerjaan Impian",
      color: "text-blue-600"
    },
    {
      icon: Search,
      text: "Cari dengan Mudah",
      color: "text-green-600"
    },
    {
      icon: Users,
      text: "Terhubung dengan Perusahaan",
      color: "text-purple-600"
    },
    {
      icon: Trending_up,
      text: "Karir yang Menjanjikan",
      color: "text-orange-600"
    }
  ];
  return `${`<div class="fixed inset-0 z-50 flex items-center justify-center bg-gradient-to-br from-blue-600 via-purple-600 to-indigo-700"><div class="text-center text-white"> <div class="mb-8"><h1 class="text-4xl sm:text-6xl font-bold mb-4 tracking-tight" data-svelte-h="svelte-19n9xp0">Job Portal</h1> <p class="text-lg sm:text-xl text-blue-100 opacity-90" data-svelte-h="svelte-at7pkl">Platform Lowongan Kerja Terbaik</p></div>  <div class="space-y-4">${each(steps, (step, index) => {
    return `${index === currentStep ? `<div class="flex items-center justify-center space-x-3">${validate_component(step.icon || missing_component, "svelte:component").$$render($$result, { class: "h-6 w-6 " + step.color }, {}, {})} <span class="text-lg font-medium">${escape(step.text)}</span> </div>` : ``}`;
  })}</div>  <div class="mt-8 flex justify-center space-x-2">${each(Array(3), (_, i) => {
    return `<div class="w-2 h-2 bg-white rounded-full animate-pulse" style="${"animation-delay: " + escape(i * 0.2, true) + "s;"}"></div>`;
  })}</div></div>  <div class="absolute inset-0 overflow-hidden pointer-events-none"><div class="absolute -top-40 -right-40 w-80 h-80 bg-white opacity-10 rounded-full animate-pulse"></div> <div class="absolute -bottom-40 -left-40 w-80 h-80 bg-white opacity-10 rounded-full animate-pulse" style="animation-delay: 1s;"></div></div></div>`}`;
});
const Layout = create_ssr_component(($$result, $$props, $$bindings, slots) => {
  return `${validate_component(SplashScreen, "SplashScreen").$$render($$result, {}, {}, {})} ${$$result.head += `<!-- HEAD_svelte-1rw46st_START -->${$$result.title = `<title>Job Portal - Temukan Pekerjaan Impian Anda</title>`, ""}<meta name="description" content="Portal lowongan kerja terbaik dengan berbagai posisi menarik dari perusahaan ternama"><!-- HEAD_svelte-1rw46st_END -->`, ""} <div class="min-h-screen bg-gray-50"> <nav class="bg-white border-b border-gray-200 sticky top-0 z-50"><div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8"><div class="flex justify-between h-16"> <div class="flex items-center"><a href="/" class="flex items-center space-x-3 group"><div class="p-2 bg-blue-600 rounded-lg group-hover:scale-110 transition-transform duration-200">${validate_component(Briefcase, "Briefcase").$$render($$result, { class: "h-5 w-5 text-white" }, {}, {})}</div> <div class="hidden sm:block"><span class="text-xl font-bold text-gray-900" data-svelte-h="svelte-hg9mrt">Job Portal</span> <p class="text-xs text-gray-500 -mt-1" data-svelte-h="svelte-19it56f">Temukan Karir Impian</p></div> <div class="sm:hidden"><span class="text-lg font-bold text-gray-900" data-svelte-h="svelte-1tva0gw">Job Portal</span></div></a></div>  <div class="hidden md:flex items-center space-x-6"><a href="/" class="flex items-center space-x-2 text-gray-700 hover:text-blue-600 px-4 py-2 rounded-lg hover:bg-blue-50 transition-all duration-200 group">${validate_component(Search, "Search").$$render($$result, { class: "h-4 w-4" }, {}, {})} <span class="font-medium" data-svelte-h="svelte-6e6bbt">Lowongan</span></a> <a href="/applications" class="flex items-center space-x-2 text-gray-700 hover:text-blue-600 px-4 py-2 rounded-lg hover:bg-blue-50 transition-all duration-200 group">${validate_component(File_text, "FileText").$$render($$result, { class: "h-4 w-4" }, {}, {})} <span class="font-medium" data-svelte-h="svelte-1ozsz4q">Lamaran Saya</span></a> <div class="flex items-center space-x-2 px-4 py-2 bg-blue-600 rounded-lg text-white">${validate_component(User, "User").$$render($$result, { class: "h-4 w-4" }, {}, {})} <span class="text-sm font-medium" data-svelte-h="svelte-vi61yv">User</span></div></div>  <div class="md:hidden flex items-center"><button class="p-2 rounded-lg text-gray-700 hover:text-blue-600 hover:bg-blue-50 transition-colors">${`${validate_component(Menu, "Menu").$$render($$result, { class: "h-6 w-6" }, {}, {})}`}</button></div></div>  ${``}</div></nav>  <main>${slots.default ? slots.default({}) : ``}</main>  <footer class="bg-white border-t border-gray-200 mt-16"><div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8"><div class="text-center"><div class="flex justify-center items-center space-x-2 mb-4">${validate_component(Briefcase, "Briefcase").$$render($$result, { class: "h-5 w-5 text-blue-600" }, {}, {})} <span class="text-lg font-semibold text-gray-900" data-svelte-h="svelte-1baoh1j">Job Portal</span></div> <p class="text-gray-500 text-sm mb-6" data-svelte-h="svelte-h9kp9">© 2025 Job Portal. Untuk keperluan technical test Roomah.</p> <div class="flex flex-col sm:flex-row justify-center space-y-2 sm:space-y-0 sm:space-x-8 text-sm text-gray-400"><button class="hover:text-blue-600 transition-colors" data-svelte-h="svelte-1aqxne1">Tentang Kami</button> <button class="hover:text-blue-600 transition-colors" data-svelte-h="svelte-1o4vekc">Kebijakan Privasi</button> <button class="hover:text-blue-600 transition-colors" data-svelte-h="svelte-1xe67xp">Syarat &amp; Ketentuan</button> <button class="hover:text-blue-600 transition-colors" data-svelte-h="svelte-1jmzu7q">Kontak</button></div></div></div></footer></div>`;
});
export {
  Layout as default
};
