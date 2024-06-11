import type { PlasmoMessaging } from "@plasmohq/messaging"
import type { CreateJobProxyRequest } from "~types/job"
import { V1Client } from "~client/v1-client"
const handler: PlasmoMessaging.MessageHandler = async (req, res) => {
  const client = new V1Client();
    const body: CreateJobProxyRequest = req.body;

    const CreateJobBody = {
      upwork_job_request: body.job
    }
    const response = await client.createJob(CreateJobBody, body.authentication_token)
    res.send(response)
    return;
}
 
export default handler
