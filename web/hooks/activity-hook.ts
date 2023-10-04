import activityServices from "@/services/activity-services";
import { getAccessToken, getSession } from "@auth0/nextjs-auth0";
import { useUser } from "@auth0/nextjs-auth0/client";
import { useQuery, useQueryClient, useMutation } from "react-query";

const useAllActivities = async () => {

//   const { user } = useUser();
//   console.log('USER',user)

    return useQuery(["activities"],  activityServices.getAllActivities);
};

// const usePostById = () => {
//   return useQuery(["posts"], exampleService.getByPostId());
// };

// const useCreatePost = () => {
//   const queryClient = useQueryClient();
//   return useMutation(
//     () => {
//       return exampleService.addPost();
//     },
//     {
//       onSuccess: () => {
//         queryClient.invalidateQueries("posts");
//       },
//     }
//   );
// };

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
   useAllActivities 
  
};