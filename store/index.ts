import { configureStore } from "@reduxjs/toolkit";
import { apiSliceCombiner, apiMiddlewares } from "@/features/api/slises/sliceCombiner";
import counterReducer from "./slices/test";

const makeStore = () => configureStore({
	reducer: {
		counter: counterReducer,
		api: apiSliceCombiner
	},
	middleware: getDefaultMiddleware =>
		getDefaultMiddleware().concat(apiMiddlewares),
});

export const store = makeStore();

export type AppStore = ReturnType<typeof makeStore>;
export type RootState = ReturnType<AppStore["getState"]>;
export type AppDispatch = AppStore["dispatch"];