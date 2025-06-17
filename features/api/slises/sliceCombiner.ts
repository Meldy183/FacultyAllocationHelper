import { combineReducers, Middleware } from "redux";
import { authSlice } from "@/features/api/slises/authSlice";

//try add reducers to combiner and add it in configureStore
export const apiSliceCombiner = combineReducers({
	[authSlice.reducerPath]: authSlice.reducer,
});

export const apiMiddlewares: Middleware[] = [authSlice.middleware];