import { UserButton } from "@clerk/chrome-extension";

export default function Header(){

    return (
        <header className="flex flex-row justify-center items-center">
            <div className="flex flex-1"></div>
            <UserButton
                afterSignOutUrl="/popup.html"
            />
        </header>
    )
}