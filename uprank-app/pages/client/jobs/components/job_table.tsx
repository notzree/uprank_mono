import * as React from "react";
import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar";
import {
    ColumnDef,
    ColumnFiltersState,
    GlobalFilterColumnDef,
    SortingState,
    VisibilityState,
    flexRender,
    getCoreRowModel,
    getFilteredRowModel,
    getPaginationRowModel,
    getSortedRowModel,
    useReactTable,
} from "@tanstack/react-table";
import {
    AlertDialog,
    AlertDialogAction,
    AlertDialogCancel,
    AlertDialogContent,
    AlertDialogDescription,
    AlertDialogFooter,
    AlertDialogHeader,
    AlertDialogTitle,
    AlertDialogTrigger,
} from "@/components/ui/alert-dialog";
import { ArrowUpDown, ChevronDown, MoreHorizontal } from "lucide-react";

import { Button } from "@/components/ui/button";
import { Checkbox } from "@/components/ui/checkbox";
import {
    DropdownMenu,
    DropdownMenuCheckboxItem,
    DropdownMenuContent,
    DropdownMenuItem,
    DropdownMenuLabel,
    DropdownMenuSeparator,
    DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
import { Input } from "@/components/ui/input";
import {
    Table,
    TableBody,
    TableCell,
    TableHead,
    TableHeader,
    TableRow,
} from "@/components/ui/table";
import type { UpworkFreelancer } from "@/types/freelancer";
import { UpworkJob } from "@/types/job";

export default function JobDataTable({
    freelancers,
    job,
}: {
    freelancers: UpworkFreelancer[];
    job: UpworkJob;
}) {
    const columns: ColumnDef<UpworkFreelancer>[] = [
        {
            id: "select",
            header: ({ table }) => (
                <Checkbox
                    checked={
                        table.getIsAllPageRowsSelected() ||
                        (table.getIsSomePageRowsSelected() && "indeterminate")
                    }
                    onCheckedChange={(value) =>
                        table.toggleAllPageRowsSelected(!!value)
                    }
                    aria-label="Select all"
                />
            ),
            cell: ({ row }) => (
                <Checkbox
                    checked={row.getIsSelected()}
                    onCheckedChange={(value) => row.toggleSelected(!!value)}
                    aria-label="Select row"
                />
            ),
            enableSorting: false,
            enableHiding: false,
        },
        {
            id: "name",
            accessorKey: "name",
            header: ({ column }) => {
                return (
                    <Button
                        variant="ghost"
                        onClick={() =>
                            column.toggleSorting(column.getIsSorted() === "asc")
                        }
                    >
                        Name
                        <ArrowUpDown className="ml-2 h-4 w-4" />
                    </Button>
                );
            },
            cell: ({ row }) => {
                const name = row.original.name;
                const photo_url = row.original.photo_url;
                const freelancer_link = row.original.id;
                return (
                    <div style={{ display: "flex", alignItems: "center" }}>
                        <Avatar>
                            <AvatarImage src={photo_url} />
                            <AvatarFallback>:d</AvatarFallback>
                        </Avatar>
                        <span style={{ marginLeft: "8px" }}>
                            <a
                                href={freelancer_link}
                                target="_blank"
                                className=" underline"
                            >
                                {name}
                            </a>
                        </span>
                    </div>
                );
            },
            enableSorting: true,
            enableHiding: false,
        },
        {
            id: "Location",
            header: "Region Details",
            cell: ({ row }) => (
                <div className="text-right">
                    {row.original.city} | {row.original.country}
                </div>
            ),
            enableSorting: true, // Enable sorting
        },
        {
            id: "description",
            header: "Description",
            cell: ({ row }) => (
                <AlertDialog>
                    <AlertDialogTrigger asChild>
                        <Button variant="outline">View</Button>
                    </AlertDialogTrigger>
                    <AlertDialogContent>
                        <AlertDialogHeader>
                            <AlertDialogTitle>
                                {row.original.name}
                            </AlertDialogTitle>
                            <AlertDialogDescription className=" overflow-y-scroll max-h-96">
                                {row.original.description}
                            </AlertDialogDescription>
                        </AlertDialogHeader>
                        <AlertDialogFooter>
                            <AlertDialogAction>Close</AlertDialogAction>
                        </AlertDialogFooter>
                    </AlertDialogContent>
                </AlertDialog>
            ),
            enableSorting: false,
        },
        {
            id: "skills",
            header: ({ column }) => {
                return (
                    <Button
                        variant="ghost"
                        onClick={() =>
                            column.toggleSorting(column.getIsSorted() === "asc")
                        }
                    >
                        Percent skills matched
                        <ArrowUpDown className="ml-2 h-4 w-4" />
                    </Button>
                );
            },
            cell: ({ row }) => {
                return (
                    <div className="text-right flex justify-end items-center gap-x-4">
                        {row.getValue("skills") + "%" || "No skills indicated"}
                        <AlertDialog>
                            <AlertDialogTrigger asChild>
                                <Button variant="outline">View</Button>
                            </AlertDialogTrigger>
                            <AlertDialogContent>
                                <AlertDialogHeader>
                                    <AlertDialogTitle>
                                        {row.original.name}
                                    </AlertDialogTitle>
                                    <AlertDialogDescription className=" overflow-y-scroll max-h-96">
                                        {
                                            row.original.skills.map(
                                                (skill, index) => {
                                                    return (
                                                        <div key={index}>
                                                            {skill}
                                                        </div>
                                                    );
                                                }
                                            )
                                        }
                                    </AlertDialogDescription>
                                </AlertDialogHeader>
                                <AlertDialogFooter>
                                    <AlertDialogAction>Close</AlertDialogAction>
                                </AlertDialogFooter>
                            </AlertDialogContent>
                        </AlertDialog>
                    </div>
                );
            },
            accessorKey: "skills",
            accessorFn: (row) => {
                let required_skills = job.skills;
                if (!required_skills) {
                    return null;
                }
                let freelancer_skills = row.skills;
                let matched_skills = 0;
                for (let skill of required_skills) {
                    if (freelancer_skills.includes(skill)) {
                        matched_skills += 1;
                    }
                }
                const percent_matched = Math.round(
                    (matched_skills / required_skills.length) * 100
                );
                return percent_matched;
            },
            enableSorting: true,
        },
        {
            id: "proposed_rate",
            header: ({ column }) => {
                return (
                    <Button
                        variant="ghost"
                        onClick={() =>
                            column.toggleSorting(column.getIsSorted() === "asc")
                        }
                    >
                        Proposed Rate
                        <ArrowUpDown className="ml-2 h-4 w-4" />
                    </Button>
                );
            },
            cell: ({ row }) => {
                return (
                    <div className="text-right font-medium">
                        {row.getValue("proposed_rate")}
                    </div>
                );
            },
            accessorFn: (row) => {
                let currency =
                    row.hourly_charge_currency || row.fixed_charge_currency;
                let amount =
                    row.fixed_charge_amount || row.hourly_charge_amount;
                let formatted_string = `${currency} ${amount}`;
                if (row.fixed_charge_amount) {
                    formatted_string += "/hr";
                }
                return formatted_string;
            },
            enableSorting: true,
        },
        {
            id: "recent_hours",
            accessorKey: "recent_hours",
            header: ({ column }) => {
                return (
                    <Button
                        variant="ghost"
                        onClick={() =>
                            column.toggleSorting(column.getIsSorted() === "asc")
                        }
                    >
                        Recent hours
                        <ArrowUpDown className="ml-2 h-4 w-4" />
                    </Button>
                );
            },
            cell: ({ row }) => {
                let hours = "0";
                if (row.original.recent_hours) {
                    hours = row.original.recent_hours.toFixed(0);
                }
                return <div className="text-right">{hours}</div>;
            },
            enableSorting: true,
            enableHiding: true,
        },
        {
            id: "total_hours",
            accessorKey: "total_hours",
            header: ({ column }) => {
                return (
                    <Button
                        variant="ghost"
                        onClick={() =>
                            column.toggleSorting(column.getIsSorted() === "asc")
                        }
                    >
                        Total hours
                        <ArrowUpDown className="ml-2 h-4 w-4" />
                    </Button>
                );
            },
            cell: ({ row }) => {
                let hours = "0";
                if (row.original.total_hours) {
                    hours = row.original.total_hours.toFixed(0);
                }
                return <div className="text-right">{hours}</div>;
            },
            enableSorting: true,
            enableHiding: true,
        },
        {
            id: "average_recent_earnings",
            accessorKey: "average_recent_earnings",
            header: ({ column }) => {
                return (
                    <Button
                        variant="ghost"
                        onClick={() =>
                            column.toggleSorting(column.getIsSorted() === "asc")
                        }
                    >
                        Average Recent Earnings
                        <ArrowUpDown className="ml-2 h-4 w-4" />
                    </Button>
                );
            },
            cell: ({ row }) => {
                let earnings = "0";
                if (row.original.average_recent_earnings) {
                    earnings = row.original.average_recent_earnings.toFixed(0);
                }
                return <div className="text-right">${earnings}</div>;
            },
            accessorFn: (row) => row.average_recent_earnings || 0,
            enableSorting: true,
            enableHiding: true,
        },
        {
            id: "recent_earnings",
            accessorKey: "recent_earnings",
            header: ({ column }) => {
                return (
                    <Button
                        variant="ghost"
                        onClick={() =>
                            column.toggleSorting(column.getIsSorted() === "asc")
                        }
                    >
                        Recent Earnings
                        <ArrowUpDown className="ml-2 h-4 w-4" />
                    </Button>
                );
            },
            cell: ({ row }) => {
                let earnings = "0";
                if (row.original.recent_earnings) {
                    earnings = row.original.recent_earnings.toFixed(0);
                }
                return <div className="text-right">${earnings}</div>;
            },
            accessorFn: (row) => row.recent_earnings || 0,
            enableSorting: true,
            enableHiding: true,
        },
        {
            id: "combined_total_earnings",
            accessorKey: "combined_total_earnings",
            header: ({ column }) => {
                return (
                    <Button
                        variant="ghost"
                        onClick={() =>
                            column.toggleSorting(column.getIsSorted() === "asc")
                        }
                    >
                        Total Earnings
                        <ArrowUpDown className="ml-2 h-4 w-4" />
                    </Button>
                );
            },
            cell: ({ row }) => {
                let earnings = "0";
                if (row.original.combined_total_earnings) {
                    earnings = row.original.combined_total_earnings.toFixed(0);
                }
                return <div className="text-right">${earnings}</div>;
            },
            accessorFn: (row) => row.combined_total_earnings || 0,
            enableSorting: true,
            enableHiding: true,
        },
        {
            id: "specialization_score",
            header: ({ column }) => {
                return (
                    <Button
                        variant="ghost"
                        onClick={() =>
                            column.toggleSorting(column.getIsSorted() === "asc")
                        }
                    >
                        Specialization score
                        <ArrowUpDown className="ml-2 h-4 w-4" />
                    </Button>
                );
            },
            cell: ({ row }) => {
                return (
                    <div className="text-right">
                        {row.getValue("specialization_score")}
                    </div>
                );
            },
            accessorKey: "specialization_score",
            accessorFn: (row) => {
                if (row.edges && row.edges.freelancer_inference_data?.finalized_rating_score) {
                    return row.edges.freelancer_inference_data.finalized_rating_score.toFixed(
                        2
                    ) || 0;
                } else {
                    return 0;
                }
            },
            enableSorting: true,
        },
        {
            id: "budget_adherence_percentage",
            header: ({ column }) => {
                return (
                    <Button
                        variant="ghost"
                        onClick={() =>
                            column.toggleSorting(column.getIsSorted() === "asc")
                        }
                    >
                        Budget adherence percentage
                        <ArrowUpDown className="ml-2 h-4 w-4" />
                    </Button>
                );
            },
            cell: ({ row }) => {
                return (
                    <div className="text-right">
                        {row.getValue("budget_adherence_percentage")+"%" || "No data"}
                    </div>
                );
            },
            accessorKey: "budget_adherence_percentage",
            accessorFn: (row) => {
                if (row.edges && row.edges.freelancer_inference_data?.budget_adherence_percentage) {
                    return row.edges.freelancer_inference_data.budget_adherence_percentage.toFixed(2) || 0;
                } else {
                    return -1;
                }
            },
            enableSorting: true,
        },
        {
            id: "budget_overrun_percentage",
            header: ({ column }) => {
                return (
                    <Button
                        variant="ghost"
                        onClick={() =>
                            column.toggleSorting(column.getIsSorted() === "asc")
                        }
                    >
                        Budget adherence percentage
                        <ArrowUpDown className="ml-2 h-4 w-4" />
                    </Button>
                );
            },
            cell: ({ row }) => {
                const value = row.getValue("budget_overrun_percentage");
                if (value === -1) {
                    return (
                        <div className="text-right text-muted-foreground">
                            No data
                        </div>
                    );
                }
                return (
                    <div className="text-right">
                        {row.getValue("budget_overrun_percentage")+"%"}
                    </div>
                );
            },
            accessorKey: "budget_overrun_percentage",
            accessorFn: (row) => {
                if (row.edges && row.edges.freelancer_inference_data?.budget_overrun_percentage) {
                    return row.edges.freelancer_inference_data.budget_overrun_percentage.toFixed(2) || 0;
                } else {
                    return -1;
                }
            },
            enableSorting: true,
        },
        {
            id: "actions",
            enableHiding: false,
            cell: ({ row }) => {
                const payment = row.original;

                return (
                    <DropdownMenu>
                        <DropdownMenuTrigger asChild>
                            <Button variant="ghost" className="h-8 w-8 p-0">
                                <span className="sr-only">Open menu</span>
                                <MoreHorizontal className="h-4 w-4" />
                            </Button>
                        </DropdownMenuTrigger>
                        <DropdownMenuContent align="end">
                            <DropdownMenuLabel>Actions</DropdownMenuLabel>
                            <DropdownMenuItem
                                onClick={() =>
                                    navigator.clipboard.writeText(payment.id)
                                }
                            >
                                Copy payment ID
                            </DropdownMenuItem>
                            <DropdownMenuSeparator />
                            <DropdownMenuItem>View customer</DropdownMenuItem>
                            <DropdownMenuItem>
                                View payment details
                            </DropdownMenuItem>
                        </DropdownMenuContent>
                    </DropdownMenu>
                );
            },
            enableSorting: false,
        },
    ];
    const [sorting, setSorting] = React.useState<SortingState>([]);
    const [columnFilters, setColumnFilters] =
        React.useState<ColumnFiltersState>([]);
    const [columnVisibility, setColumnVisibility] =
        React.useState<VisibilityState>({});
    const [rowSelection, setRowSelection] = React.useState({});
    const data = freelancers;
    const table = useReactTable({
        data,
        columns,
        onSortingChange: setSorting,
        onColumnFiltersChange: setColumnFilters,
        getCoreRowModel: getCoreRowModel(),
        getPaginationRowModel: getPaginationRowModel(),
        getSortedRowModel: getSortedRowModel(),
        getFilteredRowModel: getFilteredRowModel(),
        onColumnVisibilityChange: setColumnVisibility,
        onRowSelectionChange: setRowSelection,
        state: {
            sorting,
            columnFilters,
            columnVisibility,
            rowSelection,
        },
    });
    if (!job || !freelancers){
        return (
            <div>
                Missing data
            </div>
        )
    }

    

    return (
        <div className="w-full overflow-auto">
            <div className="flex items-center py-4">
                <Input
                    placeholder="Filter freelancers..."
                    value={
                        (table.getColumn("name")?.getFilterValue() as string) ??
                        ""
                    }
                    onChange={(event) =>
                        table
                            .getColumn("name")
                            ?.setFilterValue(event.target.value)
                    }
                    className="max-w-sm"
                />
                <DropdownMenu>
                    <DropdownMenuTrigger asChild>
                        <Button variant="outline" className="ml-auto">
                            Columns <ChevronDown className="ml-2 h-4 w-4" />
                        </Button>
                    </DropdownMenuTrigger>
                    <DropdownMenuContent align="end">
                        {table
                            .getAllColumns()
                            .filter((column) => column.getCanHide())
                            .map((column) => {
                                return (
                                    <DropdownMenuCheckboxItem
                                        key={column.id}
                                        className="capitalize"
                                        checked={column.getIsVisible()}
                                        onCheckedChange={(value) =>
                                            column.toggleVisibility(!!value)
                                        }
                                    >
                                        {column.id}
                                    </DropdownMenuCheckboxItem>
                                );
                            })}
                    </DropdownMenuContent>
                </DropdownMenu>
            </div>
            <div className="rounded-md border">
                <Table>
                    <TableHeader>
                        {table.getHeaderGroups().map((headerGroup) => (
                            <TableRow key={headerGroup.id}>
                                {headerGroup.headers.map((header) => {
                                    return (
                                        <TableHead key={header.id}>
                                            {header.isPlaceholder
                                                ? null
                                                : flexRender(
                                                      header.column.columnDef
                                                          .header,
                                                      header.getContext()
                                                  )}
                                        </TableHead>
                                    );
                                })}
                            </TableRow>
                        ))}
                    </TableHeader>
                    <TableBody>
                        {table.getRowModel().rows?.length ? (
                            table.getRowModel().rows.map((row) => (
                                <TableRow
                                    key={row.id}
                                    data-state={
                                        row.getIsSelected() && "selected"
                                    }
                                >
                                    {row.getVisibleCells().map((cell) => (
                                        <TableCell key={cell.id}>
                                            {flexRender(
                                                cell.column.columnDef.cell,
                                                cell.getContext()
                                            )}
                                        </TableCell>
                                    ))}
                                </TableRow>
                            ))
                        ) : (
                            <TableRow>
                                <TableCell
                                    colSpan={columns.length}
                                    className="h-24 text-center"
                                >
                                    No results.
                                </TableCell>
                            </TableRow>
                        )}
                    </TableBody>
                </Table>
            </div>
            <div className="flex items-center justify-end space-x-2 py-4">
                <div className="flex-1 text-sm text-muted-foreground">
                    {table.getFilteredSelectedRowModel().rows.length} of{" "}
                    {table.getFilteredRowModel().rows.length} row(s) selected.
                </div>
                <div className="space-x-2">
                    <Button
                        variant="outline"
                        size="sm"
                        onClick={() => table.previousPage()}
                        disabled={!table.getCanPreviousPage()}
                    >
                        Previous
                    </Button>
                    <Button
                        variant="outline"
                        size="sm"
                        onClick={() => table.nextPage()}
                        disabled={!table.getCanNextPage()}
                    >
                        Next
                    </Button>
                </div>
            </div>
        </div>
    );
}
