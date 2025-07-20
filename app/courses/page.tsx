"use client";
import React from "react";
import Wrapper from "@/shared/ui/wrapper";
import SideBar from "@/shared/ui/wrapper/sidebar";
import CourseField from "@/shared/ui/CourseField";
import SideBarContent from "@/app/courses/SideBarContent";
import styles from "./styles.module.scss";
import { Button } from "@/shared/ui/button";
import AddCourseMenu from "../../features/ui/course/CreateCourseMenu";
import { useGetAllCoursesQuery } from "@/features/api/slises/courses/insex";

const CoursesPage: React.FC = () => {
	const { data, error } = useGetAllCoursesQuery({
		allocation_finished: false
	});

	React.useEffect(() => {
		console.log(data);
	}, [data])

	return (
		<Wrapper>
			<SideBar hiddenText={ "Filters" }>
				<SideBarContent />
			</SideBar>
			<div className={ styles.headerContainer }>
				<div className={styles.name}>Courses</div>
				<AddCourseMenu><Button className={ styles.button }>Add a new course</Button></AddCourseMenu>
			</div>
			<div className={ styles.container }>
				<div className={ styles.courses }>
					{
						error && <span>smth went wrong. cant load courses</span>
					}
					{
						data?.courses.map((course) => (
							<div key={ course.course_id } className={ styles.field }><CourseField { ...course } /></div>
						))
					}
				</div>
			</div>
		</Wrapper>
	)
}


export default CoursesPage;