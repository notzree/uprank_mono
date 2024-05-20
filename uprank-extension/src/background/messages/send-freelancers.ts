import type { PlasmoMessaging } from "@plasmohq/messaging"
import type {  Send_Freelancer_Body, Send_Freelancer_Response, Add_Freelancers_Request, Add_Freelancers_Response} from "~types/freelancer"
//todo: Implement this, currently is just copy pasated over from the send-jobs file.
//then need to create the backend for this endpoint.
const handler: PlasmoMessaging.MessageHandler = async (req, res)   => {
    const body: Send_Freelancer_Body = req.body;
    const response = await fetch(
      `${process.env.PLASMO_PUBLIC_BACKEND_URL}/api/private/jobs/${body.job_id}/add_freelancers`,
      {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${body.authentication_token}`
        },
        body: JSON.stringify({
          "freelancer_data": body.freelancers
        } as Add_Freelancers_Request)
        
      }
    )
    const response_json: Add_Freelancers_Response = await response.json()
    console.log(response_json)
    res.send({
      ok: response.ok,
      count: response_json.result.count
    } as Send_Freelancer_Response)
    return;
}
 
export default handler


