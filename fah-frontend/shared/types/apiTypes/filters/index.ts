export interface FilterInterface {
  title: string;
  filter_id?: string;
}

export interface GroupFilterInterface {
  group_name: string;
  filters: FilterInterface[];
}