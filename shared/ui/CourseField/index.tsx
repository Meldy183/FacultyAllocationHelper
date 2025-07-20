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
import TAElement from "@/features/ui/course/FacultyMemberBlock";
import AssignNewMember from "@/features/ui/course/AssignNewMember";
import { CourseType } from "@/shared/types/ui/courses";

type Props = CourseType & {}

const CourseField: React.FC<Props> = (props) => {
	console.log(props.ti)

	return <div className={ styles.card }>
		<div className={ styles.header }>
			<div className={ styles.information }>
				<div className={ styles.courseName }>
					<Dialog>
						<DialogTrigger className={ "cursor-pointer" }>
							<div className={ styles.name }>{ props.brief_name }</div>
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
					{
						props.pi?.allocation_status === "assigned" ? <TAElement { ...props.pi } /> : <span>not allocated</span>
					}
				</div>
				<div className={ styles.instructor }>
					<div className={ styles.title }>Tutor</div>
					{
						props.ti?.allocation_status === "assigned" ? <TAElement { ...props.ti } /> : <span>not allocated</span>
					}
				</div>
			</div>
		</div>
		<div className={ styles.body }>
			<div className={ styles.title }>Teaching assistants</div>
			<div className={ styles.assistance }>
				<AssignNewMember />
				{ props.tas?.map(ta => <TAElement { ...ta } key={ ta.profile_data.profile_id } />) }
			</div>
		</div>
	</div>
}

export default CourseField;