import { CourseType } from "@/shared/types/ui/courses";

type Language = {
  language_code: string;
};

type WorkloadStatsEntry = {
  lec_hours: number;
  tut_hours: number;
  lab_hours: number;
  elective_hours: number;
  rate: number;
};

export interface UserDataInterface {
  year: number;
  name_eng: string;
  name_ru: string;
  alias: string;
  email: string;
  position_name: string;
  institute_names: string[];
  workload: number;
  student_type: string;
  degree: boolean;
  fsro: string;
  languages: Language[];
  employnment_type: string;
  hiring_status: string;
  mode: string;
  max_load: number;
  frontal_hours: number;
  extra_activities: number;
  workload_stats: {
    t1?: WorkloadStatsEntry;
    t2?: WorkloadStatsEntry;
    t3?: WorkloadStatsEntry;
    total: WorkloadStatsEntry;
  };
  courses: CourseType[];
}

export type CreateSimpleUserDataInterface = {
  year: number,
  name_eng: string,
  email: string,
  alias: string,
  institute_ids: number[],
  position_id: number
}

export type GetSimpleUserDataInterface =  {
  profile_id: 0,
  name_eng: string,
  alias: string,
  email: string,
  position_name: string,
  institute_names: string[],
}
