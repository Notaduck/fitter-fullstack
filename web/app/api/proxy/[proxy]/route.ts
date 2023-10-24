
import { getAccessToken } from "@auth0/nextjs-auth0";
import { NextRequest, NextResponse } from "next/server";

export async function GET(req: NextRequest) {

    return { Hello : 1}
    // const res = new NextResponse();
    // const { accessToken } = await getAccessToken(req, res);
    // console.log('access token', accessToken)
    // const path = req.nextUrl.pathname.replace("/api/", "");
    // console.log('access token', accessToken)
    // const EXTERNAL_API_URL = "http://fitter-go:3030";
    // console.log('access token', accessToken)
    
    // // return 1337 
    // return await fetch(`https:/${EXTERNAL_API_URL}/${path}${req.nextUrl.search}`, {
    //     headers: {
    //         Authorization: `Bearer ${accessToken}`,
    //     },
    // });
}
