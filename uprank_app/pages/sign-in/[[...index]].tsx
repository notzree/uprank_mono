import { SignIn } from "@clerk/nextjs";
import Link from "next/link";
import { IoMdArrowBack } from "react-icons/io";
export default function Page() {
  return (
    <div className="flex flex-col space-y-2 h-screen w-screen justify-center items-center">
        <Link className="h-14 justify-start"  href="/">
            <div className="flex items-center space-x-2">
                <IoMdArrowBack />
                <span className="font-semibold">Back to home</span>
            </div>
        </Link>
         <SignIn redirectUrl="/client/dashboard" />
    </div>
)
}
