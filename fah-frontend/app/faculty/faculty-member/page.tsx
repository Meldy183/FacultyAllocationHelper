import React from "react";
import Wrapper from "@/components/ui/wrapper";
import Image from "next/image";
import FacultyIcon from "@/public/icons/faculty/faculty-member/faculty-member-icon.svg"
import styles from "./styles.module.scss";
import CourseField from "@/components/ui/CourseField";

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


const FacultyMember: React.FC = () => {
	return <Wrapper>
		<div className={ styles.container }>
			<div className={ styles.content }>
				<div className={ styles.header }>
					<div className={ styles.informationContainer }>
						<Image className={ styles.icon } src={ FacultyIcon } alt={ "faculty image" }/>
						<div className={ styles.name }>
							<div className={ styles.engName }>Name Surname</div>
							<div className={ styles.rusName }>Фамилия Имя Отчество</div>
							<div className={ styles.contacts }>
								<div className={ styles.phoneNumber }>8-xxx-xxx-xx-xx</div>
								<div className={ styles.telegram }>@alias</div>
							</div>
						</div>
						<div className={ styles.email }>n.surname@innopolis.university</div>
					</div>
					<div className={ styles.workSection }>
						<div className={ styles.department }>Department: ________________</div>
						<div className={ styles.position }>Position: ______________</div>
						<div className={ styles.workload }>workload: 0.2</div>
					</div>
					<div className={ styles.courses }>
						<div className={ styles.title }>Teaching courses:</div>
						<div className={ styles.courses }>
							<div className={ styles.course }><CourseField { ...courseMock } /></div>
							<div className={ styles.course }><CourseField { ...courseMock } /></div>
							<div className={ styles.course }><CourseField { ...courseMock } /></div>
							<div className={ styles.course }><CourseField { ...courseMock } /></div>
						</div>
					</div>
				</div>
			</div>
		</div>
	</Wrapper>
}

export default FacultyMember