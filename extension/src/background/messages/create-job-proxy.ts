import type { PlasmoMessaging } from "@plasmohq/messaging"
import type { CreateJobProxyRequest, PlatformJobRequests} from "~types/job"
import { V1Client } from "~client/v1-client"
const handler: PlasmoMessaging.MessageHandler = async (req, res) => {
  const client = new V1Client();
    const body: CreateJobProxyRequest = req.body;

    const platform_job_requests: PlatformJobRequests = {
      upwork_request: body.job
    }
    const CreateJobBody = {
      origin: "upwork",
      platform_job_requests: platform_job_requests
    }


    const response = await client.createJob(CreateJobBody, body.authentication_token)
    res.send(response)
    return;
}

export default handler
