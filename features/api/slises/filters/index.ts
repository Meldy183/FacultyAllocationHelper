import { createApi, fetchBaseQuery } from "@reduxjs/toolkit/query/react";
import { API_PATH } from "@/shared/configs/constants/api/paths";
import { GetFiltersType } from "@/shared/types/api/profile";
import { GetCourseFilterProcess, RawFiltersResponse } from "@/shared/types/api/filters";
import { transformRawFilters } from "@/shared/lib/transformFilter";

export const filterSlice = createApi({
  reducerPath: "api/filters",
  baseQuery: fetchBaseQuery({
    baseUrl: `${ API_PATH }/filters`,
    credentials: "include"
  }),
  endpoints: (builder) => ({
    getCourseFilters: builder.query<GetCourseFilterProcess["responseBody"], GetCourseFilterProcess["requestParams"]>({
      query: () => ({
        url: "/course",
        method: "GET"
      })
    }),
    getFilters: builder.query<GetFiltersType["responseBody"], GetFiltersType["requestQuery"]>({
      query: () => ({
        url: "/profile",
        method: "GET",
      }),
      transformResponse: (response: RawFiltersResponse) => transformRawFilters(response)
    }),
  })
})

export const {
  useGetCourseFiltersQuery,
  useGetFiltersQuery
} = filterSlice;