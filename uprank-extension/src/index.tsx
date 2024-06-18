import { SignIn, useAuth, useUser } from "@clerk/chrome-extension"
import cssText from "data-text:~style.css"
import { useEffect, useState } from "react"
import { scrape_job_action, scrape_freelancers_action } from "~constants"
import Footer from "./components/footer"
import Header from "./components/header"
import { extractJobId, is_upwork_freelancer, is_upwork_job } from "./utils/url-functions"
import { getWithExpiry, removeItem, setWithExpiry } from "./utils/local-storage-functions"
import { sendToBackground } from "@plasmohq/messaging"
import type { CreateJobClientResponse, Job } from "~types/job"
import { V1Client } from "~client/v1-client"
import type { CreateFreelancerProxyRequest, CreateFreelancerResponse, ScrapeFreelancerResponse } from "~types/freelancer"
export const getStyle = () => {
  const style = document.createElement("style")
  style.textContent = cssText
  return style
}

export default function PopUpEntry() {
  const client = new V1Client()
  const { isSignedIn, user, isLoaded } = useUser()
  const { getToken } = useAuth()
  const [currentURL, setCurrentURL] = useState(null)
  const [isJobValid, setIsJobValid] = useState(false)
  const [jobFreelancerCount, setJobFreelancerCount] = useState(0)
  const [message, setMessage] = useState("")

  useEffect(() => {
    async function validate_job_id(id: string) {
      // Check if the data is already cached
      const cachedData = getWithExpiry(id)
      // if (cachedData) {
      //   setIsJobValid(cachedData.exists)
      //   return
      // }

      const upwork_job_response = await client.getUpworkJob(id, await getToken())
      console.log(upwork_job_response);
        if (upwork_job_response.error_msg == null){
          setIsJobValid(true);
          if (upwork_job_response.upwork_job.edges.upworkfreelancer != null){
            setJobFreelancerCount(upwork_job_response.upwork_job.edges.upworkfreelancer.length);
          }
        } else {
          setIsJobValid(false);
        }

      // Cache the data
      // setWithExpiry(id, data, 60000); //1 minute cache invalidation
    }

    chrome.tabs.query({ active: true, currentWindow: true }, function (tabs) {
      var currentTab = tabs[0]
      if (currentTab) {
        setCurrentURL(currentTab.url)
        if (
          currentTab.url != "ERROR: No tabs found" &&
          currentTab.url.includes("https://www.upwork.com")
        ) {
          const job_id = extractJobId(currentTab.url)
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
  const handleCreateJob = async () => {
    chrome.tabs.query({ active: true, currentWindow: true }, function (tabs) {
      chrome.tabs.sendMessage(
        tabs[0].id,
        { action: scrape_job_action, jobId: extractJobId(currentURL)},
        async function (scrape_response) {
          console.log("received response");
          console.log(scrape_response);
          if (!scrape_response.missingFields) {
            const db_response = await sendToBackground({
              //@ts-ignore
              name: "create-job-proxy",
              body: {
                job: scrape_response.job,
                authentication_token: await getToken()
              }
            }) as CreateJobClientResponse;
            if (db_response.error_msg != null){
              setMessage(`Error ${db_response.error_msg} `)
            } 
            setMessage("Job added successfully!")
            removeItem(extractJobId(currentURL)) //Remove stale data from cache
            setIsJobValid(true)
          }
          else {
            setMessage("Hey! We need some help filling these fields out. Richard you need to add a function to manuall fill out missing fields")
            //todo: Allow user to manually add job.
          }
        }
      )
    })
  }

  const handleAddFreelancers = async () => {
    chrome.tabs.query({ active: true, currentWindow: true }, function (tabs) {
      chrome.tabs.sendMessage(
        tabs[0].id,
        { action: scrape_freelancers_action, jobId: extractJobId(currentURL)},
        async function (scrape_response) {
          if (chrome.runtime.lastError) {
            console.error("Runtime error:", chrome.runtime.lastError);
            setMessage("Error communicating with content script.");
            return;
          }
          if ( scrape_response && !scrape_response.missing_fields && scrape_response.freelancers.length > 0){ //not missing any fields and greater than 0 and not equal to current count (equal to current count => no new freelancers to add)
            console.log("sending freelancer create req from indx")
            const db_response = await sendToBackground({
              //@ts-ignore
              name: "create-upwork-freelancer-proxy",
              body: {
                update: jobFreelancerCount > 0,
                freelancers: scrape_response.freelancers,
                authentication_token: await getToken(),
                job_id: extractJobId(currentURL)
              } as CreateFreelancerProxyRequest
            });
            if (db_response.error_msg != null){
              setMessage("Error persisting data to DB. Please try again.")
            }
            setMessage("Success! You will be notified shortly when your Uprank is ready.")
            removeItem(extractJobId(currentURL)) //Remove stale data from cache
            setJobFreelancerCount(db_response.count);
          } else {
            setMessage("Hey! We need some help filling these fields out. Richard you need to add a function to manually fill out missing fields")
          }

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
        <div className=" space-x-4 justify-start flex">
          <h1>{message}</h1>
          {(is_upwork_freelancer(currentURL) && !isJobValid) &&  <p>Click the "View Job Post" to get started</p>}
          {(is_upwork_job(currentURL) && isJobValid) &&  <p>Click the "Review Proposals" section to get started</p>}
          {isJobValid && jobFreelancerCount!=0 && <a className="bg-white hover:bg-gray-100 text-gray-800 font-semibold py-2 px-4 border border-gray-400 rounded shadow">
            View Uprank
            </a>}
          {isJobValid && is_upwork_freelancer(currentURL) && (
            <button
              onClick={() => handleAddFreelancers()}
              className="bg-white hover:bg-gray-100 text-gray-800 font-semibold py-2 px-4 border border-gray-400 rounded shadow">
                {jobFreelancerCount > 0 ? "Update Freelancers" : "Add Freelancers"}
            </button>
          )}
          {!isJobValid && is_upwork_job(currentURL) && (
            <button
              onClick={() => handleCreateJob()}
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





