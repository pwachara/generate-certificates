<script lang="ts">
	import { goto } from '$app/navigation'
	import { pb } from '$lib/pocketbase'

	let email = $state("")
	let password = $state("")
	let showPassword = $state(false)
	let error = $state("")

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

<div class="bg-gray-100 min-h-screen flex items-center justify-center p-4">
    <section class="bg-white rounded-lg shadow-md w-full md:w-3/5 max-w-2xl mx-auto my-16 md:my-36">
        <!-- Header -->
        <header class="border-b border-gray-200 py-4">
            <h2 class="text-2xl font-bold text-center">Enter Email Address</h2>
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
                    type="submit" 
                    class="px-6 py-2 bg-primary-600 text-white font-medium rounded-md hover:bg-primary-700 focus:outline-none focus:ring-2 focus:ring-primary-500 focus:ring-offset-2 transition-colors">
                    Send
                </button>
                <a href="/login" class="px-6 py-2 bg-secondary-500 text-white font-medium rounded-md hover:bg-secondary-600 focus:outline-none focus:ring-2 focus:ring-secondary-500 focus:ring-offset-2 transition-colors text-center">
                    Cancel
                </a>
            </div>
            
        </footer>
    </section>
</div>