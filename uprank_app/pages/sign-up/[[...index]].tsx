
import { SignUp } from "@clerk/nextjs";



//bug where users who are signed in are able to access the initail page component at first. Should run a check
export default function Page() {

        return (
            <div className="flex flex-col h-screen w-screen justify-center items-center">
                <div className="py-2">
                </div>
                 <SignUp unsafeMetadata={{ "completed_onboarding": false}} afterSignUpUrl="/client/onboarding/onboarding-form"/>
            </div>

        )
    }







function UserIcon(props: any) {
  return (
    <svg
      {...props}
      xmlns="http://www.w3.org/2000/svg"
      width="24"
      height="24"
      viewBox="0 0 24 24"
      fill="none"
      stroke="currentColor"
      strokeWidth="2"
      strokeLinecap="round"
      strokeLinejoin="round"
    >
      <path d="M19 21v-2a4 4 0 0 0-4-4H9a4 4 0 0 0-4 4v2" />
      <circle cx="12" cy="7" r="4" />
    </svg>
  )
}



