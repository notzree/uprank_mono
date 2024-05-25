import express from "express";
import { logger } from "./middleware";
var cors = require("cors");
import helmet from "helmet";
var morgan = require("morgan");

const app = express();
const port = process.env.PORT || 8000;

app.use(helmet());
app.use(
    cors({
        origin: "https://your-other-server.com", // Restrict to your other server
        methods: ["GET", "POST", "PUT", "DELETE"],
        credentials: true,
    })
);
app.use(express.json());
app.use(morgan("combined"));
app.listen(port, () => {
    console.log(`Server is running on http://localhost:${port}`);
});

app.get("/healthz", (req, res) => {
    res.status(200).send("OK");
});

app.get("/:job_id", (req, res) => {
    const jobId = req.params.job_id;

    res.send(`Job ID: ${jobId}`);
});
