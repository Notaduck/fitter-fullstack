import { getAccessToken, getSession } from '@auth0/nextjs-auth0'
import axios, { AxiosRequestConfig, AxiosRequestHeaders, AxiosResponseHeaders } from 'axios'
import type { NextApiRequest, NextApiResponse } from 'next'
 
type ResponseData = {
  message: string
}
 
export default async function handler(
  req: NextApiRequest,
  res: NextApiResponse<ResponseData>
) {

    const { accessToken} = await getAccessToken(req,res)

    const { idToken } = await getSession(req,res)
    const headers : Partial<AxiosRequestHeaders> = {
     Authorization: `Bearer ${idToken}`  
    }

    const activities = await axios.get('http://fitter-go:3030/activity', {
      headers
    })

    console.log('access token ======>',activities)
    // console.log(s?.accessToken)
    // console.log(s?.idToken)

    res.status(200).json(activities)
}