<script>
	import { goto } from "$app/navigation";
    import { pb } from "$lib/pocketbase"
	import { onMount } from "svelte";
    // Svelte 5 code for the component
    let tabSet = $state(0); // 0 for current employees, 1 for ex-employees
    let search = $state("");

    onMount(() => {

        if(!pb.authStore.isValid){

            goto("/login", { invalidateAll:true })
        }
    })
    
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
        <h3 class="text-2xl font-bold p-6 pb-2">Course Listing</h3>

        <!-- New Employee Button -->
        <div class="flex flex-row justify-between px-6 py-4">
            <button class="px-4 py-2 bg-gray-600 text-white font-medium rounded-md hover:bg-gray-700 focus:outline-none focus:ring-2 focus:ring-primary-500 focus:ring-offset-2 transition-colors">
                <a href="/courses/create">
                    + New Course
                </a>
            </button>
        </div>

        <div class="px-6 py-4">
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
                                <th>Employment Number</th>
                                <th>Date Employed</th>
                                <th>Gender</th>
                                <th>Status</th>
                                <th>Action</th>
                            </tr>
                        </thead>
                        <tbody>
                            {#each searchedEmployees as employee (employee.id)}
                                <tr>
                                    <td>{employee.firstName} {employee.middleName} {employee.lastName}</td>
                                    <td>{employee.expand.departmentId.name}</td>
                                    <td>{employee.empNumber}</td>
                                    <td>{employee.dateEmployed ? employee.dateEmployed : "N/A"}</td>
                                    <td>{employee.gender}</td>
                                    <td><span class="px-2 py-1 bg-green-100 text-green-800 text-xs rounded-full">{employee.status}</span></td>
                                    <td>
                                        <button class="w-6 h-6 text-gray-600 hover:text-gray-800" aria-label="button">
                                            <a href={`./organizations/edit/${employee.id}`} aria-label="link">
                                                <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="currentColor"><path d="M12.8995 6.85453L17.1421 11.0972L7.24264 20.9967H3V16.754L12.8995 6.85453ZM14.3137 5.44032L16.435 3.319C16.8256 2.92848 17.4587 2.92848 17.8492 3.319L20.6777 6.14743C21.0682 6.53795 21.0682 7.17112 20.6777 7.56164L18.5563 9.68296L14.3137 5.44032Z"></path></svg>
                                            </a>
                                        </button>
                                    </td>
                                </tr>
                            {/each}
                        </tbody>
                    </table>
                </section>
            </div>
        </div>
    </div>
</section>

    <style>
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
