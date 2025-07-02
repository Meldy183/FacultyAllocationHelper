"use client";
import React from "react";
import Wrapper from "@/shared/ui/wrapper";
import SideBar from "@/shared/ui/wrapper/sidebar";
import CourseField from "@/shared/ui/CourseField";
import SideBarContent from "@/app/courses/SideBarContent";
import styles from "./styles.module.scss";
import { Button } from "@/shared/ui/button";
import Link from "next/link";

// import { Sheet, SheetContent, SheetDescription, SheetHeader, SheetTitle, SheetTrigger } from "@/shared/ui/sheet";

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
			<SideBar hiddenText={ "Filters" }>
				<SideBarContent />
			</SideBar>
			<div className={ styles.headerContainer }>
				<div className={styles.name}>Courses</div>
				<Button className={ styles.button }><Link href={ "faculty/add-faculty" }>Add a new course</Link></Button>
			</div>
			<div className={ styles.container }>
				<div className={ styles.courses }>
					<div className={ styles.field }><CourseField { ...courseMock } /></div>
					<div className={ styles.field }><CourseField { ...courseMock } /></div>
					<div className={ styles.field }><CourseField { ...courseMock } /></div>
					<div className={ styles.field }><CourseField { ...courseMock } /></div>
					<div className={ styles.field }><CourseField { ...courseMock } /></div>
					<div className={ styles.field }><CourseField { ...courseMock } /></div>
				</div>
			</div>
		</Wrapper>
	)
}


export default CoursesPage;