import { getAccessToken, getSession } from '@auth0/nextjs-auth0';
import axios, { AxiosHeaders } from 'axios';
import type { NextApiRequest, NextApiResponse } from 'next';
import { ResponseData } from './hello';


export default async function handler(
  req: NextApiRequest,
  res: NextApiResponse<ResponseData>
) {

  const { accessToken } = await getAccessToken(req, res);

  const { idToken } = await getSession(req, res);
  const headers: Partial<AxiosHeaders> = {
    Authorization: `Bearer ${idToken}`
  };

  const activities = await axios.get('http://fitter-go:3030/activity', {
    headers
  });



  res.json(activities.data);
}
