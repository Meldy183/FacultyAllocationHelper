import React from "react";
import Image from "next/image";
import Link from "next/link";
import { Tooltip, TooltipContent, TooltipTrigger } from "@/shared/ui/tooltip";
import { Button } from "@/shared/ui/button";
import { TooltipArrow } from "@radix-ui/react-tooltip";
import settingsIcon from "@/public/icons/svg/settings.svg";
import styles from "./styles.module.scss";
import userIcon from "@/public/icons/faculty/faculty-member/faculty-member-icon.svg";

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
					<div className={ styles.name }>{ courseName }</div>
					<div className={ styles.icon }>
						<Image src={ settingsIcon } alt={ "settings" } className={ styles.icon } />
					</div>
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
				{ faculties.map((faculty, i) => <TAElement { ...faculty } key={ i } />) }
			</div>
		</div>
	</div>
}

const TAElement: React.FC<Faculty> = (faculty) => {
	return (
		<>
			<Tooltip>
				<TooltipTrigger>
					<div className={ styles.facultyElement }>
						<span className={ styles.menuTrigger }>{ faculty.name[0] + ". " + faculty.surname }</span>
					</div>
				</TooltipTrigger>
				<TooltipContent side={ "right" } className={ styles.contextMenu }>
					<TooltipArrow className={ `${ styles.arrow } fill-white` }  />
					<div className={ styles.menu }>
						<div className={ styles.header }>
							<div className={ styles.head }>
								<Image src={ userIcon } alt={ "user icon" } className={ styles.userImage } />
								<div className={ styles.information }>
									<div className={ styles.name }>Name Surname</div>
									<div className={ styles.tg }>@alias</div>
								</div>
							</div>
							<div className={ styles.email }>n.surname@innopolis.university</div>
							<div className={ styles.workInformation }>
								<div className={ styles.department }>
									<div className={ styles.placeholder }>Department:</div>
									<div className={ styles.value }>______________</div>
								</div>
								<div className={ styles.department }>
									<div className={ styles.placeholder }>Position:</div>
									<div className={ styles.value }>______________</div>
								</div>
								<div className={ styles.workload }>
									<div className={ styles.number }>0.2</div>
									<div className={ styles.text }>Workload</div>
								</div>
							</div>
						</div>
						<ul className={ styles.labList }>
							<li className={ styles.lab }>AGLA I - 2 lab</li>
							<li className={ styles.lab }>MA I - 3 labs</li>
							<li className={ styles.lab }>AGLA I - 1 tut</li>
							<li className={ styles.lab }>Smt else - 1 lab</li>
							<li className={ styles.lab }>Smt else - 1 lab</li>
						</ul>
					</div>
				</TooltipContent>
			</Tooltip>
		</>
	)
}

export default CourseField;