import type { PlasmoMessaging } from "@plasmohq/messaging"
import type {  Send_Freelancer_Body, Send_Freelancer_Response} from "~types/freelancer"
//todo: Implement this, currently is just copy pasated over from the send-jobs file.
//then need to create the backend for this endpoint.
const handler: PlasmoMessaging.MessageHandler = async (req, res)   => {
    const body: Send_Freelancer_Body = req.body;
    const response = await fetch(
      `${process.env.PLASMO_PUBLIC_BACKEND_URL}/api/private/job/${body.job_id}/add_freelancers`,
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
    } as Send_Freelancer_Response)
    return;
}
 
export default handler


