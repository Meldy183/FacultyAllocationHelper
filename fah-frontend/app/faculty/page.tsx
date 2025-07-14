"use client";

import React, { useEffect, useState } from "react";
import Image from "next/image";
import Wrapper from "@/shared/ui/wrapper";
import SideBar from "@/shared/ui/wrapper/sidebar";
import SideBarContent from "@/app/faculty/SideBarContent";
import TeacherAssistance from "@/app/faculty/teacherAssistantField";
import CreateFacultyMenu from "@/features/ui/faculty/CreateNewFaculty";
import { useLazyGetMembersByParamQuery } from "@/features/api/slises/profile";
import { UserDataInterface } from "shared/types/api/profile";
import { useAppSelector } from "@/features/store/hooks";
import { FilterGroup } from "shared/types/api/filters";
import { transformWorkingFilters } from "@/shared/lib/transformFilter";
import { useDebounce } from "@/shared/hooks/useDebounce";
import { debounceTime } from "@/shared/configs/constants/dev/debounceTime";
import loaderSvg from "@/public/icons/svg/loader.svg";
import wrongSvg from "@/public/icons/svg/wrong.svg";
import styles from "./styles.module.scss";

const AssistantsPage: React.FC = () => {
	const filters: FilterGroup[] = useAppSelector(state => state.facultyFilters.filters);
	const [getUsers, { data, error, isError, isLoading }] = useLazyGetMembersByParamQuery();
	const [users, setUsers] = useState<UserDataInterface[]>([]);

	const debouncedFilters = useDebounce(filters, debounceTime);

	useEffect(() => {
		const transformedFilters = transformWorkingFilters(debouncedFilters);
		getUsers(transformedFilters);
	}, [debouncedFilters, getUsers]);

	useEffect(() => {
		if (data) setUsers(data?.profiles || []);
	}, [data, error, isLoading]);

	return <Wrapper>
		<SideBar hiddenText={ "Filters" }><SideBarContent/></SideBar>
		<div className={ styles.headerContainer }>
			<div className={styles.name}>Faculty list</div>
			<CreateFacultyMenu />
		</div>
		<div className={ styles.assistance }>
			{
				isError
					? <div className={ styles.wrongMessage }>
						<div className={ styles.wrongText }>something went wrong: <>{ error && 'data' in error ? (error.data! as { message: string }).message : 'An error occurred' }</></div>
						<Image className={ styles.wrongImage } src={ wrongSvg } alt={ "something went wrong" } />
					</div>
					: <ul className={styles.list}>
						<li className={styles.header}>
							<div className={styles.colName}>Name, alias</div>
							<div className={styles.colEmail}>Email</div>
							<div className={styles.colInstitute}>Institute</div>
							<div className={styles.colPosition}>Position</div>
						</li>
						{
							isLoading ?
								<><Image className={ styles.loadingImage } src={ loaderSvg } alt={ "loading" } /></>
								: users.map((item, i) => <TeacherAssistance {...item} key={ i } />)
						}
					</ul>
			}
		</div>
	</Wrapper>
}

export default AssistantsPage;