import { FilterGroup, FiltersRequest } from "@/shared/types/api/filters";

export function transformWorkingFilters(filters: FilterGroup[]): FiltersRequest {
	const obj: FiltersRequest = {};
	filters.forEach(filter => {
		obj[filter.name] = filter.items.map(item => item.value.toString());
	})
	return obj;
}