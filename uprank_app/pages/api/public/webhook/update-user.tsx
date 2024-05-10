import prisma from "@/prisma/client";
import { Webhook } from "svix";
import { UserWebhookEvent, WebhookEvent } from "@clerk/nextjs/server";
import { NextApiRequest, NextApiResponse } from "next";
import { buffer } from "micro";
import { UserJSON } from "@clerk/nextjs/server";
export const config = {
    api: {
        bodyParser: false,
    },
};

export default async function handler(
    req: NextApiRequest,
    res: NextApiResponse
) {
    if (req.method !== "POST") {
        return res.status(405);
    }
    // You can find this in the Clerk Dashboard -> Webhooks -> choose the webhook
    const WEBHOOK_SECRET = process.env.WEBHOOK_SECRET;

    if (!WEBHOOK_SECRET) {
        throw new Error(
            "Please add WEBHOOK_SECRET from Clerk Dashboard to .env or .env.local"
        );
    }

    // Get the headers
    const svix_id = req.headers["svix-id"] as string;
    const svix_timestamp = req.headers["svix-timestamp"] as string;
    const svix_signature = req.headers["svix-signature"] as string;

    // If there are no headers, error out
    if (!svix_id || !svix_timestamp || !svix_signature) {
        return res
            .status(400)
            .json({ error: "Error occured -- no svix headers" });
    }

    console.log(
        "headers",
        req.headers,
        svix_id,
        svix_signature,
        svix_timestamp
    );
    // Get the body
    const body = (await buffer(req)).toString();

    // Create a new Svix instance with your secret.
    const wh = new Webhook(WEBHOOK_SECRET);

    let evt: WebhookEvent;

    // Verify the payload with the headers
    try {
        evt = wh.verify(body, {
            "svix-id": svix_id,
            "svix-timestamp": svix_timestamp,
            "svix-signature": svix_signature,
        }) as WebhookEvent;
    } catch (err) {
        console.error("Error verifying webhook:", err);
        return res.status(400).json({ Error: err });
    }

    // Get the ID and type
    const { id } = evt.data;
    const eventType = evt.type;
    const user_event = evt.data as UserJSON;

    console.log(`Webhook with and ID of ${id} and type of ${eventType}`);
    let primary_email_address = "";
    const primary_email_address_id = user_event.primary_email_address_id;

    for (const email of user_event.email_addresses) {
        if (email.id == primary_email_address_id) {
            primary_email_address = email.email_address;
        }
    }
    const user_type = user_event.unsafe_metadata?.type;
    if (user_type == "freelancer") {
        const freelancer = await prisma.freelancer.update({
            where: {
                id: user_event.id,
            },
            data: {
                first_name: user_event.first_name,
                last_name: user_event.last_name,
                email: primary_email_address,
                profile_image_url: user_event.image_url,
            },
        });
    } else if (user_type == "client") {
        const client = await prisma.client.update({
            where: {
                id: user_event.id,
            },
            data: {
                first_name: user_event.first_name,
                last_name: user_event.last_name,
                email: primary_email_address,
                profile_image_url: user_event.image_url,
            },
        });
    } else {
        return res.status(400).json({ error: "User type not recognized" });
    }
    return res.status(200).json({ response: "Success" });
}
