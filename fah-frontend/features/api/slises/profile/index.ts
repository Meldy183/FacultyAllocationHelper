import { createApi,  fetchBaseQuery, } from "@reduxjs/toolkit/query/react";
import { API_PATH } from "@/shared/configs/constants/api/paths";
import { buildQuery } from "@/shared/lib/buildQuery";
import { ProfileTag } from "@/shared/configs/constants/dev/cache/tags/profile";
import { CreateMember, GetAllUsers, GetMemberProcessType } from "@/shared/types/api/profile";
import { GetSimpleUserDataInterface } from "@/shared/types/ui/faculties";

export const memberSlice = createApi({
  reducerPath: "api/profile",
  baseQuery: fetchBaseQuery({
    baseUrl: `${ API_PATH }/profile/`,
    credentials: "include"
  }),
  tagTypes: [ProfileTag],
  refetchOnReconnect: true,
  endpoints: (builder) => ({
    getUser: builder.query<GetMemberProcessType["responseBody"], GetMemberProcessType["requestQuery"]>({
      query: ({ id }) => ({
        url: `getProfile/${ id }`,
        method: "GET",
      }),
      providesTags: (result, err, arg) => [{ type: ProfileTag, id: arg.id }]
    }),
    getMembersByParam: builder.query<GetAllUsers["responseBody"], GetAllUsers["requestQuery"]>({
      query: (query) => ({
        url: `getAllProfiles${ buildQuery(query) }`,
        method: "GET",
      }),
      providesTags: (result = { profiles: [] }) =>
        [
          ProfileTag,
          ...result.profiles.map((profile: GetSimpleUserDataInterface) => ({ type: ProfileTag, id: profile.alias }) as const)
        ]
    }),
    createUser: builder.mutation<CreateMember["responseBody"], CreateMember["requestBody"]>({
      query: (body) => ({
        url: "addProfile",
        method: "POST",
        body: {
          year: 2026,
          ...body
        }
      }),
      invalidatesTags: [ProfileTag],
    })
  })
});

export const {
  useGetUserQuery,
  useLazyGetMembersByParamQuery,
  useCreateUserMutation,
} = memberSlice;