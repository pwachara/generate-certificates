<script lang="ts">
	import '../../app.css';
	import favicon from '$lib/assets/favicon.svg';
	import { pb } from "$lib/pocketbase"
	import { goto} from '$app/navigation';
    import { page } from '$app/state';

	let { children } = $props();

	const user = pb.authStore.record

	let fullName = user?.name

	let nameInitials = getInitials(fullName)

	function getInitials(fullName){
		if (!fullName || typeof fullName !== 'string') return '';
        
        const names = fullName.trim().split(/\s+/);
        
        if (names.length === 0) return '';
        if (names.length === 1) return names[0].charAt(0).toUpperCase();
        
        const firstName = names[0];
        const lastName = names[names.length - 1];
        
        return (firstName.charAt(0) + lastName.charAt(0)).toUpperCase();
	}

	function logout() {
		pb.authStore.clear()
        goto("/login", {invalidateAll:true} )
	}
    
</script>

<svelte:head>
	<link rel="icon" href={favicon} />
</svelte:head>

<section class="bg-gray-100 flex h-screen">
    <!-- Mobile overlay -->
    <div class="overlay" id="overlay"></div>

    <!-- Sidebar -->
    <aside class="sidebar bg-gray-800 text-white w-64 flex flex-col">
        <!-- App Title -->
        <div class="p-4 flex items-center justify-between border-b border-gray-700">
            <div class="flex items-center">
                <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" width="36" height="36" fill="rgba(240,187,64,1)"><path d="M4 11.3333L0 9L12 2L24 9V17.5H22V10.1667L20 11.3333V18.0113L19.7774 18.2864C17.9457 20.5499 15.1418 22 12 22C8.85817 22 6.05429 20.5499 4.22263 18.2864L4 18.0113V11.3333ZM6 12.5V17.2917C7.46721 18.954 9.61112 20 12 20C14.3889 20 16.5328 18.954 18 17.2917V12.5L12 16L6 12.5ZM3.96927 9L12 13.6846L20.0307 9L12 4.31541L3.96927 9Z"></path></svg>
                <h1 class="text-xl font-poppins font-semibold ml-2">Loka Certificates</h1>
            </div>
            <button id="close-sidebar" class="md:hidden text-gray-400 hover:text-white" aria-label="button">
                <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
                </svg>
            </button>
        </div>

        <!-- Navigation Links -->
        <nav class="flex-1 overflow-y-auto py-4">
            <ul>
                <li class="mb-1">
                    <a class:active={page.url.pathname === "/"}  href="/" class="flex items-center px-4 py-3  text-gray-300 hover:bg-gray-700 hover:text-white">
                        <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 12l2-2m0 0l7-7 7 7M5 10v10a1 1 0 001 1h3m10-11l2 2m-2-2v10a1 1 0 01-1 1h-3m-6 0a1 1 0 001-1v-4a1 1 0 011-1h2a1 1 0 011 1v4a1 1 0 001 1m-6 0h6" />
                        </svg>
                        <span class="ml-3">Dashboard</span>
                    </a>
                </li>
                <li class="mb-1">
                    <a class:active={page.url.pathname === "/trainees"} href="/trainees" class="flex items-center px-4 py-3 text-gray-300 hover:bg-gray-700 hover:text-white">
                        <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0zm6 3a2 2 0 11-4 0 2 2 0 014 0zM7 10a2 2 0 11-4 0 2 2 0 014 0z" />
                        </svg>
                        <span class="ml-3">Trainees</span>
                    </a>
                </li>
                <li class="mb-1">
                    <a class:active={page.url.pathname === "/organizations"}  href="/organizations" class="flex items-center px-4 py-3 text-gray-300 hover:bg-gray-700 hover:text-white">
                        <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 9V7a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2m2 4h10a2 2 0 002-2v-6a2 2 0 00-2-2H9a2 2 0 00-2 2v6a2 2 0 002 2zm7-5a2 2 0 11-4 0 2 2 0 014 0z" />
                        </svg>
                        <span class="ml-3">Organizations</span>
                    </a>
                </li>
                <li class="mb-1">
                    <a class:active={page.url.pathname === "/courses"} href="/courses" class="flex items-center px-4 py-3 text-gray-300 hover:bg-gray-700 hover:text-white">
                        <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" viewBox="0 0 24 24" fill="currentColor"><path d="M17 15.2454V22.1169C17 22.393 16.7761 22.617 16.5 22.617C16.4094 22.617 16.3205 22.5923 16.2428 22.5457L12 20L7.75725 22.5457C7.52046 22.6877 7.21333 22.6109 7.07125 22.3742C7.02463 22.2964 7 22.2075 7 22.1169V15.2454C5.17107 13.7793 4 11.5264 4 9C4 4.58172 7.58172 1 12 1C16.4183 1 20 4.58172 20 9C20 11.5264 18.8289 13.7793 17 15.2454ZM9 16.4185V19.4676L12 17.6676L15 19.4676V16.4185C14.0736 16.7935 13.0609 17 12 17C10.9391 17 9.92643 16.7935 9 16.4185ZM12 15C15.3137 15 18 12.3137 18 9C18 5.68629 15.3137 3 12 3C8.68629 3 6 5.68629 6 9C6 12.3137 8.68629 15 12 15Z"></path></svg>
                        <span class="ml-3">Courses</span>
                    </a>
                </li>
                <li class="mb-1">
                    <a href="/" class="flex items-center px-4 py-3 text-gray-300 hover:bg-gray-700 hover:text-white">
                        <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 3h2l.4 2M7 13h10l4-8H5.4M7 13L5.4 5M7 13l-2.293 2.293c-.63.63-.184 1.707.707 1.707H17m0 0a2 2 0 100 4 2 2 0 000-4zm-8 2a2 2 0 11-4 0 2 2 0 014 0z" />
                        </svg>
                        <span class="ml-3">Purchases</span>
                    </a>
                </li>
                <li class="mb-1">
                    <a href="/" class="flex items-center px-4 py-3 text-gray-300 hover:bg-gray-700 hover:text-white">
                        <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 11V7a4 4 0 00-8 0v4M5 9h14l1 12H4L5 9z" />
                        </svg>
                        <span class="ml-3">Sales</span>
                    </a>
                </li>
                <li class="mb-1">
                    <a href="/" class="flex items-center px-4 py-3 text-gray-300 hover:bg-gray-700 hover:text-white">
                        <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m5.618-4.016A11.955 11.955 0 0112 2.944a11.955 11.955 0 01-8.618 3.04A12.02 12.02 0 003 9c0 5.591 3.824 10.29 9 11.622 5.176-1.332 9-6.03 9-11.622 0-1.042-.133-2.052-.382-3.016z" />
                        </svg>
                        <span class="ml-3">Receipts</span>
                    </a>
                </li>
                <li class="mb-1">
                    <a href="/" class="flex items-center px-4 py-3 text-gray-300 hover:bg-gray-700 hover:text-white">
                        <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
                        </svg>
                        <span class="ml-3">Payments</span>
                    </a>
                </li>
                <li class="mb-1">
                    <a href="/" class="flex items-center px-4 py-3 text-gray-300 hover:bg-gray-700 hover:text-white">
                        <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 17v-2m3 2v-4m3 4v-6m2 10H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z" />
                        </svg>
                        <span class="ml-3">Reports</span>
                    </a>
                </li>
            </ul>
        </nav>

        <!-- User Profile -->
        <div class="p-4 border-t border-gray-700">
            <div class="flex items-center">
                <div class="w-8 h-8 rounded-full bg-yellow-500 flex items-center justify-center text-white font-semibold">
                    { nameInitials }
                </div>
                <div class="ml-3">
                    <p class="text-sm font-medium">{ user?.name}</p>
                    <p class="text-xs text-gray-400">Admin</p>
                </div>
            </div>
        </div>
    </aside>

	    <!-- Main Content -->
    <div class="flex-1 flex flex-col overflow-hidden">
        <!-- Top Bar -->
        <header class="bg-white border-b border-gray-200">
            <div class="flex items-center justify-between p-4">
                <div class="flex items-center">
                    <button id="menu-toggle" class="md:hidden text-gray-600 hover:text-gray-900" aria-label="button">
                        <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16M4 18h16" />
                        </svg>
                    </button>
                    <h2 class="ml-2 md:ml-0 text-xl font-poppins font-semibold text-gray-800">Dashboard</h2>
                </div>
                {#if pb.authStore.isValid}
                    <div class="relative flex items-center space-x-4 hover:cursor-pointer">                    
                        <!-- User Menu -->
                        <div class="relative">
                                <button class="flex items-center focus:outline-none hover:cursor-pointer" popovertarget="dropdown" popovertargetaction="toggle">
                                    <div class="w-8 h-8 rounded-full bg-yellow-500 flex items-center justify-center text-white font-semibold">
                                        {nameInitials}
                                    </div>
                                    <span class="ml-2 text-gray-700 font-medium hidden md:block">{ fullName }</span>
                                </button>
                                <div id="dropdown" popover class="absolute right-0 top-16 w-48">
                                    <ul>
                                        <li>
                                            <button class="bg-gray-300 font-semibold w-full flex items-center justify-center px-4 py-2 hover:cursor-pointer" onclick={ logout }>Logout</button>
                                        </li>
                                    </ul>
                                </div>
                        </div>
                    </div>
                {/if}
            </div>
        </header>

		{@render children?.()}

    </div>


</section>

<style>
	.active{
		background-color: #fbbf24;
		font-weight: bold;
		color: white;
	}

       /* Override popover positioning */
    [popover] {
      position: absolute;
      inset: auto 0 auto auto; /* anchor below, aligned right */
      margin-top: 4rem;
      border: 1px solid #ccc;
      border-radius: 8px;
      padding: 0.5rem 0;
      background: white;
      box-shadow: 0 4px 12px rgba(0,0,0,0.15);
      min-width: 150px;
    }

    [popover] ul {
      list-style: none;
      margin: 0;
      padding: 0;
    }

    [popover] li {
      padding: 0.5rem 1rem;
      cursor: pointer;
    }

    [popover] li:hover {
      background: #f2f2f2;
    }


</style>