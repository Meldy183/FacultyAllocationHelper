import React from "react";
import { Button } from "@/components/ui/button";
import styles from "./styles.module.scss";
import Link from "next/link";

interface Course {
	courseName: string;
	PI: Faculty;
	tutor?: Faculty;
	faculties: Faculty[];
}

interface Faculty {
	name: string;
	surname: string;
	department: string[];
	role: string;
	workload: number;
}

const CourseField: React.FC<Course> = ({ courseName, PI, tutor, faculties }) => {
	return <div className={ styles.card }>
		<div className={ styles.header }>
			<div className={ styles.title }>{ courseName }</div>
			<Button><Link href={ "/courses/manage" }>Manage</Link></Button>
		</div>
		<div className={ styles.instructors }>
			<div className={ styles.instructor }>
				<div className={ styles.role }> PI</div>
				<div className={ styles.name }> <span>{ PI.name[0] + ". " + PI.surname }</span> </div>
			</div>
			<div className={ styles.instructor }>
				<div className={ styles.role }> Tutor </div>
				<div className={ styles.name }> <span>{ tutor?.name || "not assigned" }</span> </div>
			</div>
		</div>
		<div className={ styles.faculties }>
			<div className={ styles.title }>Faculty</div>
			<div className={ styles.facultyMainBlock }>
				{
					faculties.map((faculty, i) => <div key={ i } className={ styles.facultyBlock }><span>{ faculty.name[0] + ". " + faculty.surname }</span></div>)
				}
			</div>
		</div>
	</div>
}

export default CourseField;