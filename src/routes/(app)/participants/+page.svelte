<script>
    // Svelte 5 code for the component
    let tabSet = 0; // 0 for current employees, 1 for ex-employees
    let search = '';
    
    // Sample data - in a real app this would come from a store or API
    const employees = $state([
        {
            id: 1,
            firstName: 'John',
            middleName: 'Michael',
            lastName: 'Smith',
            expand: { departmentId: { name: 'Engineering' } },
            empNumber: 'EMP-10245',
            dateEmployed: '2020-01-15',
            gender: 'Male',
            status: 'Active'
        },
        {
            id: 2,
            firstName: 'Sarah',
            middleName: '',
            lastName: 'Johnson',
            expand: { departmentId: { name: 'Marketing' } },
            empNumber: 'EMP-10246',
            dateEmployed: '2021-03-22',
            gender: 'Female',
            status: 'Active'
        },
        {
            id: 3,
            firstName: 'Robert',
            middleName: '',
            lastName: 'Williams',
            expand: { departmentId: { name: 'Sales' } },
            empNumber: 'EMP-10247',
            dateEmployed: '2019-08-05',
            gender: 'Male',
            status: 'Active'
        }
    ])
    
    const exEmployees = $state([
        {
            id: 4,
            firstName: 'Jennifer',
            middleName: '',
            lastName: 'Lopez',
            expand: { departmentId: { name: 'HR' } },
            empNumber: 'EMP-10123',
            dateEmployed: '2018-02-10',
            dateLeft: '2022-06-15',
            gender: 'Female',
            status: 'Inactive'
        },
        {
            id: 5,
            firstName: 'David',
            middleName: '',
            lastName: 'Brown',
            expand: { departmentId: { name: 'Finance' } },
            empNumber: 'EMP-10124',
            dateEmployed: '2017-05-18',
            dateLeft: '2021-11-30',
            gender: 'Male',
            status: 'Inactive'
        }
    ])
    
    // Computed values for searched employees

    let searchedEmployees = $derived(employees.filter(emp => 
            `${emp.firstName} ${emp.middleName} ${emp.lastName}`
                .toLowerCase()
                .includes(search.toLowerCase())
        ))
    
    let searchedExEmployees = $derived(exEmployees.filter(emp => 
            `${emp.firstName} ${emp.middleName} ${emp.lastName}`
                .toLowerCase()
                .includes(search.toLowerCase())
        ))
    
    // Function to update active tab
    function setTabSet(value) {
        tabSet = value;
    }

</script>

