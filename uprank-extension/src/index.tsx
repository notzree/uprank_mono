import {
  SignedIn,
  SignedOut,
  SignIn,
  SignUp,
  useClerk,
  UserButton,
  useUser
} from "@clerk/chrome-extension"
import cssText from "data-text:~style.css"

import Footer from "./components/footer"
import Header from "./components/header"
import { useEffect } from "react"
import { useStorage } from "@plasmohq/storage/hook"


export const getStyle = () => {
  const style = document.createElement("style")
  style.textContent = cssText
  return style
}

export default function PopUpEntry() {
  const { isSignedIn, user, isLoaded } = useUser()
  const [plasmo_isSignedIn, setPlasmo_isSignedIn] = useStorage("signedIn", isSignedIn)
  useEffect(() => {
    setPlasmo_isSignedIn(isSignedIn)
  },[isSignedIn, isLoaded])


  console.log("isSignedIn", user)
  if (isLoaded && !isSignedIn) {
    return (
      <div className="flex flex-col w-96 h-[150px] px-4 py-8 mb-10">
        <div className="mb-1">
          <SignIn
          forceRedirectUrl="/popup.html"
            appearance={{
              elements: {
                footerAction: { display: "none" }
              }
            }}
          />
        </div>
        <div className="flex justify-center flex-col">
          <Footer />
        </div>
      </div>
    )
  }
  
  return (
    <div className="flex flex-col w-96 h-[300px] px-4 py-8 mb-10">
      <Header />
      <div>
      <h1 className=" text-black">Actions</h1>


      </div>
      <Footer />
    </div>
  )
}
