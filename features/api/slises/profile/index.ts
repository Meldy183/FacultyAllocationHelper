import { createApi,  fetchBaseQuery, } from "@reduxjs/toolkit/query/react";
import { API_PATH } from "@/shared/configs/constants/api/paths";
import { GetMemberProcessType, GetAllUsers, CreateMember, GetFiltersType } from "shared/types/api/profile";
import { transformRawFilters } from "@/shared/lib/transformFilter";
import { RawFiltersResponse } from "shared/types/api/filters";
import { buildQuery } from "@/shared/lib/buildQuery";
import { ProfileTag } from "@/shared/configs/constants/dev/cache/tags/profile";
import { instituteList, roleList } from "@/shared/configs/constants/ui";

export const memberSlice = createApi({
  reducerPath: "api/profile",
  baseQuery: fetchBaseQuery({
    baseUrl: `${ API_PATH }/profile/`,
    credentials: "include"
  }),
  tagTypes: [ProfileTag],
  refetchOnReconnect: true,
  endpoints: (builder) => ({
    getFilters: builder.query<GetFiltersType["responseBody"], GetFiltersType["requestQuery"]>({
      query: () => ({
        url: "filters",
        method: "GET",
      }),
      transformResponse: (response: RawFiltersResponse) => transformRawFilters(response)
    }),
    getUser: builder.query<GetMemberProcessType["responseBody"], GetMemberProcessType["requestQuery"]>({
      query: ({ id }) => ({
        url: `getProfile/${ id }`,
        method: "GET",
      }),
      providesTags: [ProfileTag]
    }),
    getMembersByParam: builder.query<GetAllUsers["responseBody"], GetAllUsers["requestQuery"]>({
      query: (query) => ({
        url: `getAllProfiles${ buildQuery(query) }`,
        method: "GET",
      })
    }),
    createUser: builder.mutation<CreateMember["responseBody"], CreateMember["requestBody"]>({
      query: (body) => ({
        url: "addProfile",
        method: "POST",
        body: body
      }),
      invalidatesTags: [ProfileTag],
      async onQueryStarted(newUserBody, { dispatch, queryFulfilled, getState }) {
        const patchResult = dispatch(
          memberSlice.util.updateQueryData(
            'getMembersByParam',
            {},
            (draft) => {
              console.log(newUserBody);
              console.log(draft);
              console.log("<<--->>");
              const newUser = {
                //@ts-ignore
                institute: instituteList.find(item => item.id === newUserBody.institute_id).name,
                //@ts-ignore
                position: roleList.find(item => item.id === newUserBody.position_id).name,
                ...newUserBody
              }
              //@ts-ignore
              draft.profiles = [newUser, ...draft.profiles];
            }
          )
        );

        try {
          await queryFulfilled; // Wait for the actual API call to complete
        } catch {
          patchResult.undo(); // If the API call fails, revert the optimistic update
        }
      },
    })
  })
});

export const {
  useGetUserQuery,
  useLazyGetMembersByParamQuery,
  useCreateUserMutation,
  useGetFiltersQuery
} = memberSlice;