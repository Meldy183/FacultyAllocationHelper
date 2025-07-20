import { GetSimpleUserDataInterface } from "@/shared/types/ui/faculties";

export type CourseTeacher = {
  allocation_status: string;
  profile_data: GetSimpleUserDataInterface & {
    is_confirmed: boolean;
    classes: string[];
  }
}

export type CourseType = {
  course_id: number;
  brief_name: string;
  official_name: string;
  academic_year_name: string;
  semester_name: string;
  study_program_names: string[];
  responsible_institute_name: string;
  track_names: string[];
  allocation_not_finished: boolean;
  mode: string;
  study_year: number;
  form: string;
  lecture_hours: number;
  lab_hours: number;
  groups_needed: number;
  groups_taken: number;
  pi?: CourseTeacher;
  ti?: CourseTeacher;
  tas: CourseTeacher[];
};