import { FilterGroup } from "@/shared/types/api/filters";
import { CreateMemberType } from "@/shared/types/resolvers/profile";
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
  requestQuery: {}
}

export type GetAllUsers = {
  requestQuery: { [key: string]: string[] }
  responseBody: {
    profiles: GetSimpleUserDataInterface[];
  },
}

export type CreateMember = {
  requestBody: CreateMemberType,
  responseBody: CreateSimpleUserDataInterface
}