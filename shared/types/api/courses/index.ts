import { CourseType } from "@/shared/types/ui/courses";
import { FilterItem } from "@/shared/types/api/filters";
import { CreateCourseType } from "@/shared/types/resolvers/course";
import { RawFilters } from "@/app/(pages)/faculty/module/features/FacultyFilters";

export type GetCoursesProcess = {
  requestParams: {
    // исправить на: allocation_not_finished
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