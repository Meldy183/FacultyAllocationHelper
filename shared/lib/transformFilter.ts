import { FilterGroup, FiltersRequest, RawFiltersResponse } from "shared/types/api/filters";

export function transformRawFilters(raw: RawFiltersResponse): FilterGroup[] {
	return Object.entries(raw.filters).map(([name, rawItems]) => ({
		name,
		items: rawItems.map(rawItem => {
			const [itemName, value] = Object.entries(rawItem)[0];
			return { name: itemName, value };
		})
	}));
}

export function transformWorkingFilters(filters: FilterGroup[]): FiltersRequest {
	const obj: FiltersRequest = {};
	filters.forEach(filter => {
		obj[filter.name] = filter.items.map(item => item.name);
	})
	return obj;
}