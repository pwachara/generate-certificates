<script lang="ts">
	import { goto } from '$app/navigation'
	import { pb } from '$lib/pocketbase'
  import "../../../app.css"

	let email = $state("")
  let dialog: HTMLDialogElement
  let error = $state("")

  function displaySentMessage(){
      dialog.showModal()
      
  }

	async function submitEmail() {

		await pb.collection('users').requestPasswordReset(email)
					.then(r => {
						return goto("/login", { invalidateAll: true })
					})
					.catch(e => {
							return error = "Error " + e.status + " - " + e.message
						}
					)	
	}

	function closeDialog() {
      submitEmail()
		  dialog.close()
	}

</script>

<div class="bg-white min-h-screen flex items-center justify-center p-4">
  <section class="w-full max-w-md mx-auto">
    
    <!-- Header -->
    <header class="text-center mb-8">
      <h2 class="text-2xl font-medium text-gray-700">Forgot Password</h2>
      <p class="text-xs font-medium text-gray-600 mt-4">Enter the email associated with your account and we’ll send you a recovery link</p>
    </header>
    
    <!-- Form -->
    <div>
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
      </form>
    </div>
    
    <!-- Footer -->
    <footer class="mt-6">
      <div class="flex flex-col gap-3">
        <button
          onclick={ displaySentMessage } 
          type="submit" 
          class="w-full bg-black text-white py-3 rounded-md font-medium hover:cursor-pointer hover:transform hover:scale-x-105 transition-colors">
          Send
        </button>
        
        <a href="/login" class="w-full text-center bg-gray-200 text-gray-700 py-3 rounded-md font-medium hover:bg-gray-300 transition-colors">
          Cancel
        </a>
      </div>
    </footer>
  </section>
</div>

<!-- Dialog -->
<dialog id="dialog" bind:this={dialog} class="w-full max-w-sm rounded-lg p-6">
  <div class="flex flex-col items-center gap-3 mb-5 text-center">
    <h3 class="font-bold text-xl text-gray-900">Email Sent</h3>
    <p class="text-base text-gray-800 font-medium">But wait!</p>
    <p class="text-sm text-gray-600">Check your email only if you are already registered</p>
  </div>
  <button 
    class="w-full bg-black text-white py-3 rounded-md font-medium hover:transform hover:scale-x-105 transition-colors" 
    onclick={ closeDialog }>
    Close
  </button>
</dialog>



<style>
    dialog {
        margin: auto;           /* removes default UA margin */
        inset: 0;               /* ensures it can be centered by flex */
        position: static;       /* prevents browser from forcing its own positioning */
        border: none;
        border-radius: 0.5rem;
        padding: 1.5rem;
        background: white;
        animation: fade-out 0.7s ease-out;
    }

    dialog:open {
        animation: fade-in 0.7s ease-out;
    }

    dialog:open::backdrop {
        background-color: black;
        animation: backdrop-fade-in 0.7s ease-out forwards;
    }

    /* Animations */
    @keyframes fade-in {
        0% { opacity: 0; transform: scaleY(0); }
        100% { opacity: 1; transform: scaleY(1); }
    }

    @keyframes fade-out {
        0% { opacity: 1; transform: scaleY(1); }
        100% { opacity: 0; transform: scaleY(0); }
    }

    @keyframes backdrop-fade-in {
        0% { opacity: 0; }
        100% { opacity: 0.5; }
    }


</style>

