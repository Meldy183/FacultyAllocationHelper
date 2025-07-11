import { createSlice, PayloadAction } from "@reduxjs/toolkit";
import { FilterGroup } from "@/shared/types/apiTypes/filters";

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

      const groupIndex = state.filters.findIndex((g) => g.name === groupName);

      if (groupIndex === -1) {
        state.filters.push({
          name: action.payload.name,
          items: [action.payload.items[0]],
        });
        return;
      }

      const group = state.filters[groupIndex];
      const itemExists = group.items.some((item) => item.name === newItem.name);

      state.filters[groupIndex] = {
        ...group,
        items: itemExists
          ? group.items.filter((item) => item.name !== newItem.name)
          : [...group.items, newItem],
      };
    }
  }
});

export const { toggleFilter } = facultyFilterSlice.actions;
export default facultyFilterSlice.reducer;