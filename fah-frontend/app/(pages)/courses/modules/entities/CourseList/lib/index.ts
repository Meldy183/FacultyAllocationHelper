import { FilterItem } from "@/shared/types";

type rawFilters = {
    allocaion_not_finished: boolean,
    academic_year: FilterItem[],
    semester: FilterItem[],
    study_program: FilterItem[],
    institute: FilterItem[],
}

export const transformFilters = (filters: rawFilters) => {
    const searchQueries = new URLSearchParams();

    searchQueries.append("profile_version_id", "");
    searchQueries.append("year", "2026");
    searchQueries.append("allocation_not_finished", filters.allocaion_not_finished.toString());
    filters.academic_year.forEach((academicYear) => {
        searchQueries.append("academic_year_id", academicYear.value.toString());
    })
    filters.semester.forEach((semester) => {
        searchQueries.append("semester_ids", semester.value.toString());
    })
    filters.study_program.forEach((studyProgram) => {
        searchQueries.append("study_program_ids", studyProgram.value.toString());
    })
    filters.institute.forEach((institute) => {
        searchQueries.append("responsible_institute_ids", institute.value.toString());
    })
    return searchQueries.toString();
}