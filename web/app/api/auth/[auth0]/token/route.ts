// app/my-api/route.js
import { NextResponse } from 'next/server';
import { getAccessToken } from '@auth0/nextjs-auth0/edge'; // Note the /edge import

export async function GET() {
  const { accessToken } = await getAccessToken();
  console.log(accessToken)
  return NextResponse.json({ foo: 'bar' });
}

export const runtime = 'edge';