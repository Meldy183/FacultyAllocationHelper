import { createSlice, PayloadAction } from "@reduxjs/toolkit";
import { FilterGroup, RawFilter, FilterItem } from "@/shared/types/api/filters";

export type RawFilters = { [key: string]: RawFilter[] };

const initialState: { filters: FilterGroup[] } = {
    filters: []
}

const facultyFilterSlice = createSlice({
    name: "facultyFilter",
    initialState,
    reducers: {
        toggleFilter: (state, action: PayloadAction<FilterGroup>) => {
            const groupName = action.payload.name;
            const newItem = action.payload.items[0];

            const groupIndex = state.filters.findIndex((g: FilterGroup) => g.name === groupName);

            if (groupIndex === -1) {
                state.filters.push({
                    name: action.payload.name,
                    items: [action.payload.items[0]],
                });
                return;
            }

            const group = state.filters[groupIndex];
            const itemExists = group.items.some((item: FilterItem) => item.name === newItem.name);

            state.filters[groupIndex] = {
                ...group,
                items: itemExists
                    ? group.items.filter((item: FilterItem) => item.name !== newItem.name)
                    : [...group.items, newItem],
            };
        }
    }
});

export const { toggleFilter } = facultyFilterSlice.actions;
export default facultyFilterSlice.reducer;