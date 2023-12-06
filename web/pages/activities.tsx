
import { QueryClient, dehydrate, useQuery } from "react-query";
import React from "react";
import Image from 'next/image'
import RootLayout, { Layout } from "@/app/layout";
import { Navigation } from "@/app/components/navigation/navigation";
import { useAllActivities } from "@/hooks/activity-hook";
import { GetServerSideProps, GetStaticProps, InferGetServerSidePropsType, InferGetStaticPropsType } from "next";
import axios from "axios";
import { getSession, withPageAuthRequired } from "@auth0/nextjs-auth0";
import { ActivityList } from "@/app/components/activities/activitiesList";



export default function ActivitiesPage({ user }){

  const {data, isFetched} = useAllActivities()

  return (

    <main>

      <h1>Activiteis IndexPage</h1>
     { isFetched && <ActivityList activities={data!}/>}
    </main>
  );
}
export const getServerSideProps = withPageAuthRequired();

