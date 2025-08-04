import { createApi, fetchBaseQuery } from "@reduxjs/toolkit/query/react";
import { API_PATH } from "@/shared/configs/constants/api/paths";
import { GetCourseFilterProcess } from "@/shared/types/api/filters";

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
    })
  })
})

export const {
  useGetCourseFiltersQuery,
} = filterSlice;