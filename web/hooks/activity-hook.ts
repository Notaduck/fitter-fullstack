import activityServices from "@/services/activity-services";
import { getAccessToken, getSession } from "@auth0/nextjs-auth0";
import { useUser } from "@auth0/nextjs-auth0/client";
import axios from "axios";
import { useQuery, useQueryClient, useMutation } from "@tanstack/react-query";


export enum ACTIVITY_QUERY_KEYS {
    USE_ALL_ACTIVITIES = 'all_activities',
    USE_ACTIVITY = 'activity'

}


const useAllActivities = () => {
    return useQuery({
        queryKey: [ACTIVITY_QUERY_KEYS.USE_ALL_ACTIVITIES],
        queryFn: activityServices.getAllActivities
    });
};

const useActivity = (id: number) => {
    return useQuery({
        queryKey: [ACTIVITY_QUERY_KEYS.USE_ACTIVITY, id],
        queryFn: () => activityServices.getActivityById(id)
    });
};


const useCreateActivities = () => {
    const queryClient = useQueryClient();
    return useMutation(
        (data) => {
            return activityServices.addActivity(data);
        },
        {
            onSuccess: () => {
                queryClient.invalidateQueries(
                    {
                        queryKey: [ACTIVITY_QUERY_KEYS.USE_ALL_ACTIVITIES]
                    }
                );
            },
        }
    );
};

// const useUpdatePost = () => {
//   const queryClient = useQueryClient();
//   return useMutation(
//     () => {
//       return exampleService.updatePost();
//     },
//     {
//       onSuccess: () => {
//         queryClient.invalidateQueries("posts");
//       },
//     }
//   );
// };

// const useDeletePost = () => {
//   const queryClient = useQueryClient();
//   return useMutation(
//     () => {
//       return exampleService.deletePost();
//     },
//     {
//       onSuccess: () => {
//         queryClient.invalidateQueries("posts");
//       },
//     }
//   );
// };

export {
    useAllActivities,
    useActivity,
    useCreateActivities


};
