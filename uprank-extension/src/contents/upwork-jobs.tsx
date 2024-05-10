import { ClerkProvider, SignedIn, useUser } from "@clerk/chrome-extension"
import cssText from "data-text:~/style.css"
import type { PlasmoCSConfig, PlasmoRender } from "plasmo"
import { useEffect, useState } from "react"
import { createRoot } from "react-dom/client"

import type { Freelancer } from "~types/freelancer"

export const config: PlasmoCSConfig = {
  matches: ["https://www.upwork.com/*/*/applicants*"]
}

export function Overlay() {
  const handleClick = async () => {
    // Send a message to content script
    console.log("Scraping job...")
    await preProcessPage()
  }

  return (
    <div>
      <div className="block p-6 max-w-sm bg-white rounded-lg border border-gray-200 shadow-md dark:bg-gray-800 dark:border-gray-700 dark:hover:bg-gray-700">
        <h1 className="mb-2 text-xl font-bold tracking-tight text-gray-900 dark:text-white">
          Uprank
        </h1>
        <div className="flex flex-row space-x-2 px-2">
          <button onClick={() => handleClick()}>Scrape Job</button>
        </div>
        <h2 className="mb-2 text-2xl font-bold tracking-tight text-gray-600 dark:text-white"></h2>
        <p className="font-normal text-gray-700 dark:text-gray-400"></p>
      </div>
    </div>
  )
}

function preProcessPage() {
  //prepares the page for scraping by loading all of the data;
  // Find the button by its data-test attribute and ensure it is treated as an HTMLElement
  const buttons = Array.from(document.querySelectorAll("button"))
  // Find a button that contains the text "Load More"
  const loadMoreButton: HTMLButtonElement = buttons.find(
    (button) => button.textContent.trim() === "Load More"
  )
  if (loadMoreButton) {
    console.log("Loading more freelancers...")
    loadMoreButton.click()
    // Set a timeout before clicking again to allow page/data to load
    setTimeout(preProcessPage, 1000) // Adjust delay as needed for network response times
  } else {
    console.log("All freelancers loaded.")
    scrapeData()
  }
}
async function scrapeData() {
  const name_element = document.querySelectorAll("span.text-base-sm.pr-2x")
  const names = Array.from(name_element, (span) => span.textContent.trim())

  const price_elements = Array.from(document.querySelectorAll("span"))
  const prices = price_elements
    .filter(
      (element) =>
        element.textContent.includes("/hr") && //TODO: MAKE THIS WORK FOR FIXED PRICE
        element.textContent.startsWith("$")
    )
    .map((element) => element.textContent.trim().replace("$", "")) //todo: shorten this bitch
  const earning_elements = Array.from(
    document.querySelectorAll("div.air3-popper-trigger")
  )
  const total_earnings = earning_elements
    .filter(
      (element) =>
        element.textContent.includes("earned") &&
        element.textContent.startsWith("$")
    )
    .map((element) => element.textContent.trim().replace("$", "")) //todo: shorten this bitch

  const location_element = Array.from(
    document.querySelectorAll("div.text-base-sm.text-light-on-inverse")
  )
  const locations = Array.from(location_element, (span) =>
    span.textContent.trim()
  )
  const cv_elements = Array.from(
    document.querySelectorAll("div.air3-line-clamp.is-clamped")
  )
  const cv = cv_elements.map((element) => element.textContent.trim())

  const clickabledivs = document.querySelectorAll("div[data-ev-contractor_uid]")
  clickabledivs.forEach((div: HTMLElement, index) => {
    setTimeout(() => {
      div.click() // Open the popup
      setTimeout(() => {
        const links = document.querySelectorAll("a.up-n-link.d-block.my-3x")
        links.forEach((link) => {
          const href = link.getAttribute("href")
          if (href && href.startsWith("https://www.upwork.com/fl/")) {
            console.log("Found link:", href)
            // Perform further actions here
          }
        })
        const backButton = document.querySelector(
          "button.m-0.p-0.air3-btn.air3-btn-link.d-none.d-md-block"
        ) as HTMLButtonElement
        backButton.click() // Close the popup
      }, 1750) // Wait for the popup to load and links to be available // This needs to be tweaked based on network speed which is extremely difficult to do.
      //TODO: THE MOST IMPORTANT ONE, IMPLEMENT MUTATION OBSERVER TO WAIT FOR THE POPUP TO LOAD INSTEAD OF HAVING JANKY TIMINGS
    }, index * 2000) // Increase delay for each div to ensure previous actions complete
  })

  let freelancers: Freelancer[] = names.map((name, index) => ({
    name: name,
    price: prices[index],
    location: locations[index],
    totalEarnings: total_earnings[index],
    cv: cv[index],
    jobs: []
  }))
  console.log(freelancers)
}

export const getStyle = () => {
  const style = document.createElement("style")
  style.textContent = cssText
  return style
}

export default Overlay
