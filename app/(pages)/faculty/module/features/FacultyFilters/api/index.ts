import { filterSlice } from "@/features/api/slises/filters";
import { transformRawFilters } from "../lib";
import { GetFacultyFiltersProcessType } from "@/shared/types/api/profile";
import { RawFilters } from "../models";

const facultyFiltersEndpoints = filterSlice.injectEndpoints({
    endpoints: (builder) => ({
        getProfileFilters: builder.query<GetFacultyFiltersProcessType["responseBody"], GetFacultyFiltersProcessType["requestQuery"]>({
            query: () => ({
                url: "/profile",
                method: "GET",
            }),
            transformResponse: (response: RawFilters) => transformRawFilters(response),
        }),
    }),
    overrideExisting: true
})

export const { useGetProfileFiltersQuery } = facultyFiltersEndpoints;