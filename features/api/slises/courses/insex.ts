import { createApi, fetchBaseQuery } from "@reduxjs/toolkit/query/react";
import { API_PATH } from "@/shared/configs/constants/api/paths";
import { GetCoursesProcess } from "@/shared/types/api/courses";

export const coursesSlice = createApi({
  reducerPath: "api/courses",
  baseQuery: fetchBaseQuery({
    baseUrl: `${ API_PATH }/academic`,
    credentials: "include"
  }),
  endpoints: (builder) => ({
    getAllCourses: builder.query<GetCoursesProcess["responseBody"], GetCoursesProcess["requestParams"]>({
      query: (body) => ({
        url: "getCourseList",
        method: "GET"
      })
    })
  })
});

export const {
  useGetAllCoursesQuery
} = coursesSlice;