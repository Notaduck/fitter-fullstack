

import { QueryClient, dehydrate, useQuery } from "react-query";
import React from "react";
import Image from 'next/image'
import RootLayout, { Layout } from "@/app/layout";
import { Navigation } from "@/app/components/navigation/navigation";
import { useAllActivities } from "@/hooks/activity-hook";
import { GetServerSideProps, GetStaticProps, InferGetServerSidePropsType, InferGetStaticPropsType } from "next";
import axios from "axios";
import { getSession, withPageAuthRequired } from "@auth0/nextjs-auth0";



export default function ActivitiesPage({ user }){

  const { data } = useAllActivities()




  

  return (

<main className="flex min-h-screen flex-col items-center justify-between p-23">
      
      <h1>Activiteis IndexPage</h1>
      {user?.name}


      
    </main>
    

  );
}
export const getServerSideProps = withPageAuthRequired();

