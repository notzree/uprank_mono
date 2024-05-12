import { NextApiRequest, NextApiResponse } from "next";

const enableCors = (fn: (arg0: NextApiRequest, arg1: NextApiResponse) => any) => async (req: NextApiRequest, res: NextApiResponse) => {
    res.setHeader('Access-Control-Allow-Credentials', 'true');
    res.setHeader('Access-Control-Allow-Origin', '*'); // replace "*" with specific origins in production
    res.setHeader('Access-Control-Allow-Methods', 'GET, POST, PUT, DELETE, PATCH, OPTIONS');
    res.setHeader(
        'Access-Control-Allow-Headers',
        'Authorization, X-CSRF-Token, X-Requested-With, Accept, Accept-Version, Content-Length, Content-MD5, Content-Type, Date, X-Api-Version'
    );

    // Handle the preflight request
    if (req.method === 'OPTIONS') {
        res.status(204).end();
        return;
    }
  
    return await fn(req, res);
}

export default enableCors;
