"use client";

import React from "react";
import { Button } from "@/components/ui/button";
import SideBar from "@/components/ui/wrapper/sidebar";
import SideBarContent from "@/app/courses/SideBarContent";
import styles from "./styles.module.scss";

const AssistantsPage: React.FC = () => {
	return <>
		<SideBar hiddenText={ "Filters" }><SideBarContent/></SideBar>
		<div className={ styles.container }>
			<Button className={ styles.button }>Add a TA</Button>
			<div className={ styles.head }>
				<ul>
					<li>Name</li>
					<li>E-mail</li>
					<li>Department</li>
					<li>Years of teaching</li>
					<li>Workload</li>
				</ul>
				<div className={ styles.courses }>
					{ new Array(8).fill(0).map((_, i) => <Track key={ i }/>) }
				</div>
			</div>
		</div>
	</>
}

const Track: React.FC = () => {
	return <div className={ styles.course }>
		<div className={ styles.element }>Teaching assistant 1</div>
		<div className={ styles.element }>t.assistant@innopolis.university</div>
		<div className={ styles.element }>Robotics</div>
		<div className={ styles.element }>2017-2025</div>
		<div className={ styles.element }>0.25 (2h/w)</div>
	</div>
}

export default AssistantsPage;