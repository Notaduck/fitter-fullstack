"use client"

import { Activity } from "@/services/activity-services"
import { ColumnDef } from "@tanstack/react-table"
import { format } from "util"


export const columns: ColumnDef<Activity>[] = [
  {
    accessorKey: "distance",
    header: "Distance",
  },
  {
    accessorKey: "localTimestamp",
    header: "Date",
    cell: props => <p> ${format(props.getValue<Date>(),'yyyy')}</p>,
  },
]