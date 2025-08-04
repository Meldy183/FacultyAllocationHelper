import { RawFilter } from "@/app/(pages)/faculty/module/features/FacultyFilters";

export type FilterItem = { name: string; value: number };
export type FilterGroup = { name: string; items: FilterItem[] };

export type FiltersRequest = { [key: string]: string[] };

export type GetCourseFilterProcess = {
	requestParams: {},
	responseBody: {
		allocaion_not_finished: boolean,
		academic_year: RawFilter[],
		semester: RawFilter[],
		study_program: RawFilter[],
		institute: RawFilter[],
	}
}