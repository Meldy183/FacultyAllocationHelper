import { createApi, fetchBaseQuery } from "@reduxjs/toolkit/query/react";
import { API_PATH } from "@/shared/configs/constants/api/paths";
import { GetFacultyFiltersProcessType } from "@/shared/types/api/profile";
import { GetCourseFilterProcess, RawFilters } from "@/shared/types/api/filters";
import { transformRawFilters } from "@/shared/lib/transformFilter";

export const filterSlice = createApi({
  reducerPath: "api/filter",
  baseQuery: fetchBaseQuery({
    baseUrl: `${ API_PATH }/filter`,
    credentials: "include"
  }),
  endpoints: (builder) => ({
    getCourseFilters: builder.query<GetCourseFilterProcess["responseBody"], GetCourseFilterProcess["requestParams"]>({
      query: () => ({
        url: "/course",
        method: "GET"
      })
    }),
    getProfileFilters: builder.query<GetFacultyFiltersProcessType["responseBody"], GetFacultyFiltersProcessType["requestQuery"]>({
      query: () => ({
        url: "/profile",
        method: "GET",
      }),
      transformResponse: (response: RawFilters) => transformRawFilters(response)
    }),
  })
})

export const {
  useGetCourseFiltersQuery,
  useGetProfileFiltersQuery
} = filterSlice;