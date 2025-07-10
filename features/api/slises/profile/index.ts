import { createApi, fetchBaseQuery } from "@reduxjs/toolkit/query/react";
import { API_PATH } from "@/shared/configs/constants";
import { GetMemberProcessType, GetAllUsers, CreateMember, GetFiltersType } from "@/shared/types/apiTypes/members";
import { transformFilters } from "@/shared/lib/transformFilter";
import { RawFiltersResponse } from "@/shared/types/apiTypes/filters";

export const memberSlice = createApi({
  reducerPath: "api/profile",
  baseQuery: fetchBaseQuery({
    baseUrl: `${ API_PATH }/profile/`,
    credentials: "include"
  }),
  endpoints: (builder) => ({
    getFilters: builder.query<GetFiltersType["responseBody"], GetFiltersType["requestParams"]>({
      query: () => ({
        url: "filters",
        method: "GET",
      }),
      transformResponse: (response: RawFiltersResponse) => transformFilters(response)
    }),
    getUser: builder.query<GetMemberProcessType["responseBody"], GetMemberProcessType["requestQuery"]>({
      query: ({ id }) => ({
        url: `getUser/${ id }`,
        method: "GET",
      })
    }),
    getMembersByParam: builder.query<GetAllUsers["responseBody"], GetAllUsers["requestParams"]>({
      query: (query) => ({
        url: `getAllUsers`,
        method: "GET",
        params: query
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
  useGetMembersByParamQuery,
  useCreateUserMutation,
  useGetFiltersQuery
} = memberSlice;