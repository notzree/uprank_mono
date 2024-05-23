import type { PlasmoMessaging } from "@plasmohq/messaging"
import type {  CreateFreelancerProxyRequest, ScrapeFreelancerResponse, Scraped_Freelancer_Data} from "~types/freelancer"
//todo: Implement this, currently is just copy pasated over from the send-jobs file.
//then need to create the backend for this endpoint.
const handler: PlasmoMessaging.MessageHandler = async (req, res)   => {
  console.log("Create Freelancer Proxy Request")
    const body: CreateFreelancerProxyRequest = req.body;
    const response = await fetch(
      `${process.env.PLASMO_PUBLIC_BACKEND_URL}/v1/private/jobs/${body.job_id}/freelancers`,
      {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${body.authentication_token}`
        },
        body: JSON.stringify(body.freelancers as Scraped_Freelancer_Data[])
        
      }
    )
    const response_json = await response.json()
    console.log(response_json)
    res.send({
      ok: response.ok,
      count: response_json.result.count
    } as ScrapeFreelancerResponse)
    return;
}
 
export default handler


