import express from "express";
import { logger } from "./middleware";
var cors = require("cors");
import helmet from "helmet";
var morgan = require("morgan");
const { SQSClient, ReceiveMessageCommand, DeleteMessageCommand } = require("@aws-sdk/client-sqs");

const PORT = process.env.SCRAPER_PORT || 8000;
const MAIN_BACKEND_URL = process.env.MAIN_BACKEND_URL || "http://localhost:8000";
const SCRAPER_QUEUE_URL = process.env.SCRAPER_QUEUE_URL;


const sqsClient = new SQSClient({ region: "us-east-2" });

async function receiveMessages() {
  const params = {
    QueueUrl: SCRAPER_QUEUE_URL,
    MaxNumberOfMessages: 10,
    WaitTimeSeconds: 20,
  };

  try {
    const data = await sqsClient.send(new ReceiveMessageCommand(params));
    if (data.Messages) {
      for (const message of data.Messages) {
        console.log("Received message:", message.Body);

        // Process the message here

        // Delete the message
        const deleteParams = {
          QueueUrl: SCRAPER_QUEUE_URL,
          ReceiptHandle: message.ReceiptHandle,
        };
        await sqsClient.send(new DeleteMessageCommand(deleteParams));
        console.log("Deleted message:", message.MessageId);
      }
    } else {
      console.log("No messages received");
    }
  } catch (err) {
    console.error("Error receiving messages:", err);
  }
}

// Set up an interval to poll the queue
setInterval(receiveMessages, 30000); // Poll every 30 seconds








const app = express();
app.use(helmet());
app.use(
    cors({
        origin: MAIN_BACKEND_URL, // Restrict to other server
        methods: ["GET", "POST", "PUT", "DELETE"],
        credentials: true,
    })
);
app.use(express.json());
app.use(morgan("combined"));
app.listen(PORT, () => {
    console.log(`Server is running on http://localhost:${PORT}`);
});

app.get("/healthz", (req, res) => {
    res.status(200).send("OK");
});

app.get("/:job_id", (req, res) => {
    const jobId = req.params.job_id;
    res.send(`Job ID: ${jobId}`);
});
