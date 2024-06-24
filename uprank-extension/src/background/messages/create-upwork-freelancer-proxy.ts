import type { PlasmoMessaging } from "@plasmohq/messaging"
import { V1Client } from "~client/v1-client";
import type {  CreateFreelancerProxyRequest, CreateFreelancerResponse, Scraped_Freelancer_Data} from "~types/freelancer"
//todo: Implement this, currently is just copy pasated over from the send-jobs file.
//then need to create the backend for this endpoint.
const handler: PlasmoMessaging.MessageHandler = async (req, res)   => {
  console.log("sending freelancer create req")
  const client = new V1Client();
  const body: CreateFreelancerProxyRequest = req.body;
  var response = null;
    if (body.update){
      response = await client.updateUpworkFreelancers(body.job_id, body, body.authentication_token)
    } else {
      response = await client.createUpworkFreelancers(body.job_id, body, body.authentication_token)
    }
    res.send(response)
    return;
}

export default handler


