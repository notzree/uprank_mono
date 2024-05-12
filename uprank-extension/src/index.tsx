import { SignIn, useAuth, useUser } from "@clerk/chrome-extension"
import cssText from "data-text:~style.css"
import { useEffect, useState } from "react"
import { scrape_job_action, scrape_freelancers_action } from "~constants"
import Footer from "./components/footer"
import Header from "./components/header"
import { extractJobId, is_upwork_freelancer, is_upwork_job } from "./utils/url-functions"
import { getWithExpiry, setWithExpiry } from "./utils/local-storage-functions"

import type { Job } from "~types/job"
export const getStyle = () => {
  const style = document.createElement("style")
  style.textContent = cssText
  return style
}

export default function PopUpEntry() {
  const { isSignedIn, user, isLoaded } = useUser()
  const { getToken } = useAuth()
  const [currentURL, setCurrentURL] = useState(null)
  const [isJobValid, setIsJobValid] = useState(false)

  useEffect(() => {
    async function validate_job_id(id: string) {
 
      // Check if the data is already cached
      const cachedData = getWithExpiry(id)
      if (cachedData) {
        setIsJobValid(cachedData.exists)
        return
      }

      const response = await fetch(
        `${process.env.PLASMO_PUBLIC_BACKEND_URL}/api/private/job/${id}`,
        {
          method: "GET",
          headers: {
            "Content-Type": "application/json",
            Authorization: `Bearer ${await getToken()}`
          }
        }
      )
      const data = await response.json()
      console.log(data)
      setIsJobValid(data.exists)

      // Cache the data
      setWithExpiry(id, data, 900000); //15 minute cache invalidation
    }

    chrome.tabs.query({ active: true, currentWindow: true }, function (tabs) {
      var currentTab = tabs[0]
      if (currentTab) {
        setCurrentURL(currentTab.url)
        if (
          currentTab.url != "ERROR: No tabs found" &&
          currentTab.url.includes("https://www.upwork.com")
        ) {
          console.log("upwork tabs")
          const job_id = extractJobId(currentTab.url)
          console.log(job_id)
          validate_job_id(job_id)
        }
      } else {
        return "ERROR: No tabs found"
      }
    })
  }, [])

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

  const handleAddJob = async () => {
    chrome.tabs.query({ active: true, currentWindow: true }, function (tabs) {
      chrome.tabs.sendMessage(
        tabs[0].id,
        { action: scrape_job_action, jobId: extractJobId(currentURL)},
        function (response) {
          console.log("Content script activated:")
          console.log(response)
        }
      )
    })
  }
  const handleAddFreelancers = async () => {
    chrome.tabs.query({ active: true, currentWindow: true }, function (tabs) {
      chrome.tabs.sendMessage(
        tabs[0].id,
        { action: scrape_freelancers_action, jobId: extractJobId(currentURL)},
        function (response) {
          console.log("Content script activated:")
          console.log(response)
        }
      )
    })
  }
  return (
    <div className="flex flex-col w-96 h-[300px] px-4 py-8 mb-10">
      <Header />
      <div>
        <h1 className=" text-black">
          {currentURL && "Uprank Job ID: " + extractJobId(currentURL)}
        </h1>
        <div>
          {(is_upwork_freelancer(currentURL) && !isJobValid) &&  <p>Click the "View Job Post" to get started</p>}
          {isJobValid && is_upwork_freelancer(currentURL) && (
            <button
              onClick={() => handleAddFreelancers()}
              className="bg-white hover:bg-gray-100 text-gray-800 font-semibold py-2 px-4 border border-gray-400 rounded shadow">
              Add freelancers
            </button>
          )}
          {!isJobValid && is_upwork_job(currentURL) && (
            <button
              onClick={() => handleAddJob()}
              className="bg-white hover:bg-gray-100 text-gray-800 font-semibold py-2 px-4 border border-gray-400 rounded shadow">
              Add job
            </button>
          )}
        </div>
      </div>
      <Footer />
    </div>
  )
}




interface getJobResult {
  exists: boolean
  job: Job | null
}
