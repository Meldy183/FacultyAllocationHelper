import React from "react";
import Wrapper from "@/shared/ui/wrapper";
import styles from "./styles.module.scss";
import Link from "next/link";

const CourseInformation: React.FC = () => {
	return (
		<Wrapper>
			<div className={styles.container}>
				<div className={styles.card}>
					{/* Header */}
					<div className={styles.header}>
						<div className={styles.userInfo}>
							<div>
								<h1 className={styles.name}>Course name</h1>
								<p className={styles.subName}>Official name / Официальное название курса</p>
							</div>
						</div>
						{/* <Button><Link href={ "/" }>Edit profile</Link></Button> */}

					</div>

					{/* Course Info */}
					{/* <div className={styles.section}>
						<div className={styles.row}><strong>Position:</strong> Professor</div>
						<div className={styles.row}><strong>Institute:</strong> Институт разработки ПО и программной инженерии</div>
					</div> */}

					<div className={styles.section}>
						<h2 className={styles.title}>Course Information</h2>
						<div className={styles.grid}>
							<div>Responsible institute: Institute name</div>
							<div>Program: program code</div>
							<div>Track: track</div>
							<div>Mode: mode</div>
							<div>Form: form</div>
							<div>Semester: semester</div>
						</div>
					</div>

					<div className={styles.section}>
						<h2 className={styles.title}>УП</h2>
						<div className={styles.grid}>
							<div>Lecture hours (per course): N</div>
							<div>Practical class hours: N</div>
						</div>
					</div>
				</div>

				{/* Instructors */}
				<div className={styles.card}>
					<div className={styles.sectionCardWhite}>
						<h2 className={styles.name}>Instructors on this course</h2>
						<h3 className={styles.subName}>Primary instructor</h3>
						<ul className={ styles.list }>
							<li className={styles.primatyTutorItem}>
								<div className={styles.nameAliasPrimaryTutor}>
									<h2>Name Surname</h2>
									<div>@alias</div>
								</div>
								<div className={styles.taEmail}>n.surname@innopolis.university</div>
							</li>
						</ul>
						<h3 className={styles.subName}>Tutor instructor</h3>
				 				<ul className={ styles.list }>
							<li className={styles.primatyTutorItem}>
								<div className={styles.nameAliasPrimaryTutor}>
									<h2>Name Surname</h2>
									<div>@alias</div>
								</div>
								<div className={styles.taEmail}>n.surname@innopolis.university</div>
							</li>
						</ul>

						<h3 className={styles.subName}>Teaching assistants</h3>
						<div className={ styles.assistance }>
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
			 			</div>
							
					</div>
				</div>

				
			</div>
		</Wrapper>
	);
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
				
				
			</li>
		</ul>
	</Link>
}
export default CourseInformation;