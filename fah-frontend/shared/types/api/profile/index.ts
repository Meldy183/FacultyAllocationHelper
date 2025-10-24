import { FilterGroup } from "@/shared/types/api/filters";
import { CreateFacultyType } from "@/features/CreateFacultyForm";
import {
  CreateSimpleUserDataInterface,
  GetSimpleUserDataInterface,
  UserDataInterface
} from "@/shared/types/ui/faculties";

export type GetMemberProcessType = {
  requestQuery: {
    id: string;
  },
  responseBody: UserDataInterface,
}

export type GetFacultyFiltersProcessType = {
  responseBody: FilterGroup[],
  requestQuery: object
}

export type GetAllUsers = {
  requestQuery: { [key: string]: string[] }
  responseBody: {
    profiles: GetSimpleUserDataInterface[];
  },
}

export type CreateFacultyProcessType = {
  requestBody: CreateFacultyType,
  responseBody: CreateSimpleUserDataInterface
}