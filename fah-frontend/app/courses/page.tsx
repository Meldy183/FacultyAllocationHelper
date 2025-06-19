"use client";
import React from "react";
import Wrapper from "@/components/ui/wrapper";
import SideBar from "@/components/ui/wrapper/sidebar";
import CourseField from "@/components/ui/CourseField";
import SideBarContent from "@/app/courses/SideBarContent";
import styles from "./styles.module.scss";

const faculty = {
	name: "Name",
	surname: "Surname",
	department: [],
	role: "TA",
	workload: 0.2
}

const courseMock = {
	courseName: "CourseName 1",
	PI: faculty,
	faculties: [faculty, faculty, faculty, faculty, faculty, faculty, faculty, faculty, faculty],
}

const CoursesPage: React.FC = () => {
	return (
		<Wrapper>
			<SideBar hiddenText={ "Filters" }><SideBarContent /></SideBar>
			<div className={ styles.container }>
				<div className={ styles.field }><CourseField { ...courseMock } /></div>
				<div className={ styles.field }><CourseField { ...courseMock } /></div>
				<div className={ styles.field }><CourseField { ...courseMock } /></div>
				<div className={ styles.field }><CourseField { ...courseMock } /></div>
				<div className={ styles.field }><CourseField { ...courseMock } /></div>
			</div>
		</Wrapper>
	)
}


export default CoursesPage;