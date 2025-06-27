import { configureStore } from "@reduxjs/toolkit";
import counterReducer from "./slices/test";
import { authSlice } from "@/features/api/slises/authSlice";

const makeStore = () => configureStore({
	reducer: {
		counter: counterReducer,
		[authSlice.reducerPath]: authSlice.reducer,
	},
	middleware: getDefaultMiddleware =>
		getDefaultMiddleware().concat(authSlice.middleware),
});

export const store = makeStore();

export type AppStore = ReturnType<typeof makeStore>;
export type RootState = ReturnType<AppStore["getState"]>;
export type AppDispatch = AppStore["dispatch"];