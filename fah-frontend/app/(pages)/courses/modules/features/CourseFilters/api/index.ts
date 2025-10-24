import { filterSlice } from "@/features/api/slises/filters";
import { GetCourseFilterProcess } from "@/shared/types";

const courseFiltersEndpoints = filterSlice.injectEndpoints({
    endpoints: (builder) => ({
        getCourseFilters: builder.query<GetCourseFilterProcess["responseBody"], GetCourseFilterProcess["requestParams"]>({
            query: () => ({
                url: "/course",
                method: "GET"
            })
        })
    })
})

export const { useGetCourseFiltersQuery } = courseFiltersEndpoints;