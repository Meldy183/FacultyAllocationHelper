import { createApi, fetchBaseQuery } from "@reduxjs/toolkit/query/react";
import type {
	RegisterProcessType,
	LoginProcessType,
	RefreshProcessType,
	LogoutProcessType 
} from "@/shared/types/api/auth";
import { API_PATH } from "@/shared/configs/constants/api/paths";

export const authSlice = createApi({
	reducerPath: "api/auth",
	baseQuery: fetchBaseQuery({
		baseUrl: `${ API_PATH }/auth`,
		credentials: "include"
	}),
	endpoints: (builder) => ({
		register: builder.mutation<RegisterProcessType["response"], RegisterProcessType["request"]>({
			query: (body) => ({
				url: "/register",
				method: "POST",
				// добавить тело запроса
				body: body
			}),
		}),
		login: builder.mutation<LoginProcessType["response"], LoginProcessType["request"]>({
			query: (body) => ({
				url: "/login",
				method: "POST",
				body: body
			})
		}),
		refresh: builder.mutation<RefreshProcessType["response"], RefreshProcessType["request"]>({
			query: () => ({
				url: "/refresh",
				method: "POST",
				body: {}
			})
		}),
		logout: builder.mutation<LogoutProcessType["response"], LogoutProcessType["request"]>({
			query: () => ({
				url: "/logout",
				method: "POST",
				body: {}
			})
		}),
		session: builder.query({
			query: () => ({
				url: "/session",
				method: "GET",
				body: {}
			})
		})
	})
});

export const {
	useRegisterMutation,
	useLoginMutation,
	useRefreshMutation,
	useLogoutMutation,
	useSessionQuery
} = authSlice;