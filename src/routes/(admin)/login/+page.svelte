<script lang="ts">
	import "../../../app.css"
	import { goto } from '$app/navigation'
	import { pb } from '$lib/pocketbase'
	import { onMount } from 'svelte';

	let email = $state("")
	let password = $state("")
	let showPassword = $state(false)
	let error = $state("")
	let processing = $state(false)

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

		processing = true

		if(email === "" || password === ""){
			processing = false
			return error = "email and password are required fields"
		}

		await pb.collection('users').authWithPassword(email, password)
					.then(r => {

						processing = false

						return goto("/", { invalidateAll: true })
					})
					.catch(e => {

						processing = false

						if(e.status === 400){

							return error = "Wrong email or password"

						} else {

							return error = "Error " + e.status + " - " + e.message
						}
					})	
	}
</script>

<div class="bg-white min-h-screen flex items-center justify-center p-4">
  <section class="w-full max-w-md mx-auto">
    
    <!-- Header -->
    <header class="text-center mb-8">
      <h2 class="text-2xl font-bold text-gray-900">Login</h2>
    </header>

    <!-- Form -->
    <form class="flex flex-col gap-4">
      
      <!-- Email Field -->
      <div class="flex flex-col">
        <input
          class="w-full px-4 py-3 rounded-md bg-gray-100 placeholder-gray-500 focus:outline-none focus:ring-1 focus:ring-black"
          type="email"
          title="Email"
          name="email"
          id="email"
          placeholder="Email*"
          bind:value={email}
          required
        />
      </div>
      
      <!-- Password Field -->
      <div class="flex flex-col relative">
        <input
          class="w-full px-4 py-3 rounded-md bg-gray-100 placeholder-gray-500 focus:outline-none focus:ring-1 focus:ring-black"
          type={showPassword ? "text" : "password"}
          title="Password"
          name="password"
          id="password"
          placeholder="Password*"
          bind:value={password}
          required
        />
        <button 
          type="button" 
          class="absolute right-3 top-3 text-gray-400 hover:text-gray-600" 
          onclick={toggleShowPassword}>
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

    </form>

    <!-- Footer -->
    <footer class="mt-6">
      <!-- Error message -->
      {#if error}
        <div class="flex justify-center items-center text-sm gap-2 mb-4">
          <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-5 h-5 text-red-500">
            <path stroke-linecap="round" stroke-linejoin="round" d="M12 9v3.75m9-.75a9 9 0 1 1-18 0 9 9 0 0 1 18 0Zm-9 3.75h.008v.008H12v-.008Z" />
          </svg>
          <p class="text-red-500 font-medium">Error: { error }</p>
        </div>
      {/if}

      <div class="flex flex-col gap-3">
        <button 
          onclick={loginUser}
          type="submit" 
          class="w-full bg-black text-white py-3 rounded-md font-medium hover:transform hover:scale-x-105 transition-colors hover:cursor-pointer">
          {#if processing}
            Processing...
          {:else}
            Login
          {/if}
        </button>

        <a href="/" class="w-full text-center bg-gray-200 text-gray-700 py-3 rounded-md font-medium hover:bg-gray-300 transition-colors">
          Clear
        </a>
      </div>
      
      <div class="mt-6 text-center text-xs text-gray-500">
        <p>Forgot Password? <a href="/forgot-password" class="text-black hover:underline">Click here</a></p>
      </div>
    </footer>
  </section>
</div>
