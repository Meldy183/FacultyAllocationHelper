import { CreateFacultyProcessType } from "@/shared/types/api/profile";
import { ProfileTag } from "@/shared/configs/constants/dev/cache/tags/profile";
import { memberSlice } from "@/features/api/slises/profile";

const createFacultyEndpoint = memberSlice.injectEndpoints({
    endpoints: (builder) => ({
        createFaculty: builder.mutation<CreateFacultyProcessType["responseBody"], CreateFacultyProcessType["requestBody"]>({
            query: (body) => ({
                url: "addProfile",
                method: "POST",
                body: {
                    year: 2026,
                    ...body
                }
            }),
            invalidatesTags: [ProfileTag],
        })
    })
})

export const { useCreateFacultyMutation } = createFacultyEndpoint;