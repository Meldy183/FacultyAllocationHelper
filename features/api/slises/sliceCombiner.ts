import { combineReducers, Middleware } from "redux";
import { index } from "@/features/api/slises/auth";

//try add reducers to combiner and add it in configureStore
export const apiSliceCombiner = combineReducers({
	[index.reducerPath]: index.reducer,
});

export const apiMiddlewares: Middleware[] = [index.middleware];