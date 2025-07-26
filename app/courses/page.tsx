"use client";

import React, { useEffect } from "react";
import Wrapper from "@/shared/ui/wrapper";
import SideBar from "@/shared/ui/wrapper/sidebar";
import CourseField from "@/shared/ui/CourseField";
import SideBarContent from "@/app/courses/SideBarContent";
import styles from "./styles.module.scss";
import { Button } from "@/shared/ui/button";
import AddCourseMenu from "../../features/ui/course/CreateCourseMenu";
import { useLazyGetAllCoursesQuery } from "@/features/api/slises/courses";
import { useAppSelector } from "@/features/store/hooks";
import { useDebounce } from "@/shared/hooks/useDebounce";
import { FilterItem } from "@/shared/types/api/filters";

type rawFilters = {
	allocaion_not_finished: boolean,
	academic_year: FilterItem[],
	semester: FilterItem[],
	study_program: FilterItem[],
	institute: FilterItem[],
}

const transformFilters = (filters: rawFilters) => {
	const searchQueries = new URLSearchParams();

	searchQueries.append("profile_version_id", "");
	searchQueries.append("year", "2026");
	searchQueries.append("allocation_not_finished", filters.allocaion_not_finished.toString());
	filters.academic_year.forEach((academicYear) => {
		searchQueries.append("academic_year_id", academicYear.value.toString());
	})
	filters.semester.forEach((semester) => {
		searchQueries.append("semester_ids", semester.value.toString());
	})
	filters.study_program.forEach((studyProgram) => {
		searchQueries.append("study_program_ids", studyProgram.value.toString());
	})
	filters.institute.forEach((institute) => {
		searchQueries.append("responsible_institute_ids", institute.value.toString());
	})
	return searchQueries.toString();
}

const CoursesPage: React.FC = () => {
	const filters = useAppSelector(state => state.courseFilters.filters);
	const [getCourses, { data, error }] = useLazyGetAllCoursesQuery();

	const debouncedFilters = useDebounce(filters);

	useEffect(() => {
		getCourses(transformFilters(debouncedFilters));
	}, [debouncedFilters, getCourses]);

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