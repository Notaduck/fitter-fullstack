import React, { FC, PropsWithChildren, ReactNode, createContext, useEffect, useMemo, useState } from "react";
import { getSession, withPageAuthRequired } from "@auth0/nextjs-auth0";
import { useRouter } from "next/router";
import { ACTIVITY_QUERY_KEYS, useActivity } from "@/hooks/activity-hook";

import Section from "@/app/components/Section"
import Container from "@/app/components/Container"
import Button from "@/app/components/Button"
import Map from "@/app/components/Map"

import styles from '@/styles/Home.module.scss'
import { QueryClient, dehydrate } from "@tanstack/react-query";
import { AreaChart, Point } from "@/app/components/Charts/LineChart";
import simplify from 'simplify-js'

const DEFAULT_CENTER = [55.676098, 12.568337]
const KM = 1000;

type DataSet = {
    speed: Point[]
    heartRate: Point[]
}


export default function ActivityPage() {


    const router = useRouter()
    const activityId = parseInt(router.query.activityId as string, 10)

    const { data, error, isPending } = useActivity(activityId)


    if (isPending) return 'Loading...'

    if (error) return 'An error has occurred: ' + error.message



    let dataSet: DataSet = data.records.reduce((acc, curr) => {
        let distance = curr.Distance / 100000 // meters to km
        let speed = curr.Speed * 3.600000 / 1000
        if (curr.Distance === 0) {
            return acc
        }

        acc.speed.push({
            id: curr.ID,
            x: distance,
            y: speed
        })

        acc.heartRate.push({
            id: curr.ID,
            x: distance,
            y: curr.HeartRate,
        })

        return acc
    }
        , {
            speed: [],
            heartRate: [],
        } as DataSet)

    return (

        <main className="w-full h-full ">


            <Section>
                <Container>
                    <h1 className={styles.title}>
                        Next.js Leaflet Starter
                    </h1>

                    <Map activity={data} className={styles.homeMap} width="800" height="400" zoom={12} center={DEFAULT_CENTER}>
                        {({ TileLayer }) => (
                            <>
                                <TileLayer
                                    url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
                                    attribution="&copy; <a href=&quot;http://osm.org/copyright&quot;>OpenStreetMap</a> contributors"
                                />
                            </>
                        )}
                    </Map>


                </Container>
            </Section>

            <Row>
                <Column size={"small"} >
                    < AreaChart
                        entries={dataSet.speed}
                        xUnit="km"
                        yUnit="km/h"
                        color="#43ccfe"
                        avgValue={1}
                        maxValue={1}
                    />
                </Column>

            </Row>



            <Row>
                <Column size="small">

                    < AreaChart
                        entries={dataSet.heartRate}
                        xUnit="km"
                        yUnit="bpm"
                        color="#43ccfe"
                        avgValue={1}
                        maxValue={1}
                    />

                </Column>
            </Row>
        </main >
    );
}


export const getServerSideProps = withPageAuthRequired({
    async getServerSideProps(ctx) {
        const activityId = parseInt(ctx.query.activityId as string, 10)
        const session = getSession(ctx.req, ctx.res);
        const queryClient = new QueryClient();

        await queryClient.prefetchQuery({
            queryKey: [ACTIVITY_QUERY_KEYS.USE_ACTIVITY],
            queryFn: () => useActivity(activityId)
        });

        return {
            props: {
                data: dehydrate(queryClient)
            }
        };
    }
});


const Row: FC<PropsWithChildren> = ({ children }) => (
    <div className="grid  grid-flow-col  gap-4 mb-4 w-4/4">
        {children}
    </div>
)

type ColumnProps = {
    size: 'small' | 'default' | 'large',
}

const Column: FC<PropsWithChildren<ColumnProps>> = ({ children, size = 'default' }) => {

    const sizes: Record<ColumnProps['size'], string> = {
        small: 'h-48',
        default: 'h-48',
        large: 'h-96',
    }

    return (
        <div className={`flex items-center justify-center ${sizes[
            size]} mb-4 rounded  `
        }>
            {children}
        </div>
    )
}