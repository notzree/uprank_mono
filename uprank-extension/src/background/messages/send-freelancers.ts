import type { PlasmoMessaging } from "@plasmohq/messaging"
import type { ScrapedFreelancerData, SendFreelancerBody, SendFreelancerResponse} from "~types/freelancer"
 
//todo: Implement this, currently is just copy pasated over from the send-jobs file.
//then need to create the backend for this endpoint.
const handler: PlasmoMessaging.MessageHandler = async (req, res)   => {
    const body: SendFreelancerBody = req.body;
    const response = await fetch(
      `${process.env.PLASMO_PUBLIC_BACKEND_URL}/api/private/freelancer/${body.jobId}/addFreelancers`,
      {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${body.authentication_token}`
        },
        body: JSON.stringify(body.freelancers)
      }
    )
    const response_json = await response.json()
    res.send({
      ok: response.ok,
      count: response_json?.count,
    } as SendFreelancerResponse)
    return;
}
 
export default handler


