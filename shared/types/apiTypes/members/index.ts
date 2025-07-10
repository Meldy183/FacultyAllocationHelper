import { FilterGroup } from "@/shared/types/apiTypes/filters";

type Language = {
  language: string;
};

type Course = {
  id: string;
};

type WorkloadStatsClasses = {
  lec: number;
  tut: number;
  lab: number;
  elec: number;
  rate: number;
};

type WorkloadStatsUniteStat = {
  id: string;
  classes: WorkloadStatsClasses;
};

type WorkloadStatsTotal = {
  totalLec: number;
  totalTut: number;
  totalLab: number;
  totalElec: number;
  totalRate: number;
};

type WorkloadStats = {
  uniteStat: WorkloadStatsUniteStat[];
  total: WorkloadStatsTotal;
};

export interface UserDataInterface {
  profile_id: number;
  nameEng: string;
  nameRu: string;
  alias: string;
  email: string;
  position: string;
  institute: string;
  workload: number;
  studentType: string;
  degree: boolean;
  FSRO: string;
  languages: Language[];
  courses: Course[];
  employnmentType: string;
  hiringStatus: string;
  mode: string;
  maxLoad: number;
  frontalHours: number;
  extraActivities: number;
  workloadStats: WorkloadStats;
};

export type GetMemberProcessType = {
  requestQuery: {
    id: string;
  },
  responseBody: UserDataInterface,
}

export type GetFiltersType = {
  responseBody: FilterGroup[],
  requestParams: {}
}

export type GetAllUsers = {
  requestParams: { [key: string]: string[] }
  responseBody: {
    data: UserDataInterface[]
  },
}

export type CreateMember = {
  requestBody: {
    nameEng: string,
    email: string,
    alias: string,
    institute: string,
    position: string,
  },
  responseBody: {
    message: string
  }
}