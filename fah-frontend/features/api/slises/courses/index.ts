import { createApi, fetchBaseQuery } from "@reduxjs/toolkit/query/react";
import { API_PATH } from "@/shared/configs/constants/api/paths";
import { CreateCourseProcess, GetCoursesProcess } from "@/shared/types/api/courses";

export const coursesSlice = createApi({
  reducerPath: "api/courses",
  baseQuery: fetchBaseQuery({
    baseUrl: `${ API_PATH }/academic`,
    credentials: "include"
  }),
  tagTypes: ["apiCourses"],
  refetchOnReconnect: true,
  endpoints: () => ({})
});