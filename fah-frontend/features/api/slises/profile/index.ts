import { createApi, fetchBaseQuery } from "@reduxjs/toolkit/query/react";
import { API_PATH } from "@/shared/configs/constants/api/paths";
import { GetMemberProcessType, GetAllUsers, CreateMember, GetFiltersType } from "shared/types/api/profile";
import { transformRawFilters } from "@/shared/lib/transformFilter";
import { RawFiltersResponse } from "shared/types/api/filters";
import { buildQuery } from "@/shared/lib/buildQuery";

export const memberSlice = createApi({
  reducerPath: "api/profile",
  baseQuery: fetchBaseQuery({
    baseUrl: `${ API_PATH }/profile/`,
    credentials: "include"
  }),
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
        url: `getUser/${ id }`,
        method: "GET",
      })
    }),
    getMembersByParam: builder.query<GetAllUsers["responseBody"], GetAllUsers["requestQuery"]>({
      query: (query) => ({
        url: `getAllUsers${ buildQuery(query) }`,
        method: "GET",
      })
    }),
    createUser: builder.mutation<CreateMember["responseBody"], CreateMember["requestBody"]>({
      query: (body) => ({
        url: "addUser",
        method: "POST",
        body: body
      })
    })
  })
});

export const {
  useGetUserQuery,
  useLazyGetMembersByParamQuery,
  useCreateUserMutation,
  useGetFiltersQuery
} = memberSlice;