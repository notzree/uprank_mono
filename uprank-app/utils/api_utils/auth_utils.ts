import { NextApiRequest, NextApiResponse } from "next";
import { getAuth } from "@clerk/nextjs/server";
//Returns a user id from a request or throws a 401 error
export function getIdOr401(req: NextApiRequest, res: NextApiResponse) {
    const { userId } = getAuth(req)
    if (!userId && process.env.NODE_ENV === "production") {
        res.status(401).json({ message: "User not authenticated" });
    }
    return userId;
}