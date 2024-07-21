import { SignedIn, SignedOut, SignOutButton } from "@clerk/chrome-extension"

export default function Footer() {
  return (
    <footer>
      <div className="flex flex-col items-center justify-center w-full h-24 ">
        <div className=" text-lg">
          <SignedOut>
            <p>
              Don't have an account? Click
              <a
                href={process.env.PLASMO_PUBLIC_UPRANK_APP_URL}
                className=" text-blue-500">
                {" "}
                here{" "}
              </a>
              to join.
            </p>
          </SignedOut>
        </div>
        <p className="text-sm">Made with ❤️ by <a href="https://www.richard-zhang.ca">notzree</a></p>
      </div>
    </footer>
  )
}
