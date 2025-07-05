"use client";

import React from "react";
import Wrapper from "@/shared/ui/wrapper";
import SideBar from "@/shared/ui/wrapper/sidebar";
import SideBarContent from "@/app/faculty/SideBarContent";
import styles from "./styles.module.scss";
import TeacherAssistance from "@/app/faculty/teacherAssistantField";
import { useGetMembersByParamQuery } from "@/features/api/slises/courses/members";
import CreateFacultyMenu from "../../features/ui/faculty/CreateNewFaculty";

const user = {
	"nameEng": "Fyodor Markin",
	"nameRu": "Маркин Фёдор Сергеевич",
	"alias": "@meld_i",
	"email": "f.markin@innopolis.university",
	"position": "Intern TA",
	"institute": "Институт разработки ПО и программной инженерии",
	"workload": 0.7,
	"studentType": "MS1",
	"degree": true,
	"FSRO": "employnment",
	"languages": [
		{
			"language": "Russian"
		}
	],
	"courses": [
		{
			"id": "courseInstance_id"
		}
	],
	"employnmentType": "Combination of positions",
	"hiringStatus": "??",
	"mode": "remote",
	"maxLoad": 40,
	"frontalHours": 40,
	"extraActivities": 1.5,
	"workloadStats": {
		"uniteStat": [
			{
				"id": "T1",
				"classes": {
					"lec": 1,
					"tut": 2,
					"lab": 3,
					"elec": 4,
					"rate": 5
				}
			}
		],
		"total": {
			"totalLec": 1,
			"totalTut": 2,
			"totalLab": 3,
			"totalElec": 12,
			"totalRate": 12
		}
	}
}

const data = {
	data: new Array(10).fill(user).map((user, index) => ({...user, id: index}))
}

const AssistantsPage: React.FC = () => {
	// const { data, error, isLoading } = useGetMembersByParamQuery([]);

	// if (error) return <>smth went wrong</>



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
						data?.data.map((item, i) => <TeacherAssistance {...item} key={ i } />)
					}
					</ul>
			</div>
	</Wrapper>
}

export default AssistantsPage;