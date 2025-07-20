import { configureStore } from "@reduxjs/toolkit";
import counterReducer from "./slices/test";
import facultyFilterReducer from "./slices/filters/faculty";
import { index } from "@/features/api/slises/auth";
import { memberSlice } from "@/features/api/slises/profile";
import { DevMode } from "@/shared/configs/constants/dev/DevMode";
import { coursesSlice } from "@/features/api/slises/courses/insex";

const makeStore = () => configureStore({
	reducer: {
		counter: counterReducer,
		facultyFilters: facultyFilterReducer,
		[index.reducerPath]: index.reducer,
		[memberSlice.reducerPath]: memberSlice.reducer,
		[coursesSlice.reducerPath]: coursesSlice.reducer,
	},
	middleware: getDefaultMiddleware =>
		getDefaultMiddleware().concat(index.middleware, memberSlice.middleware, coursesSlice.middleware),
	devTools: DevMode !== 'production'
});

export const store = makeStore();

export type AppStore = ReturnType<typeof makeStore>;
export type RootState = ReturnType<AppStore["getState"]>;
export type AppDispatch = AppStore["dispatch"];