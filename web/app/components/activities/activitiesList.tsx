import { Activity } from "@/services/activity-services"
import { FC } from "react"
import { DataTable } from "./dataTable"
import { columns } from "./columns"

type Props = {
    activities: Activity[]
}

export const ActivityList: FC<Props> = ({ activities }) => {

    return (
        <div className="container mx-auto py-10">
            <DataTable columns={columns} data={activities} />

        </div>
    )


}