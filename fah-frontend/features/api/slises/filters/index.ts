import { createApi, fetchBaseQuery } from "@reduxjs/toolkit/query/react";
import { API_PATH } from "@/shared/configs/constants/api/paths";

export const filterSlice = createApi({
  reducerPath: "api/filter",
  baseQuery: fetchBaseQuery({
    baseUrl: `${ API_PATH }/filter`,
    credentials: "include"
  }),
  endpoints: () => ({})
})