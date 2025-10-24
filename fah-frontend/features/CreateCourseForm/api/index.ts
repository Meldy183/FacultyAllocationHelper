import { coursesSlice } from "@/features/api/slises/courses";
import { CreateCourseProcess } from "@/shared/types";

export const { useCreateNewCourseMutation } = coursesSlice.injectEndpoints({
    endpoints: (builder) => ({
        createNewCourse: builder.mutation<CreateCourseProcess["responseBody"], CreateCourseProcess["requestBody"]>({
            query: (body) => ({
                url: "addNewCourse",
                method: "POST",
                body: body
            }),
            invalidatesTags: ["apiCourses"]
        })
    })
})