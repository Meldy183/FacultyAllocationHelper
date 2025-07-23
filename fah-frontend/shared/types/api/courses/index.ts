import { CourseType } from "@/shared/types/ui/courses";
import { FilterItem, RawFilters } from "@/shared/types/api/filters";
import { CreateCourseType } from "@/shared/types/resolvers/course";

export type GetCoursesProcess = {
  requestParams: {
    allocation_finished: boolean,
    academic_year: FilterItem[],
    semester_ids: FilterItem[],
    responsible_institute_ids: FilterItem[],
  },
  responseBody: {
    courses: CourseType[]
  }
}

export type CreateCourseProcess = {
  requestBody: CreateCourseType,
  responseBody: {}
}

export type CoursesFiltersProcess = {
  requestQuery: {},
  responseBody: {
    allocation_finished: boolean,
  } & RawFilters
}