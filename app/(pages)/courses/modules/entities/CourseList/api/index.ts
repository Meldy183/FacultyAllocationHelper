import { coursesSlice } from "@/features/api/slises/courses";
import { GetCoursesProcess } from "@/shared/types";

export const { useLazyGetAllCoursesQuery } = coursesSlice.injectEndpoints({
    endpoints: (builder) => ({
        getAllCourses: builder.query<GetCoursesProcess["responseBody"], string>({
            query: (body) => ({
                url: `getCourseList?${ body }`,
                method: "GET"
            }),
            providesTags: ["apiCourses"]
        }),
    })
})