"use client";
import React from "react";
import Link from "next/link";
import { Button } from "@/components/ui/button";
import SideBar from "@/components/ui/wrapper/sidebar";
import SideBarContent from "@/app/courses/SideBarContent";
import styles from "./styles.module.scss";
import { usePathname } from "next/navigation";

const CoursesPage: React.FC = () => {
	const pathname = usePathname();

	return (
		<>
			<SideBar hiddenText={ "Filters" }><SideBarContent /></SideBar>
			<div className={ styles.container }>
				<div className={ styles.head }>
					<ul>
						<li>Course Name</li>
						<li>Year of study</li>
						<li>Track</li>
						<li>Semester</li>
						<li>Year</li>
						<li><Button className={ styles.button }><Link href={ pathname + "/add-course" }>Add a course</Link></Button></li>
					</ul>
					<div className={ styles.courses }>
						{ new Array(8).fill(0).map((_, i) => <Track key={ i }/>) }
					</div>
				</div>
			</div>
		</>
	)
}

const Track: React.FC = () => {
	const pathname = usePathname();

	return <div className={ styles.course }>
		<div className={ styles.name }>CourseName 1</div>
		<div className={ styles.academicYear }>BS-1</div>
		<div className={ styles.track }>ISE</div>
		<div className={ styles.semester }>Fall</div>
		<div className={ styles.year }>2025</div>
		<Button className={ styles.button }><Link href={ pathname + "/add-ta" }>Add a TA</Link></Button>
	</div>
}

export default CoursesPage;