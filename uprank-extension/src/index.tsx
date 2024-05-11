import { SignIn, useAuth, useUser } from "@clerk/chrome-extension"
import cssText from "data-text:~style.css"
import { useEffect, useState } from "react"

import { useStorage } from "@plasmohq/storage/hook"

import type { Job } from "~types/job"

import Footer from "./components/footer"
import Header from "./components/header"

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
      const response = await fetch(
        `${process.env.PLASMO_PUBLIC_BACKEND_URL}/api/private/job/${id}`,
        {
          method: "GET",
          headers: {
            "Content-Type": "application/json",
            Authorization: `Bearer ${await getToken()}` // Attach the token in Authorization header
          }
        }
      )
      const data: getJobResult = await response.json()
      console.log(data)
      setIsJobValid(data.exists)
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
        { action: "fetchJob" },
        function (response) {
          console.log("Content script activated:", response.status)
          console.log(response)
        }
      )
    })
  }
  const handleAddFreelancers = async () => {
    chrome.tabs.query({ active: true, currentWindow: true }, function (tabs) {
      chrome.tabs.sendMessage(
        tabs[0].id,
        { action: "fetchFreelancers" },
        function (response) {
          console.log("Content script activated:", response.status)
          console.log(response)
        }
      )
    })
  }

  return (
    <div className="flex flex-col w-96 h-[300px] px-4 py-8 mb-10">
      <Header />
      <div>
        <h1 className=" text-black">Actions</h1>
        <div>
          {isJobValid && is_upwork_freelancer(currentURL) && (
            <button
            onClick={()=>handleAddFreelancers()}
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

function is_upwork_job(url: string) {
  const pattern =
    /^https:\/\/www\.upwork\.com\/[^/]+\/applicants\/[^/]+\/job-details$/
  return pattern.test(url)
}

function is_upwork_freelancer(url: string) {
  const pattern =
    /^https:\/\/www\.upwork\.com\/[^/]+\/applicants\/[^/]+\/applicants$/
  return pattern.test(url)
}

function extractJobId(url: string): string | null {
  const pattern = /\/applicants\/(\d+)\//
  const match = url.match(pattern)
  if (match && match[1]) {
    return match[1]
  }
  return null
}

interface getJobResult {
  exists: boolean
  job: Job | null
}
