import { configureStore } from "@reduxjs/toolkit";
import { facultyFilterReducer } from "@/app/(pages)/faculty/module/features/FacultyFilters";
import { courseFiltersReducer } from "@/app/(pages)/courses/modules/features/CourseFilters";
import { index } from "@/features/api/slises/auth";
import { memberSlice } from "@/features/api/slises/profile";
import { DevMode } from "@/shared/configs/constants/dev/DevMode";
import { coursesSlice } from "@/features/api/slises/courses";
import { filterSlice } from "@/features/api/slises/filters";

const makeStore = () => configureStore({
	reducer: {
		facultyFilters: facultyFilterReducer,
		courseFilters: courseFiltersReducer,
		[index.reducerPath]: index.reducer,
		[memberSlice.reducerPath]: memberSlice.reducer,
		[coursesSlice.reducerPath]: coursesSlice.reducer,
		[filterSlice.reducerPath]: filterSlice.reducer,
	},
	middleware: getDefaultMiddleware =>
		getDefaultMiddleware().concat(index.middleware, memberSlice.middleware, coursesSlice.middleware, filterSlice.middleware),
	devTools: DevMode !== 'production'
});

export const store = makeStore();

export type AppStore = ReturnType<typeof makeStore>;
export type RootState = ReturnType<AppStore["getState"]>;
export type AppDispatch = AppStore["dispatch"];