<section class="bg-gray-50 min-h-screen">
    <div class="container mx-auto px-4 py-8">
        <!-- Header -->
        <h3 class="text-2xl font-bold p-6 pb-2">Employee Listing</h3>

        <!-- New Employee Button -->
        <div class="flex flex-row justify-between px-6 py-4">
            <button class="px-4 py-2 bg-primary-600 text-white font-medium rounded-md hover:bg-primary-700 focus:outline-none focus:ring-2 focus:ring-primary-500 focus:ring-offset-2 transition-colors">
                <a href="/employees/create">
                    <i class="fas fa-plus mr-2"></i> New Employee
                </a>
            </button>
        </div>

        <!-- Tabs -->
        <div class="px-6 pt-4">
            <div class="border-b border-gray-200">
                <div class="flex space-x-8">
                    <button class="py-2 px-1 font-semibold text-sm tab-active" onclick={() => setTabSet(0)}>
                        Current Employees
                    </button>
                    <button class="py-2 px-1 font-semibold text-sm text-gray-500 hover:text-gray-700" onclick={() => setTabSet(1)}>
                        Ex-Employees
                    </button>
                </div>
            </div>
        </div>

        <!-- Tab Panels -->
        <div class="px-6 py-4">
            {#if tabSet === 0}
                <!-- Current Employees Panel -->
                <div>
                    <div class="mb-6">
                        <div class="relative w-full md:w-1/3">
                            <div class="absolute inset-y-0 left-0 flex items-center pl-3 pointer-events-none">
                                <i class="fas fa-search text-gray-400"></i>
                            </div>
                            <input 
                                class="pl-10 pr-4 py-2 border border-gray-300 rounded-md w-full focus:outline-none focus:ring-2 focus:ring-primary-500 focus:border-transparent" 
                                title="Search by Name" 
                                type="search" 
                                placeholder="Search by name..." 
                                bind:value={search}
                            />
                        </div>
                    </div>

                    <section class="py-4 overflow-x-auto">
                        <table class="employee-table">
                            <thead>
                                <tr>
                                    <th>Employee</th>
                                    <th>Department</th>
                                    <th class="text-center">Employment Number</th>
                                    <th class="text-center">Date Employed</th>
                                    <th class="text-center">Gender</th>
                                    <th class="text-center">Status</th>
                                    <th class="text-center">Action</th>
                                </tr>
                            </thead>
                            <tbody>
                                {#each searchedEmployees as employee (employee.id)}
                                    <tr>
                                        <td class="font-medium">{employee.firstName} {employee.middleName} {employee.lastName}</td>
                                        <td>{employee.expand.departmentId.name}</td>
                                        <td class="text-center">{employee.empNumber}</td>
                                        <td class="text-center">{employee.dateEmployed ? employee.dateEmployed : "N/A"}</td>
                                        <td class="text-center">{employee.gender}</td>
                                        <td class="text-center"><span class="px-2 py-1 bg-green-100 text-green-800 text-xs rounded-full">{employee.status}</span></td>
                                        <td class="text-center">
                                            <button class="w-6 h-6 text-blue-500 hover:text-blue-700" aria-label="button">
                                                <a href={`./employees/edit/${employee.id}`} aria-label="link">
                                                    <i class="fas fa-edit"></i>
                                                </a>
                                            </button>
                                        </td>
                                    </tr>
                                {/each}
                            </tbody>
                        </table>
                    </section>
                </div>
            {:else if tabSet === 1}
                <!-- Ex-Employees Panel -->
                <div>
                    <div class="mb-6">
                        <div class="relative w-full md:w-1/3">
                            <div class="absolute inset-y-0 left-0 flex items-center pl-3 pointer-events-none">
                                <i class="fas fa-search text-gray-400"></i>
                            </div>
                            <input 
                                class="pl-10 pr-4 py-2 border border-gray-300 rounded-md w-full focus:outline-none focus:ring-2 focus:ring-primary-500 focus:border-transparent" 
                                title="Search by Name" 
                                type="search" 
                                placeholder="Search by name..." 
                                bind:value={search}
                            />
                        </div>
                    </div>

                    <section class="py-4 overflow-x-auto">
                        <table class="employee-table">
                            <thead>
                                <tr>
                                    <th>Employee</th>
                                    <th>Department</th>
                                    <th class="text-center">Employment Number</th>
                                    <th class="text-center">Date Employed</th>
                                    <th class="text-center">Gender</th>
                                    <th class="text-center">Status</th>
                                    <th class="text-center">Date Left Employment</th>
                                </tr>
                            </thead>
                            <tbody>
                                {#each searchedExEmployees as employee (employee.id)}
                                    <tr>
                                        <td class="font-medium">{employee.firstName} {employee.middleName} {employee.lastName}</td>
                                        <td>{employee.expand.departmentId.name}</td>
                                        <td class="text-center">{employee.empNumber}</td>
                                        <td class="text-center">{employee.dateEmployed ? employee.dateEmployed : "N/A"}</td>
                                        <td class="text-center">{employee.gender}</td>
                                        <td class="text-center"><span class="px-2 py-1 bg-red-100 text-red-800 text-xs rounded-full">{employee.status}</span></td>
                                        <td class="text-center">{employee.dateLeft ? employee.dateLeft : "N/A"}</td>
                                    </tr>
                                {/each}
                            </tbody>
                        </table>
                    </section>
                </div>
            {/if}
        </div>
    </div>
</section>

    <style>
        .tab-active {
            border-bottom: 3px solid #3b82f6;
            color: #2563eb;
            font-weight: 600;
        }
        .employee-table {
            width: 100%;
            border-collapse: collapse;
        }
        .employee-table th {
            background-color: #f8fafc;
            padding: 0.75rem;
            text-align: left;
            font-weight: 600;
            color: #374151;
            border-bottom: 1px solid #e5e7eb;
        }
        .employee-table td {
            padding: 0.75rem;
            border-bottom: 1px solid #e5e7eb;
            white-space: nowrap;
        }
        .employee-table tr:hover {
            background-color: #f9fafb;
        }
    </style>
