import type { PlasmoMessaging } from "@plasmohq/messaging"
import type {  CreateFreelancerProxyRequest, CreateFreelancerResponse, Scraped_Freelancer_Data} from "~types/freelancer"
//todo: Implement this, currently is just copy pasated over from the send-jobs file.
//then need to create the backend for this endpoint.
const handler: PlasmoMessaging.MessageHandler = async (req, res)   => {
  console.log("Create Freelancer Proxy Request")
    const body: CreateFreelancerProxyRequest = req.body;
    const url = body.update ? `${process.env.PLASMO_PUBLIC_BACKEND_URL}/v1/private/jobs/${body.job_id}/freelancers/update` : `${process.env.PLASMO_PUBLIC_BACKEND_URL}/v1/private/jobs/${body.job_id}/freelancers`
    const response = await fetch(
      url,
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

    res.send({
      ok: response.ok,
      count: response_json.result.count
    } as CreateFreelancerResponse)
    return;
}
 
export default handler


