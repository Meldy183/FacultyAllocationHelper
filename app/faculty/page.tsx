"use client";

import React from "react";
import { Button } from "@/shared/ui/button";
import Wrapper from "@/shared/ui/wrapper";
import SideBar from "@/shared/ui/wrapper/sidebar";
import SideBarContent from "@/app/faculty/SideBarContent";
import styles from "./styles.module.scss";
import Link from "next/link";

const AssistantsPage: React.FC = () => {
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
					<TeacherAssistance/>
					<TeacherAssistance/>					
					<TeacherAssistance/>
					<TeacherAssistance/>
					<TeacherAssistance/>					
					<TeacherAssistance/>
					<TeacherAssistance/>
					<TeacherAssistance/>
					<TeacherAssistance/>
					<TeacherAssistance/>
					</ul>
			</div>
	</Wrapper>
}


const TeacherAssistance: React.FC = () => {
	return <Link href={ "/faculty/faculty-member" }>
		<li className={styles.row}>
			<div className={styles.colName}>
			<h2>Name Surname</h2>
			<div>@alias</div>
			</div>
			<div className={styles.colEmail}>n.surname@innopolis.university</div>
			<div className={styles.colInstitute}>Institute</div>
			<div className={styles.colPosition}>Position</div>
		</li>
	</Link>
}



// const Track: React.FC = () => {
// 	return <Link href={ "faculty/faculty-member" }>
// 		<div className={ styles.course }>
// 			<div className={ styles.element }>Teaching assistant 1</div>
// 			<div className={ styles.element }>t.assistant@innopolis.university</div>
// 			<div className={ styles.element }>Robotics</div>
// 			<div className={ styles.element }>8 (914)-888-15-36</div>
// 			<div className={ styles.element }>0.25</div>
// 		</div>
// 	</Link>
// }

export default AssistantsPage;