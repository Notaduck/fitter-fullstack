import { getAccessToken, getSession } from '@auth0/nextjs-auth0'
import axios, { AxiosHeaders, AxiosRequestConfig, AxiosRequestHeaders, AxiosResponseHeaders } from 'axios'
import type { NextApiRequest, NextApiResponse } from 'next'

type ResponseData = {
  message: string
}

export default async function handler(
  req: NextApiRequest,
  res: NextApiResponse<ResponseData>
) {

  if (req.method === 'GET') {

    const hasQueryParams = Object.keys(req.query).length > 0;

    if(hasQueryParams) {
    }

    try {

      const { accessToken } = await getAccessToken(req, res)

      const { idToken } = await getSession(req, res)
      const headers: Partial<AxiosHeaders> = {
        Authorization: `Bearer ${idToken}`
      }

      const activities = await axios.get('http://fitter-go:3030/activity', {
        headers
      })



      res.json(activities.data)
    } catch (err) {
      console.error(err)
    }



  }
}