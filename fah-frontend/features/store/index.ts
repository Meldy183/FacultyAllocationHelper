import { configureStore } from "@reduxjs/toolkit";
import counterReducer from "./slices/test";
import facultyFilterReducer from "./slices/filters/faculty";
import { authSlice } from "@/features/api/slises/authSlice";
import { memberSlice } from "@/features/api/slises/profile";

const makeStore = () => configureStore({
	reducer: {
		counter: counterReducer,
		facultyFilters: facultyFilterReducer,
		[authSlice.reducerPath]: authSlice.reducer,
		[memberSlice.reducerPath]: memberSlice.reducer,
	},
	middleware: getDefaultMiddleware =>
		getDefaultMiddleware().concat(authSlice.middleware, memberSlice.middleware),
	devTools: process.env.WORK_MODE !== 'production'
});

export const store = makeStore();

export type AppStore = ReturnType<typeof makeStore>;
export type RootState = ReturnType<AppStore["getState"]>;
export type AppDispatch = AppStore["dispatch"];