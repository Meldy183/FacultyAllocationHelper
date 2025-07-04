"use client"

import React from "react";
import Image from "next/image";
import settingsIcon from "@/public/icons/svg/settings.svg";
import styles from "./styles.module.scss";

import {
	Dialog,
	DialogContent, DialogDescription, DialogHeader,
	DialogTitle,
	DialogTrigger,
} from "@/shared/ui/dialog";
import { VisuallyHidden } from "@radix-ui/react-visually-hidden";
import CourseDialogMenuContent from "@/entities/course/ui/CourseDialogMenuContent";
import ManageCourseDialogMenu from "@/entities/course/ui/ManageCourseDialogMenu";
import TAElement from "../../../features/ui/course/FacultyMemberBlock";
import AssignNewMember from "../../../features/ui/course/AssignNewMember";

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
			<div className={ styles.information }>
				<div className={ styles.courseName }>
					<Dialog>
						<DialogTrigger className={ "cursor-pointer" }>
							<div className={ styles.name }>{ courseName }</div>
						</DialogTrigger>
						<DialogContent className={ styles.dialogMenu }>
							<VisuallyHidden>
								<DialogHeader>
									<DialogTitle/>
									<DialogDescription/>
								</DialogHeader>
							</VisuallyHidden>
							<CourseDialogMenuContent/>
						</DialogContent>
					</Dialog>
					<Dialog>
						<DialogTrigger>
							<div className={ `${ styles.icon } cursor-pointer` }>
								<Image src={ settingsIcon } alt={ "settings" } className={ styles.icon }/>
							</div>
						</DialogTrigger>
						<DialogContent className={ styles.dialogMenu }>
							<VisuallyHidden>
								<DialogHeader>
									<DialogTitle/>
									<DialogDescription/>
								</DialogHeader>
							</VisuallyHidden>
							<ManageCourseDialogMenu />
						</DialogContent>
					</Dialog>
				</div>
				<ul>
					<li>Study year:</li>
					<li>Semester:</li>
					<li>Study program:</li>
					<li>Institute:</li>
				</ul>
			</div>
			<div className={ styles.instructors }>
				<div className={ styles.instructor }>
					<div className={ styles.title }>PI</div>
					<TAElement { ...PI } />
				</div>
				<div className={ styles.instructor }>
					<div className={ styles.title }>Tutor</div>
					<TAElement { ...PI } />
				</div>
			</div>
		</div>
		<div className={ styles.body }>
			<div className={ styles.title }>Teaching assistants</div>
			<div className={ styles.assistance }>
				<AssignNewMember />
				{ faculties.map((faculty, i) => <TAElement { ...faculty } key={ i }/>) }
			</div>
		</div>
	</div>
}

export default CourseField;