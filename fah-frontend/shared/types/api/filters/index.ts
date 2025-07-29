export type RawFilter = {
	id: number,
	name: string
};

export type RawFilters = { [key: string]: RawFilter[] };

export type FilterItem = { name: string; value: number };
export type FilterGroup = { name: string; items: FilterItem[] };

export type FiltersRequest = { [key: string]: string[] };

export type GetCourseFilterProcess = {
	requestParams: {},
	responseBody: {
		// filters: {
		allocaion_not_finished: boolean,
		academic_year: RawFilter[],
		semester: RawFilter[],
		study_program: RawFilter[],
		institute: RawFilter[],
		// }
	}
}