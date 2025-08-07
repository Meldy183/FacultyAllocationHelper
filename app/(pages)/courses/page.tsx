"use client";

import React from "react";
import Wrapper from "@/shared/ui/wrapper";
import SideBar from "@/shared/ui/wrapper/sidebar";
import { CourseList } from "@/app/(pages)/courses/modules/entities/CourseList";
import { CourseFilters } from "./modules/features/CourseFilters";
import AddCourseMenu from "@/features/ui/course/CreateCourseMenu";
import { Button } from "@/shared/ui/button";
import styles from "./styles.module.scss";

const CoursesPage: React.FC = () => {
	return (
		<Wrapper>
			<SideBar hiddenText={ "Filters" }>
				<CourseFilters />
			</SideBar>
			<div className={ styles.headerContainer }>
				<div className={styles.name}>Courses</div>
				<AddCourseMenu><Button className={ styles.button }>Add a new course</Button></AddCourseMenu>
			</div>
			<CourseList />
		</Wrapper>
	)
}


export default CoursesPage;