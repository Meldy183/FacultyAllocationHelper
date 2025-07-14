import { FilterGroup } from "@/shared/types/api/filters";
import { CreateMemberType } from "@/shared/types/resolvers/profile";

type Language = {
  language: string;
};

type Course = {
  id: string;
};

type WorkloadStatsEntry = {
  lec_hours: number;
  tut_hours: number;
  lab_hours: number;
  elective_hours: number;
  rate: number;
};

export interface UserDataInterface {
  name_eng: string;
  name_ru: string;
  alias: string;
  email: string;
  position: string;
  institute: string;
  workload: number;
  student_type: string;
  degree: boolean;
  fsro: string;
  languages: Language[];
  courses: Course[];
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
}


export type GetMemberProcessType = {
  requestQuery: {
    id: string;
  },
  responseBody: UserDataInterface,
}

export type GetFiltersType = {
  responseBody: FilterGroup[],
  requestQuery: {}
}

export type GetAllUsers = {
  requestQuery: { [key: string]: string[] }
  responseBody: {
    profiles: UserDataInterface[]
  },
}

export type CreateMember = {
  requestBody: CreateMemberType,
  responseBody: {
    message: string
  }
}