import { createApi, fetchBaseQuery } from "@reduxjs/toolkit/query/react";
import type {
	RegisterProcessType,
	LoginProcessType,
	RefreshProcessType,
	LogoutProcessType 
} from "@/types/apiTypes/auth";
import { API_PATH } from "@/configs/constants";

export const authSlice = createApi({
	reducerPath: `/api`,
	baseQuery: fetchBaseQuery(
		{
			baseUrl: API_PATH,
			credentials: "include"
		}),
	endpoints: (builder) => ({
		//позже типизировать как builder.query<типы>
		register: builder.query<RegisterProcessType["response"], RegisterProcessType["request"]>({
			query: (body) => ({
				url: "/register",
				method: "POST",
				// добавить тело запроса
				body: body
			}),
		}),
		login: builder.query<LoginProcessType["response"], LoginProcessType["request"]>({
			query: (body) => ({
				url: "/login",
				method: "POST",
				body: body
			})
		}),
		refresh: builder.query<RefreshProcessType["response"], RefreshProcessType["request"]>({
			query: () => ({
				url: "/refresh",
				method: "POST",
				body: {}
			})
		}),
		logout: builder.query<LogoutProcessType["response"], LogoutProcessType["request"]>({
			query: () => ({
				url: "/logout",
				method: "POST",
				body: {}
			})
		})
	})
});

export const authApiSliceKey = "auth";