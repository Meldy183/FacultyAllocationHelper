import { FilterGroup, FiltersRequest, RawFilters } from "shared/types/api/filters";

export function transformRawFilters(raw: RawFilters): FilterGroup[] {
	const filters = Object.entries(raw).map(([name, rawItems]) => ({
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
		obj[filter.name] = filter.items.map(item => item.value.toString());
	})
	return obj;
}