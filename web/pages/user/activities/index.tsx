import React from "react";
import { ACTIVITY_QUERY_KEYS, useAllActivities } from "@/hooks/activity-hook";
import { GetServerSideProps } from "next";
import { getSession, withPageAuthRequired } from "@auth0/nextjs-auth0";
import { ActivityList } from "@/app/components/activities/activitiesList";
import { QueryClient, dehydrate } from "@tanstack/react-query";
import FileUpload from "@/app/components/FileUpload/FileUpload";

export default function ActivitiesPage(props) {

  const { data, isFetched } = useAllActivities()

  return (
    <div>
      <FileUpload />
      {isFetched && <ActivityList activities={data!} />}
    </div>
  );
}
export const getServerSideProps = withPageAuthRequired({
  async getServerSideProps(ctx) {
    const session = getSession(ctx.req, ctx.res);
    const queryClient = new QueryClient();

    await queryClient.prefetchQuery({
      queryKey: [ACTIVITY_QUERY_KEYS.USE_ALL_ACTIVITIES],
      queryFn: useAllActivities
    });

    return {
      props: {
        data: dehydrate(queryClient)
      }
    };
  }
});