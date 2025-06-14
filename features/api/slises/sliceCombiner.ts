import { combineReducers, Middleware } from "redux";
import { authSlice, authApiSliceKey } from "@/features/api/slises/authSlice";

export const apiSliceCombiner = combineReducers({
	[authApiSliceKey]: authSlice.reducer,
});

export const apiMiddlewares: Middleware[] = [authSlice.middleware];