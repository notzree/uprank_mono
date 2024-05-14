export function is_upwork_job(url: string) {
  const pattern =
    /^https:\/\/www\.upwork\.com\/[^/]+\/applicants\/[^/]+\/job-details.*$/
  return pattern.test(url)
}

export function is_upwork_freelancer(url: string) {
  const pattern =
  /^https:\/\/www\.upwork\.com\/[^/]+\/applicants\/[^/]+\/applicants.*$/
  return pattern.test(url)
}

export function extractJobId(url: string): string | null {
  const pattern = /\/applicants\/(\d+)\//
  const match = url.match(pattern)
  if (match != null && match[1]) {
    return match[1]
  }
  return null
}

//wraps the chrome.tabs.query function in a promise to make it easier to use
//Returns the url of the current tab
export function getCurrentTabUrl(): Promise<string> {
  return new Promise((resolve, reject) => {
      chrome.tabs.query({ active: true, currentWindow: true }, function(tabs) {
          if (tabs.length === 0) {
              reject(new Error("No active tab found"));
              return;
          }
          console.log(tabs[0].url);
          resolve(tabs[0].url);
      });
  });
}
