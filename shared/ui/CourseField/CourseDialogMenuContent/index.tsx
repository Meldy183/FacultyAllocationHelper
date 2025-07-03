import styles from "./styles.module.scss";
import React from "react";
import Link from "next/link";
import { Accordion, AccordionContent, AccordionItem, AccordionTrigger } from "@/shared/ui/accordion";

const CourseDialogMenuContent: React.FC = () => {
	return <div className={ styles.container }>
		<div className={ styles.card }>
			{/* Header */ }
			<div className={ styles.header }>
				<div className={ styles.userInfo }>
					<div>
						<h1 className={ styles.name }>Course name</h1>
						<p className={ styles.subName }>Official name / Официальное название курса</p>
					</div>
				</div>
				{/* <Button><Link href={ "/" }>Edit profile</Link></Button> */ }

			</div>

			{/* <div className={styles.section}>
						<div className={styles.row}><strong>Position:</strong> Professor</div>
						<div className={styles.row}><strong>Institute:</strong> Институт разработки ПО и программной инженерии</div>
					</div> */ }

			<Accordion className={ styles.section } type="single" collapsible>
				<AccordionItem value="item-1">
					<AccordionTrigger className={ `${ styles.title } cursor-pointer` }>Course Information</AccordionTrigger>
					<AccordionContent className={ styles.grid }>
						<div>Responsible institute: Institute name</div>
						<div>Program: program code</div>
						<div>Track: track</div>
						<div>Mode: mode</div>
						<div>Form: form</div>
						<div>Semester: semester</div>
					</AccordionContent>
				</AccordionItem>
			</Accordion>

			<div className={ styles.section }>
				<h2 className={ styles.title }>УП</h2>
				<div className={ styles.grid }>
					<div>Lecture hours (per course): N</div>
					<div>Practical class hours: N</div>
				</div>
			</div>
		</div>

		{/* Instructors */ }
		<div className={ styles.card }>
			<div className={ styles.sectionCardWhite }>
				<h2 className={ styles.name }>Instructors on this course</h2>
				<h3 className={ styles.subName }>Primary instructor</h3>
				<ul className={ styles.list }>
					<TeacherAssistance/>
				</ul>
				<h3 className={ styles.subName }>Tutor instructor</h3>
				<ul className={ styles.list }>
					<TeacherAssistance/>
				</ul>

				<h3 className={ styles.subName }>Teaching assistants</h3>
				{/* <div className={ styles.assistance }>
					<ul className={ styles.heading }>
						<li>Name</li>
						<li>E-mail</li>
					</ul>
					<ul className={ styles.list }>
						<li><TeacherAssistance/></li>
						<li><TeacherAssistance/></li>
						<li><TeacherAssistance/></li>
						<li><TeacherAssistance/></li>
						<li><TeacherAssistance/></li>
					</ul>
				</div> */}
				<div className={ styles.assistance }>
					<ul className={styles.list}>
						<li className={styles.tableHeader}>
							<div className={styles.colName}>Name, alias</div>
							<div className={styles.colEmail}>Email</div>
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

			</div>
		</div>


	</div>
}

const TeacherAssistance: React.FC = () => {
	return <Link href={ "/faculty/[id]" }>
		<li className={styles.row}>
			<div className={styles.colName}>
			<h2>Name Surname</h2>
			<div>@alias</div>
			</div>
			<div className={styles.colEmail}>n.surname@innopolis.university</div>
		</li>
	</Link>
}

export default CourseDialogMenuContent;