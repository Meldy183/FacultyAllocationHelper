import { memberSlice } from "@/features/api/slises/profile";
import { GetMemberProcessType } from "@/shared/types";
import { ProfileTag } from "@/shared/configs/constants/dev/cache/tags/profile";

export const { useGetUserQuery } = memberSlice.injectEndpoints({
    endpoints: (builder) => ({
        getUser: builder.query<GetMemberProcessType["responseBody"], GetMemberProcessType["requestQuery"]>({
            query: ({ id }) => ({
                url: `getProfile/${ id }`,
                method: "GET",
            }),
            providesTags: (result, err, arg) => [{ type: ProfileTag, id: arg.id }]
        }),
    })
})