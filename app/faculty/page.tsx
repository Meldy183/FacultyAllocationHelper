"use client";

import React from "react";
import { Button } from "@/shared/ui/button";
import Wrapper from "@/shared/ui/wrapper";
import SideBar from "@/shared/ui/wrapper/sidebar";
import SideBarContent from "@/app/faculty/SideBarContent";
import styles from "./styles.module.scss";
import Link from "next/link";
import TeacherAssistance from "@/app/faculty/teacherAssistantField";
import { useGetMembersByParamQuery } from "@/features/api/slises/courses/members";

const AssistantsPage: React.FC = () => {
	const { data, error, isLoading } = useGetMembersByParamQuery([]);

	console.log("render")

	React.useEffect(() => {
		console.log("its works")
	}, [])

	React.useEffect(() => {
		console.log(data);
	}, [data])

	React.useEffect(() => {
		console.log(data, error, isLoading);
	}, [data, error, isLoading])

	return <Wrapper>
		<SideBar hiddenText={ "Filters" }><SideBarContent/></SideBar>
		<div className={ styles.headerContainer }>
			<div className={styles.name}>Faculty list</div>
			<Button className={ styles.button }><Link href={ "faculty/add-faculty" }>Add a new	faculty member</Link></Button>
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
						new Array(15).fill(0).map((_, i) => <TeacherAssistance key={ i } />)
					}
					</ul>
			</div>
	</Wrapper>
}

export default AssistantsPage;