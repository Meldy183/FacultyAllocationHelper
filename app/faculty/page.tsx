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
		<div className={ styles.container }>
			<Button className={ styles.button }><Link href={ "faculty/add-faculty" }>Add a new	faculty member</Link></Button>
			{/* <div className={ styles.head }>
				<ul>
					<li>Name</li>
					<li>E-mail</li>
					<li>Department</li>
					<li>Phone number</li>
					<li>Workload</li>
				</ul>
				<div className={ styles.courses }>
					{ new Array(8).fill(0).map((_, i) => <Track key={ i }/>) }
				</div>
			</div> */}
			<div className={ styles.assistance }>
					<ul className={ styles.heading }>
						<li className={styles.nameAliasTA}>Name</li>
						<li className={styles.emailTA}>E-mail</li>
						<li className={styles.instituteTA}>Institute</li>
						<li className={styles.instituteTA}>Position</li>
					</ul>
					<ul className={ styles.list }>
						<li><TeacherAssistance/></li>													 					<li><TeacherAssistance/></li>
						<li><TeacherAssistance/></li>
						<li><TeacherAssistance/></li>
						<li><TeacherAssistance/></li>
						<li><TeacherAssistance/></li>													 					<li><TeacherAssistance/></li>
						<li><TeacherAssistance/></li>


					</ul>
			</div>
		</div>
	</Wrapper>
}


const TeacherAssistance: React.FC = () => {
	return <Link href={ "/faculty/faculty-member" }>
		<ul className={ styles.list }>
			<li className={styles.TaItem}>
				<div className={styles.nameAliasTA}>
					<h2>Name Surname</h2>
					<div>@alias</div>
				</div>
				<div className={styles.emailTA}>n.surname@innopolis.university</div>
				<div className={styles.instituteTA}>institute</div>
				<div className={styles.instituteTA}>position</div>
			</li>
		</ul>
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