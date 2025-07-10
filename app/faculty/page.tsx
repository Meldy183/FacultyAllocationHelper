"use client";

import React from "react";
import Wrapper from "@/shared/ui/wrapper";
import SideBar from "@/shared/ui/wrapper/sidebar";
import SideBarContent from "@/app/faculty/SideBarContent";
import styles from "./styles.module.scss";
import TeacherAssistance from "@/app/faculty/teacherAssistantField";
import CreateFacultyMenu from "../../features/ui/faculty/CreateNewFaculty";
import { useGetMembersByParamQuery } from "@/features/api/slises/profile";

const AssistantsPage: React.FC = () => {
	const { data, error, isLoading } = useGetMembersByParamQuery([]);

	if (error) return <>smth went wrong</>

	return <Wrapper>
		<SideBar hiddenText={ "Filters" }><SideBarContent/></SideBar>
		<div className={ styles.headerContainer }>
			<div className={styles.name}>Faculty list</div>
			<CreateFacultyMenu />
		</div>
		<div className={ styles.assistance }>
				<ul className={styles.list}>
					<li className={styles.header}>
						<div className={styles.colName}>Name, alias</div>
						<div className={styles.colEmail}>Email</div>
						<div className={styles.colInstitute}>Institute</div>
						<div className={styles.colPosition}>Position</div>
					</li>
					{
						isLoading ? <>data loading</> : data?.data.map((item, i) => <TeacherAssistance {...item} key={ i } />)
					}
					</ul>
			</div>
	</Wrapper>
}

export default AssistantsPage;