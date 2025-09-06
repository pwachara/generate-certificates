<script lang="ts">
	import { goto } from '$app/navigation'
	import { pb } from '$lib/pocketbase'
	import { onMount } from 'svelte';

	let email = $state("")
	let password = $state("")
	let showPassword = $state(false)
	let error = $state("")

	onMount(() => {
        // Client-side authentication check
        if (pb.authStore.isValid) {
            goto('/', { invalidateAll: true });
        }
    });

	function toggleShowPassword(){

		showPassword = !showPassword

	}

	async function loginUser() {

		await pb.collection('users').authWithPassword(email, password)
					.then(r => {
						return goto("/", { invalidateAll: true })
					})
					.catch(e => {
						if(e.status === 400){

							return error = "Wrong email or password"

						} else {
							return error = "Error " + e.status + " - " + e.message
						}
					})	
	}
</script>

<!-- <section class="card md:w-3/5 mx-auto my-36">
					
	<header class="card-header text-center h4">Login</header>
	<div class="p-4">
		<form class="container mx-auto flex flex-col gap-4">
			<fieldset class="container mx-auto flex flex-row gap-2">
				<label for="email" class="label place-self-center"> Email: </label>
				<input
					class="input"
					type="email"
					title="Email"
					name="email"
					id="email"
					placeholder="Email"
					bind:value={email}
				/>
			</fieldset>
			<fieldset class="container mx-auto flex flex-row gap-2">
				{#if !showPassword}
					<label for="password" class="label place-self-center"> Password: </label>
					<input
						class="input"
						type="password"
						title="Password"
						name="password"
						id="password"
						placeholder="Password"
						bind:value={password}
					/>
				{:else}
					<label for="password" class="label place-self-center"> Password: </label>
					<input
						class="input"
						type="text"
						title="Password"
						name="password"
						id="password"
						placeholder="Password"
						bind:value={password}
					/>
				{/if}
				<button class="flex items-center" onclick={toggleShowPassword} >
					{#if showPassword}
						<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-6 h-6">
							<path stroke-linecap="round" stroke-linejoin="round" d="M3.98 8.223A10.477 10.477 0 001.934 12C3.226 16.338 7.244 19.5 12 19.5c.993 0 1.953-.138 2.863-.395M6.228 6.228A10.45 10.45 0 0112 4.5c4.756 0 8.773 3.162 10.065 7.498a10.523 10.523 0 01-4.293 5.774M6.228 6.228L3 3m3.228 3.228l3.65 3.65m7.894 7.894L21 21m-3.228-3.228l-3.65-3.65m0 0a3 3 0 10-4.243-4.243m4.242 4.242L9.88 9.88" />
						</svg>
					{:else}
						<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-6 h-6">
							<path stroke-linecap="round" stroke-linejoin="round" d="M2.036 12.322a1.012 1.012 0 010-.639C3.423 7.51 7.36 4.5 12 4.5c4.638 0 8.573 3.007 9.963 7.178.07.207.07.431 0 .639C20.577 16.49 16.64 19.5 12 19.5c-4.638 0-8.573-3.007-9.963-7.178z" />
							<path stroke-linecap="round" stroke-linejoin="round" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
						</svg>
					{/if}
				</button>
			</fieldset>
		</form>
	</div>
	<footer class="card-footer justify-around text-center">
		{#if error}
            <div class="flex justify-center items-center text-lg gap-2 py-2">
                <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-6 h-6 text-red-500">
                    <path stroke-linecap="round" stroke-linejoin="round" d="M12 9v3.75m9-.75a9 9 0 1 1-18 0 9 9 0 0 1 18 0Zm-9 3.75h.008v.008H12v-.008Z" />
                </svg>
                <p class="text-red-500 font-medium">{ error }</p>
            </div>
        {/if}
		<div>
			<button 
				onclick={loginUser}
				type="submit" 
				class="btn variant-filled-primary">
					Login
			</button>
			<a href="/" class="btn variant-filled-secondary">Cancel</a>
		</div>
	</footer>
</section> -->

<div class="bg-gray-100 min-h-screen flex items-center justify-center p-4">
    <section class="bg-white rounded-lg shadow-md w-full md:w-3/5 max-w-2xl mx-auto my-16 md:my-36">
        <!-- Header -->
        <header class="border-b border-gray-200 py-4">
            <h2 class="text-2xl font-bold text-center">Login</h2>
        </header>
        
        <div class="p-6">
            <form class="flex flex-col gap-5">
                <!-- Email Field -->
                <div class="flex flex-col sm:flex-row gap-3 items-center">
                    <label for="email" class="w-24 font-medium text-gray-700">Email:</label>
                    <input
                        class="flex-1 px-4 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-primary-500 focus:border-transparent"
                        type="email"
                        title="Email"
                        name="email"
                        id="email"
                        placeholder="Email"
						bind:value={email}
                    />
                </div>
                
                <!-- Password Field -->
                <div class="flex flex-col sm:flex-row gap-3 items-center">
                    <label for="password" class="w-24 font-medium text-gray-700">Password:</label>
                    <div class="flex-1 flex items-center relative">
                        <input
                            class="w-full px-4 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-primary-500 focus:border-transparent"
                            type={showPassword? "text" : "password"}
                            title="Password"
                            name="password"
                            id="password"
                            placeholder="Password"
							bind:value={password}
                        />
                        <button type="button" class="absolute right-3 text-gray-500 hover:text-gray-700" onclick={toggleShowPassword}>
							{#if showPassword}
								<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-6 h-6">
									<path stroke-linecap="round" stroke-linejoin="round" d="M3.98 8.223A10.477 10.477 0 001.934 12C3.226 16.338 7.244 19.5 12 19.5c.993 0 1.953-.138 2.863-.395M6.228 6.228A10.45 10.45 0 0112 4.5c4.756 0 8.773 3.162 10.065 7.498a10.523 10.523 0 01-4.293 5.774M6.228 6.228L3 3m3.228 3.228l3.65 3.65m7.894 7.894L21 21m-3.228-3.228l-3.65-3.65m0 0a3 3 0 10-4.243-4.243m4.242 4.242L9.88 9.88" />
								</svg>
							{:else}
								<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-6 h-6">
									<path stroke-linecap="round" stroke-linejoin="round" d="M2.036 12.322a1.012 1.012 0 010-.639C3.423 7.51 7.36 4.5 12 4.5c4.638 0 8.573 3.007 9.963 7.178.07.207.07.431 0 .639C20.577 16.49 16.64 19.5 12 19.5c-4.638 0-8.573-3.007-9.963-7.178z" />
									<path stroke-linecap="round" stroke-linejoin="round" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
								</svg>
							{/if}
                        </button>
                    </div>
                </div>
            </form>
        </div>
        
        <!-- Footer -->
        <footer class="border-t border-gray-200 py-4 px-6">
            <!-- Error message (hidden by default) -->
            <div class="hidden justify-center items-center text-lg gap-2 py-2 mb-4 error-container">
                <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-6 h-6 text-red-500">
                    <path stroke-linecap="round" stroke-linejoin="round" d="M12 9v3.75m9-.75a9 9 0 1 1-18 0 9 9 0 0 1 18 0Zm-9 3.75h.008v.008H12v-.008Z" />
                </svg>
                <p class="text-red-500 font-medium">Error message would appear here</p>
            </div>
            
            <div class="flex flex-col sm:flex-row justify-center gap-4">
                <button 
					onclick={loginUser}
                    type="submit" 
                    class="px-6 py-2 bg-primary-600 text-white font-medium rounded-md hover:bg-primary-700 focus:outline-none focus:ring-2 focus:ring-primary-500 focus:ring-offset-2 transition-colors">
                    Login
                </button>
                <a href="/" class="px-6 py-2 bg-secondary-500 text-white font-medium rounded-md hover:bg-secondary-600 focus:outline-none focus:ring-2 focus:ring-secondary-500 focus:ring-offset-2 transition-colors text-center">
                    Cancel
                </a>
            </div>
            
            <div class="mt-6 text-center text-sm text-gray-500">
                <p>Forgot Password? <a href="/forgot-password" class="text-primary-600 hover:underline">Click here</a></p>
            </div>
        </footer>
    </section>
</div>