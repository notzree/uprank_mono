import { ClerkProvider, SignedIn, useUser } from "@clerk/chrome-extension"
import cssText from "data-text:~/style.css"
import type { PlasmoCSConfig, PlasmoRender } from "plasmo"
import { useEffect, useState } from "react"
import { createRoot } from "react-dom/client"

export const config: PlasmoCSConfig = {
  matches: ["https://www.upwork.com/freelancers/*"]
}
export const getStyle = () => {
    const style = document.createElement("style")
    style.textContent = cssText
    return style
  }

export function Overlay() {

  const handleClick = async () => {
    // Send a message to content script
    chrome.tabs.query({active: true, currentWindow: true}, (tabs) => {
      chrome.tabs.sendMessage(tabs[0].id, {action: "UPWORK_SCRAPE_FREELANCER"});
    });
  };

  return (
      <div className="block p-6 max-w-sm bg-white rounded-lg border border-gray-200 shadow-md dark:bg-gray-800 dark:border-gray-700 dark:hover:bg-gray-700">
        <h1 className="mb-2 text-3xl font-bold tracking-tight text-gray-900 dark:text-white">
          Uprank
        </h1>
        <div className="flex flex-row space-x-2 px-2">
          <button>Scrape Freelancer</button>
        </div>
        <h2 className="mb-2 text-2xl font-bold tracking-tight text-gray-600 dark:text-white"></h2>
        <p className="font-normal text-gray-700 dark:text-gray-400"></p>
      </div>
  )
}


export default Overlay
