import { FilterItem } from "@/shared/types/api/filters";
import { createSlice, PayloadAction } from "@reduxjs/toolkit";

export enum filtersEnum {
    academic_year = "academic_year",
    semester = "semester",
    study_program = "study_program",
    institute = "institute"
};

type initialStateType = {
    allocaion_not_finished: boolean,
    academic_year: FilterItem[],
    semester: FilterItem[],
    study_program: FilterItem[],
    institute: FilterItem[],
}

const initialState: initialStateType = {
    allocaion_not_finished: false,
    academic_year: [],
    semester: [],
    study_program: [],
    institute: []
}

const courseFiltersSlice = createSlice({
    name: "courseFilters",
    initialState,
    reducers: {
        toggleIsAllocated: (state) => {
            state.allocaion_not_finished = !state.allocaion_not_finished;
        },
        toggleFilters: (state, action: PayloadAction<{ name: filtersEnum; items: FilterItem[] }>) => {
            const groupName = action.payload.name;
            const newItem = action.payload.items[0];

            const isContains = state[groupName].some((g) => g.name === newItem.name);
            if (isContains) {
                state[groupName] = state[groupName].filter(item => item.name !== newItem.name);
                return;
            }

            state[groupName].push(newItem);
        }
    }
})

export const {
    toggleIsAllocated,
    toggleFilters,
} = courseFiltersSlice.actions;
export default courseFiltersSlice.reducer;