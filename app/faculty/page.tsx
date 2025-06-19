"use client";

import React from "react";
import { Button } from "@/components/ui/button";
import Wrapper from "@/components/ui/wrapper";
import SideBar from "@/components/ui/wrapper/sidebar";
import SideBarContent from "@/app/courses/SideBarContent";
import styles from "./styles.module.scss";
import Link from "next/link";

const AssistantsPage: React.FC = () => {
	return <Wrapper>
		<SideBar hiddenText={ "Filters" }><SideBarContent/></SideBar>
		<div className={ styles.container }>
			<Button className={ styles.button }><Link href={ "faculty/add-faculty" }>Add a new	faculty member</Link></Button>
			<div className={ styles.head }>
				<ul>
					<li>Name</li>
					<li>E-mail</li>
					<li>Department</li>
					<li>Phone number</li>
					<li>Workload</li>
				</ul>
				<div className={ styles.courses }>
					{ new Array(8).fill(0).map((_, i) => <Track key={ i }/>) }
				</div>
			</div>
		</div>
	</Wrapper>
}

const Track: React.FC = () => {
	return <Link href={ "faculty/faculty-member" }>
		<div className={ styles.course }>
			<div className={ styles.element }>Teaching assistant 1</div>
			<div className={ styles.element }>t.assistant@innopolis.university</div>
			<div className={ styles.element }>Robotics</div>
			<div className={ styles.element }>8 (914)-888-15-36</div>
			<div className={ styles.element }>0.25</div>
		</div>
	</Link>
}

export default AssistantsPage;