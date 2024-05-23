import type { PlasmoMessaging } from "@plasmohq/messaging"
import type { ScrapedJobData, CreateJobProxyBody } from "~types/job"
 
const handler: PlasmoMessaging.MessageHandler = async (req, res) => {
    const body: CreateJobProxyBody = req.body;
    const response = await fetch(
      `${process.env.PLASMO_PUBLIC_BACKEND_URL}/v1/private/jobs`,
      {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${body.authentication_token}`
        },
        body: JSON.stringify(body.job)
      }
    )
    res.send({
      status: response.ok
    })
    return;
}
 
export default handler