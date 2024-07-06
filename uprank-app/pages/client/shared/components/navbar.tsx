import { PackageIcon, Package2Icon, SearchIcon } from "@/public/Icons";
import { UserButton } from "@clerk/nextjs";
import Link from "next/link";
import { Input } from "@/components/ui/input";
import { useRouter } from "next/router";
export default function Navbar() {
    const currentPath = useRouter().pathname;
    const is_marketplace = (currentPath === "/client/dashboard");

    return (
        <header className="flex h-14 lg:h-[60px] items-center justify-end gap-4 border-b bg-gray-100/40 px-6 dark:bg-gray-800/40">
            <p>
                Uprank
            </p>
            <div className="flex flex-1"></div>
            <Link href="/client/dashboard">Dashboard</Link>
            <Link href="/client/dashboard/subscription">Billing</Link>
            <UserButton afterSignOutUrl="/" />
        </header>
    );
}
