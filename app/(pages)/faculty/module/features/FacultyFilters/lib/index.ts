import { FilterGroup, RawFilters } from "@/shared/types/api/filters";

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