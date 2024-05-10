import { ClerkProvider } from "@clerk/chrome-extension"
import { useState } from "react"
import "./style.css"
import PopUpEntry from "~index"

function IndexPopup() {
  const PUBLISHABLE_KEY = process.env.PLASMO_PUBLIC_CLERK_PUBLISHABLE_KEY
  if (!PUBLISHABLE_KEY) {
    throw new Error("Missing Publishable Key")
  }

  return (
    <ClerkProvider publishableKey={PUBLISHABLE_KEY}>
      <PopUpEntry />
    </ClerkProvider>
  )
}

export default IndexPopup
