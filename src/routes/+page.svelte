<script>
  export let participantName = "John Doe"
  export let eventName = "Loka Advisory Training"
  export let date = "30th August 2025"
</script>

<div
  id="certificate"
  class="relative bg-white w-[1100px] h-[800px] mx-auto border-[12px] border-yellow-500 shadow-2xl flex flex-col items-center justify-center font-serif overflow-hidden"
>
  <!-- Inner border -->
  <div class="absolute inset-6 border-[6px] border-black pointer-events-none"></div>

  <!-- Corner Ornaments -->
  <!-- Top Left -->
  <svg
    class="absolute top-6 left-6 w-40 h-40 text-yellow-500 opacity-90"
    viewBox="0 0 200 200"
    fill="none"
    stroke="currentColor"
    stroke-width="3"
  >
    <path d="M10,190 Q60,140 40,100 T100,40 Q160,20 190,10" />
    <circle cx="40" cy="100" r="5" fill="currentColor" />
    <circle cx="100" cy="40" r="4" fill="currentColor" />
  </svg>

  <!-- Top Right -->
  <svg
    class="absolute top-6 right-6 w-40 h-40 text-yellow-500 opacity-90 rotate-90"
    viewBox="0 0 200 200"
    fill="none"
    stroke="currentColor"
    stroke-width="3"
  >
    <path d="M10,190 Q60,140 40,100 T100,40 Q160,20 190,10" />
    <circle cx="40" cy="100" r="5" fill="currentColor" />
    <circle cx="100" cy="40" r="4" fill="currentColor" />
  </svg>

  <!-- Bottom Left -->
  <svg
    class="absolute bottom-6 left-6 w-40 h-40 text-yellow-500 opacity-90 -rotate-90"
    viewBox="0 0 200 200"
    fill="none"
    stroke="currentColor"
    stroke-width="3"
  >
    <path d="M10,190 Q60,140 40,100 T100,40 Q160,20 190,10" />
    <circle cx="40" cy="100" r="5" fill="currentColor" />
    <circle cx="100" cy="40" r="4" fill="currentColor" />
  </svg>

  <!-- Bottom Right -->
  <svg
    class="absolute bottom-6 right-6 w-40 h-40 text-yellow-500 opacity-90 rotate-180"
    viewBox="0 0 200 200"
    fill="none"
    stroke="currentColor"
    stroke-width="3"
  >
    <path d="M10,190 Q60,140 40,100 T100,40 Q160,20 190,10" />
    <circle cx="40" cy="100" r="5" fill="currentColor" />
    <circle cx="100" cy="40" r="4" fill="currentColor" />
  </svg>

  <!-- Header / Logo -->
  <div class="absolute top-12 left-12 flex items-center gap-4">
    <img src="/logo-loka.webp" alt="Logo" class="h-20" />
  </div>

  <!-- Main Body -->
  <div class="z-10 text-center px-24">
    <h1 class="text-5xl font-bold text-black tracking-wide">CERTIFICATE</h1>
    <h2 class="text-xl text-gray-600 italic mt-1">of Participation</h2>

    <p class="mt-12 text-xl text-gray-700">This certificate is proudly presented to</p>

    <p class="mt-6 text-4xl font-bold text-yellow-600 italic">
      {participantName}
    </p>

    <p class="mt-8 text-lg text-gray-700">
      For successfully participating in
      <span class="font-semibold text-black">{eventName}</span>
      held on {date}.
    </p>
  </div>

  <!-- Footer -->
  <div class="absolute bottom-16 w-full flex justify-between px-24">
    <div class="text-center">
      <div class="border-t-2 border-black w-48 mx-auto"></div>
      <p class="mt-2 text-sm text-gray-700">Authorized Signatory</p>
    </div>

    <div class="text-center">
      <div class="border-t-2 border-black w-48 mx-auto"></div>
      <p class="mt-2 text-sm text-gray-700">Date</p>
    </div>
  </div>
</div>
