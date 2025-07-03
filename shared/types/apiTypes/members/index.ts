import { GroupFilterInterface } from "@/shared/types/apiTypes/filters";

interface WorkloadStatsUniteStatClasses {
  lec: number;
  tut: number;
  lab: number;
  elec: number;
  rate: number;
}

interface WorkloadStatsUniteStat {
  id: string;
  classes: WorkloadStatsUniteStatClasses;
}

interface WorkloadStats {
  uniteStat: WorkloadStatsUniteStat[];
  total: {
    totalLec: number;
    totalTut: number;
    totalLab: number;
    totalElec: number;
    totalRate: number;
  };
}

interface Language {
  language: string;
}

interface Course {
  id: string;
}

export interface UserDataInterface {
  id: string;
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
}

export type GetMemberProcessType = {
  requestBody: {
    id: string;
  },
  responseBody: UserDataInterface,
}

export type GetUsersByFiltersType = {
  requestBody: GroupFilterInterface[],
  responseBody: {
    data: UserDataInterface[]
  },
}