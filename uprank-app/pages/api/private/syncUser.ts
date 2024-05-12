import { NextApiRequest, NextApiResponse } from "next";
import prisma from "@/prisma/client";
import { User } from "@prisma/client";


export type SyncUserBody = {
    user_data: User;
    metadata: {
        completed_onboarding: boolean;
    };
}


export default async function handler (req:NextApiRequest, res: NextApiResponse) {
    if (req.method !=="PUT"){
        res.setHeader('Allow', ['PUT']);
        res.status(405).end(`Method ${req.method} Not Allowed`);
    }
    const body: SyncUserBody = req.body.request_body;
    const user_data = body.user_data;
    try{
        const user = await createUser(user_data);
        res.status(200).json(user);
    }
    catch(e: any){
        res.status(500).json({error: e.message})
    }

}

//todo: edit the frontend and this request body to agree to send smth different. Have to change FE form to include 
//monthly rate calculator / transformer ting? have to think about globally set prices vs self-set prices
//must consider scenarios where people just inflate their own worth...

const createUser = async (user_data: User) => {
    try{
        return await prisma.user.create({
            data: user_data as User ,
        });
    }
    catch (e: any){
        throw new Error(e);
    }
};


