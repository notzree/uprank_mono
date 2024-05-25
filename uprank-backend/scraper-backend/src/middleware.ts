import express, { Request, Response, NextFunction } from 'express';

export const logger = (req: Request, res: Response , next: NextFunction) => {
    var req_time = new Date(Date.now()).toString();
    console.log(req.method,req.hostname, req.path, req_time, req.body);
    next();
  };

