import axios, { AxiosResponse } from "axios"
import { useQuery } from "react-query";
import { getAccessToken } from '@auth0/nextjs-auth0';

async function fetchActivities(token: string): Promise<AxiosResponse<any>> {
  try {
    const response = await axios.get('http://localhost:3030/activity', {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });

    return response;
  } catch (error) {
    // consol
    //@ts-ignore

    console.log(error?.message)
     throw error;
  }
}

const useActivitites = async () => {

    console.log("HERE ++==================")
    const x = await getAccessToken().then(t => t.accessToken).catch(err => console.error(err))
    
    console.log("access token",x)


  return useQuery({
    queryKey: ['activities'],
    queryFn: () => fetchActivities(x!),
  })
}

export { useActivitites, fetchActivities }