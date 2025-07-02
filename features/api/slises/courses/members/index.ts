import { createApi, fetchBaseQuery } from "@reduxjs/toolkit/query/react";
import { API_PATH } from "@/shared/configs/constants";
import { GetMemberProcessType, GetUsersByFiltersType } from "@/shared/types/apiTypes/members";

export const memberSlice = createApi({
  reducerPath: "api/members",
  baseQuery: fetchBaseQuery({
    baseUrl: `${ API_PATH }/profile/`,
    credentials: "include"
  }),
  endpoints: (builder) => ({
    getUser: builder.query<GetMemberProcessType["responseBody"], GetMemberProcessType["requestBody"]>({
      query: ({ id }) => ({
        url: `getUser/${ id }`,
        method: "GET",
        body: {}
      })
    }),
    //подумать, что за тип запроса
    getMembersByParam: builder.query<GetUsersByFiltersType["responseBody"], GetUsersByFiltersType["requestBody"]>({
      query: () => ({
        url: `getUser`,
        method: "GET",
        // body: body
      })
    })
  })
});

export const { useGetUserQuery, useGetMembersByParamQuery } = memberSlice;