import { FilterGroup, RawFiltersResponse } from "@/shared/types/apiTypes/filters";

export function transformFilters(raw: RawFiltersResponse): FilterGroup[] {
	return Object.entries(raw.filters).map(([name, rawItems]) => ({
		name,
		items: rawItems.map(rawItem => {
			const [itemName, value] = Object.entries(rawItem)[0];
			return { name: itemName, value };
		})
	}));
}