import { createApi, fetchBaseQuery } from "@reduxjs/toolkit/query/react";
import { API_PATH } from "@/shared/configs/constants";
import { GetMemberProcessType, GetAllUsers, CreateMember } from "@/shared/types/apiTypes/members";

export const memberSlice = createApi({
  reducerPath: "api/members",
  baseQuery: fetchBaseQuery({
    baseUrl: `${ API_PATH }/profile/`,
    credentials: "include"
  }),
  endpoints: (builder) => ({
    getUser: builder.query<GetMemberProcessType["responseBody"], GetMemberProcessType["requestQuery"]>({
      query: ({ id }) => ({
        url: `getUser/${ id }`,
        method: "GET",
      })
    }),
    //подумать, что за тип запроса
    getMembersByParam: builder.query<GetAllUsers["responseBody"], GetAllUsers["requestBody"]>({
      query: () => ({
        url: `getAllUsers`,
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

export const { useGetUserQuery, useGetMembersByParamQuery, useCreateUserMutation } = memberSlice;