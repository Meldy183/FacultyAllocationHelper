import { CourseType } from "@/shared/types/ui/courses";

export type GetCoursesProcess = {
  requestParams: {
    allocation_finished: boolean,
  },
  responseBody: {
    courses: CourseType[]
  }
}