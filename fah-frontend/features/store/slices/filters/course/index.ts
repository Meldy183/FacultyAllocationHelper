import { FilterItem } from "@/shared/types/api/filters";
import { createSlice, PayloadAction } from "@reduxjs/toolkit";

type initialStateType = {
  filters: {
    allocaion_not_finished: boolean,
    academic_year: FilterItem[],
    semester: FilterItem[],
    study_program: FilterItem[],
    institute: FilterItem[],
  }
}

const initialState: initialStateType = {
  filters: {
    allocaion_not_finished: false,
    academic_year: [],
    semester: [],
    study_program: [],
    institute: []
  }
}

export interface CourseFiltersGroup {
  name: "academic_year" | "semester" | "study_program" | "institute";
  items: FilterItem[]
}

const courseFilterSlice = createSlice({
  name: "courseFilters",
  initialState,
  reducers: {
    toggleIsAllocated: (state) => {
      state.filters.allocaion_not_finished = !state.filters.allocaion_not_finished;
    },
    toggleFilters: (state, action: PayloadAction<CourseFiltersGroup>) => {
      const groupName = action.payload.name;
      const newItem = action.payload.items[0];

      const isContains = state.filters[groupName].some((g) => g.name === newItem.name);
      if (isContains) {
        state.filters[groupName] = state.filters[groupName].filter(item => item.name !== newItem.name);
        return;
      }

      state.filters[groupName].push(newItem);
    }
  }
})

export const {
  toggleIsAllocated,
  toggleFilters,
} = courseFilterSlice.actions;
export default courseFilterSlice.reducer;