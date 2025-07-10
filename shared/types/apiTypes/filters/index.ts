export type RawFilter = { [key: string]: number };
export type RawFilters = { [key: string]: RawFilter[] };
export type RawFiltersResponse = { filters: RawFilters };

export type FilterItem = { name: string; value: number };
export type FilterGroup = { name: string; items: FilterItem[] };