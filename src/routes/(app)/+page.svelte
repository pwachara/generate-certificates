<script>
    import { goto } from "$app/navigation";
    import { pb } from "$lib/pocketbase";
    import { onMount } from "svelte";

    let noOfParticipants = 0
    
    onMount(async() => {

        if(!pb.authStore.isValid){
            goto("/login", { invalidateAll: true})
        }

        noOfParticipants = await getTotalParticipants();

    })

    async function getTotalParticipants() {
        // use getList with page = 1 and perPage = 1 to minimize payload
        const result = await pb.collection("participants").getList(1, 1);

        console.log("Total participants:", result.totalItems);
        return result.totalItems;
    }



</script>

<!-- Dashboard Content -->
<main class="flex-1 overflow-y-auto p-4 md:p-6 bg-gray-50">
    <!-- Stats Overview -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-4 md:gap-6 mb-6">
        <div class="bg-white rounded-lg shadow p-4 md:p-6">
            <div class="flex items-center">
                <div class="p-3 rounded-lg bg-aqua-light text-teal-dark">
                    <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 8h14M5 8a2 2 0 110-4h14a2 2 0 110 4M5 8v10a2 2 0 002 2h10a2 2 0 002-2V8m-9 4h4" />
                    </svg>
                </div>
                <div class="flex flex-col">
                    <h3 class="text-sm font-medium text-gray-500">Total Participants Trained</h3>
                    <p class="text-xl font-bold text-gray-900 self-center">{ noOfParticipants }</p>
                </div>
            </div>
            <div class="mt-2">
                <p class="text-xs text-green-500 font-medium">
                    <span>+2.5% from last month</span>
                </p>
            </div>
        </div>

        <div class="bg-white rounded-lg shadow p-4 md:p-6">
            <div class="flex items-center">
                <div class="p-3 rounded-lg bg-teal-light text-teal-dark">
                    <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 9V7a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2m2 4h10a2 2 0 002-2v-6a2 2 0 00-2-2H9a2 2 0 00-2 2v6a2 2 0 002 2zm7-5a2 2 0 11-4 0 2 2 0 014 0z" />
                    </svg>
                </div>
                <div class="ml-4">
                    <h3 class="text-sm font-medium text-gray-500">Total Certificates Generated</h3>
                    <p class="text-xl font-bold text-gray-900">242</p>
                </div>
            </div>
            <div class="mt-2">
                <p class="text-xs text-green-500 font-medium">
                    <span>+1 new this week</span>
                </p>
            </div>
        </div>

        <div class="bg-white rounded-lg shadow p-4 md:p-6">
            <div class="flex items-center">
                <div class="p-3 rounded-lg bg-aqua-light text-teal-dark">
                    <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0zm6 3a2 2 0 11-4 0 2 2 0 014 0zM7 10a2 2 0 11-4 0 2 2 0 014 0z" />
                    </svg>
                </div>
                <div class="ml-4">
                    <h3 class="text-sm font-medium text-gray-500">Total Certificates Sent</h3>
                    <p class="text-xl font-bold text-gray-900">156</p>
                </div>
            </div>
            <div class="mt-2">
                <p class="text-xs text-green-500 font-medium">
                    <span>+5 new this week</span>
                </p>
            </div>
        </div>

        <div class="bg-white rounded-lg shadow p-4 md:p-6">
            <div class="flex items-center">
                <div class="p-3 rounded-lg bg-teal-light text-teal-dark">
                    <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
                    </svg>
                </div>
                <div class="ml-4">
                    <h3 class="text-sm font-medium text-gray-500">Certificates not yet Sent</h3>
                    <p class="text-xl font-bold text-gray-900">3</p>
                </div>
            </div>
            <div class="mt-2">
                <p class="text-xs text-green-500 font-medium">
                    <span>+12.3% from last month</span>
                </p>
            </div>
        </div>
    </div>

    <!-- Charts and Tables -->
    <div class="grid grid-cols-1 lg:grid-cols-2 gap-6 mb-6">
        <!-- Recent Sales -->
        <div class="bg-white rounded-lg shadow overflow-hidden">
            <div class="px-4 py-5 sm:px-6 border-b border-gray-200">
                <h3 class="text-lg font-medium leading-6 text-gray-900">Recent Sales</h3>
            </div>
            <div class="px-4 py-5 sm:p-6">
                <div class="overflow-x-auto">
                    <table class="min-w-full divide-y divide-gray-200">
                        <thead>
                            <tr>
                                <th class="px-4 py-3 bg-gray-50 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Order</th>
                                <th class="px-4 py-3 bg-gray-50 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Customer</th>
                                <th class="px-4 py-3 bg-gray-50 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Amount</th>
                                <th class="px-4 py-3 bg-gray-50 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Status</th>
                            </tr>
                        </thead>
                        <tbody class="bg-white divide-y divide-gray-200">
                            <tr>
                                <td class="px-4 py-4 whitespace-nowrap text-sm font-medium text-gray-900">#INV-0035</td>
                                <td class="px-4 py-4 whitespace-nowrap text-sm text-gray-500">John Smith</td>
                                <td class="px-4 py-4 whitespace-nowrap text-sm text-gray-500">$856.00</td>
                                <td class="px-4 py-4 whitespace-nowrap">
                                    <span class="px-2 inline-flex text-xs leading-5 font-semibold rounded-full bg-green-100 text-green-800">Paid</span>
                                </td>
                            </tr>
                            <tr>
                                <td class="px-4 py-4 whitespace-nowrap text-sm font-medium text-gray-900">#INV-0034</td>
                                <td class="px-4 py-4 whitespace-nowrap text-sm text-gray-500">Emily Johnson</td>
                                <td class="px-4 py-4 whitespace-nowrap text-sm text-gray-500">$1,240.00</td>
                                <td class="px-4 py-4 whitespace-nowrap">
                                    <span class="px-2 inline-flex text-xs leading-5 font-semibold rounded-full bg-yellow-100 text-yellow-800">Pending</span>
                                </td>
                            </tr>
                            <tr>
                                <td class="px-4 py-4 whitespace-nowrap text-sm font-medium text-gray-900">#INV-0033</td>
                                <td class="px-4 py-4 whitespace-nowrap text-sm text-gray-500">Robert Williams</td>
                                <td class="px-4 py-4 whitespace-nowrap text-sm text-gray-500">$589.50</td>
                                <td class="px-4 py-4 whitespace-nowrap">
                                    <span class="px-2 inline-flex text-xs leading-5 font-semibold rounded-full bg-green-100 text-green-800">Paid</span>
                                </td>
                            </tr>
                            <tr>
                                <td class="px-4 py-4 whitespace-nowrap text-sm font-medium text-gray-900">#INV-0032</td>
                                <td class="px-4 py-4 whitespace-nowrap text-sm text-gray-500">Sarah Brown</td>
                                <td class="px-4 py-4 whitespace-nowrap text-sm text-gray-500">$2,499.99</td>
                                <td class="px-4 py-4 whitespace-nowrap">
                                    <span class="px-2 inline-flex text-xs leading-5 font-semibold rounded-full bg-red-100 text-red-800">Overdue</span>
                                </td>
                            </tr>
                        </tbody>
                    </table>
                </div>
            </div>
        </div>

        <!-- Low Stock Alert -->
        <div class="bg-white rounded-lg shadow overflow-hidden">
            <div class="px-4 py-5 sm:px-6 border-b border-gray-200">
                <h3 class="text-lg font-medium leading-6 text-gray-900">Low Stock Alert</h3>
            </div>
            <div class="px-4 py-5 sm:p-6">
                <div class="overflow-x-auto">
                    <table class="min-w-full divide-y divide-gray-200">
                        <thead>
                            <tr>
                                <th class="px-4 py-3 bg-gray-50 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Product</th>
                                <th class="px-4 py-3 bg-gray-50 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">SKU</th>
                                <th class="px-4 py-3 bg-gray-50 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Stock</th>
                                <th class="px-4 py-3 bg-gray-50 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Status</th>
                            </tr>
                        </thead>
                        <tbody class="bg-white divide-y divide-gray-200">
                            <tr>
                                <td class="px-4 py-4 whitespace-nowrap text-sm font-medium text-gray-900">Wireless Headphones</td>
                                <td class="px-4 py-4 whitespace-nowrap text-sm text-gray-500">WH-2056</td>
                                <td class="px-4 py-4 whitespace-nowrap text-sm text-gray-500">4</td>
                                <td class="px-4 py-4 whitespace-nowrap">
                                    <span class="px-2 inline-flex text-xs leading-5 font-semibold rounded-full bg-red-100 text-red-800">Low</span>
                                </td>
                            </tr>
                            <tr>
                                <td class="px-4 py-4 whitespace-nowrap text-sm font-medium text-gray-900">Bluetooth Speaker</td>
                                <td class="px-4 py-4 whitespace-nowrap text-sm text-gray-500">BS-7821</td>
                                <td class="px-4 py-4 whitespace-nowrap text-sm text-gray-500">2</td>
                                <td class="px-4 py-4 whitespace-nowrap">
                                    <span class="px-2 inline-flex text-xs leading-5 font-semibold rounded-full bg-red-100 text-red-800">Low</span>
                                </td>
                            </tr>
                            <tr>
                                <td class="px-4 py-4 whitespace-nowrap text-sm font-medium text-gray-900">USB-C Charger</td>
                                <td class="px-4 py-4 whitespace-nowrap text-sm text-gray-500">UC-3098</td>
                                <td class="px-4 py-4 whitespace-nowrap text-sm text-gray-500">7</td>
                                <td class="px-4 py-4 whitespace-nowrap">
                                    <span class="px-2 inline-flex text-xs leading-5 font-semibold rounded-full bg-yellow-100 text-yellow-800">Medium</span>
                                </td>
                            </tr>
                            <tr>
                                <td class="px-4 py-4 whitespace-nowrap text-sm font-medium text-gray-900">Phone Case</td>
                                <td class="px-4 py-4 whitespace-nowrap text-sm text-gray-500">PC-5567</td>
                                <td class="px-4 py-4 whitespace-nowrap text-sm text-gray-500">5</td>
                                <td class="px-4 py-4 whitespace-nowrap">
                                    <span class="px-2 inline-flex text-xs leading-5 font-semibold rounded-full bg-yellow-100 text-yellow-800">Medium</span>
                                </td>
                            </tr>
                        </tbody>
                    </table>
                </div>
            </div>
        </div>
    </div>

    <!-- Recent Activity -->
    <div class="bg-white rounded-lg shadow overflow-hidden mb-6">
        <div class="px-4 py-5 sm:px-6 border-b border-gray-200">
            <h3 class="text-lg font-medium leading-6 text-gray-900">Recent Activity</h3>
        </div>
        <div class="px-4 py-5 sm:p-6">
            <ul class="divide-y divide-gray-200">
                <li class="py-4">
                    <div class="flex space-x-3">
                        <div class="flex-shrink-0">
                            <div class="h-8 w-8 rounded-full bg-teal-100 text-teal-800 flex items-center justify-center">
                                <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m5.618-4.016A11.955 11.955 0 0112 2.944a11.955 11.955 0 01-8.618 3.04A12.02 12.02 0 003 9c0 5.591 3.824 10.29 9 11.622 5.176-1.332 9-6.03 9-11.622 0-1.042-.133-2.052-.382-3.016z" />
                                </svg>
                            </div>
                        </div>
                        <div class="min-w-0 flex-1">
                            <p class="text-sm font-medium text-gray-900">
                                New payment received for order #INV-0035
                            </p>
                            <p class="text-sm text-gray-500">
                                <span>2 hours ago</span>
                            </p>
                        </div>
                    </div>
                </li>
                <li class="py-4">
                    <div class="flex space-x-3">
                        <div class="flex-shrink-0">
                            <div class="h-8 w-8 rounded-full bg-blue-100 text-blue-800 flex items-center justify-center">
                                <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4.354a4 4 0 110 5.292M15 21H3v-1a6 6 0 0112 0v1zm0 0h6v-1a6 6 0 00-9-5.197M13 7a4 4 0 11-8 0 4 4 0 018 0z" />
                                </svg>
                            </div>
                        </div>
                        <div class="min-w-0 flex-1">
                            <p class="text-sm font-medium text-gray-900">
                                New customer registered: Sarah Johnson
                            </p>
                            <p class="text-sm text-gray-500">
                                <span>5 hours ago</span>
                            </p>
                        </div>
                    </div>
                </li>
                <li class="py-4">
                    <div class="flex space-x-3">
                        <div class="flex-shrink-0">
                            <div class="h-8 w-8 rounded-full bg-yellow-100 text-yellow-800 flex items-center justify-center">
                                <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15" />
                                </svg>
                            </div>
                        </div>
                        <div class="min-w-0 flex-1">
                            <p class="text-sm font-medium text-gray-900">
                                Inventory updated: Wireless Headphones stock low
                            </p>
                            <p class="text-sm text-gray-500">
                                <span>Yesterday</span>
                            </p>
                        </div>
                    </div>
                </li>
            </ul>
        </div>
    </div>
</main>
