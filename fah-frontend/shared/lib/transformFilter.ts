import { FilterGroup, FiltersRequest, RawFiltersResponse } from "shared/types/api/filters";

export function transformRawFilters(raw: RawFiltersResponse): FilterGroup[] {
	const filters = Object.entries(raw.filters).map(([name, rawItems]) => ({
		name,
		items: rawItems.map(rawItem => ({
			name: rawItem.name,
			value: rawItem.id
		}))
	}));

	console.log(filters)

	return filters;
}

export function transformWorkingFilters(filters: FilterGroup[]): FiltersRequest {
	const obj: FiltersRequest = {};
	filters.forEach(filter => {
		obj[filter.name] = filter.items.map(item => item.name);
	})
	return obj;
}