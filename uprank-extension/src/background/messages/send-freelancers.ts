import type { PlasmoMessaging } from "@plasmohq/messaging"
import type { ScrapedFreelancerData, sendFreelancerBody} from "~types/freelancer"
 
//todo: Implement this, currently is just copy pasated over from the send-jobs file.
//then need to create the backend for this endpoint.
const handler: PlasmoMessaging.MessageHandler = async (req, res) => {
    const body: sendFreelancerBody = req.body;
    const response = await fetch(
      `${process.env.PLASMO_PUBLIC_BACKEND_URL}/api/private/freelancer`,
      {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${body.authentication_token}`
        },
        body: JSON.stringify(body.freelancer)
      }
    )
    res.send({
      status: response.ok
    })
    return;
}
 
export default handler
