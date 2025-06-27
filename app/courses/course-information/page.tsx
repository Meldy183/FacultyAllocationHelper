import React from "react";
import Wrapper from "@/shared/ui/wrapper";
import styles from "./styles.module.scss";
import { Button } from "@/shared/ui/button";
import Link from "next/link";

const CourseInformation: React.FC = () => {
	return <Wrapper>
		<div className={ styles.container }>
			<div className={ styles.content }>
				<div className={ styles.title }>Course information</div>
				<div className={ styles.courseInformation }>
					<ul className={ styles.heading }>
						<li className={ styles.headingElement }>Course Name</li>
						<li className={ styles.headingElement }>Year of study</li>
						<li className={ styles.headingElement }>Track</li>
						<li className={ styles.headingElement }>Semester</li>
						<li className={ styles.headingElement }>Year</li>
						<li className={ styles.headingElement }>Status</li>
					</ul>
					<ul className={ styles.description }>
						<li>CourseName 1</li>
						<li>BS-1</li>
						<li>ISE</li>
						<li>Fall</li>
						<li>2025</li>
						<li><Button className={ styles.button }>Manage</Button></li>
					</ul>
				</div>
				<div className={ styles.assistance }>
					<div className={ styles.title }>Teaching assistants on this course:</div>
					<ul className={ styles.heading }>
						<li>Name</li>
						<li>E-mail</li>
						<li>Department</li>
						<li>Years of teaching</li>
						<li>Workload</li>
					</ul>
					<ul className={ styles.list }>
						<li><TeacherAssistance/></li>
						<li><TeacherAssistance/></li>
						<li><TeacherAssistance/></li>
						<li><TeacherAssistance/></li>
						<li><TeacherAssistance/></li>
					</ul>
				</div>
			</div>
		</div>
	</Wrapper>
}

const TeacherAssistance: React.FC = () => {
	return <Link href={ "/faculty/faculty-member" }>
		<ul className={ styles.assistant }>
			<li>Teaching assistant 1</li>
			<li>t.assistant@innopolis.university</li>
			<li>Robotics</li>
			<li>2017-2025</li>
			<li>0.25</li>
		</ul>
	</Link>
}

export default CourseInformation;