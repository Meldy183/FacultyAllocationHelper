"use client";

import React, { Suspense } from "react";
import Wrapper from "@/shared/ui/wrapper";
import SideBar from "@/shared/ui/wrapper/sidebar";
import { CourseList } from "./modules/entities/CourseList";
import { CourseFilters } from "./modules/features/CourseFilters";
import { Button } from "@/shared/ui/button";
import { CreateCourseForm } from "@/features/CreateCourseForm";
import { Dialog, DialogContent, DialogDescription, DialogHeader, DialogTitle, DialogTrigger } from "@/shared/ui/dialog";
import { VisuallyHidden } from "@radix-ui/react-visually-hidden";
import styles from "./styles.module.scss";

const CoursesPage: React.FC = () => {
	return (
		<Wrapper>
			<SideBar hiddenText={ "Filters" }>
				<CourseFilters />
			</SideBar>
			<div className={ styles.headerContainer }>
				<div className={styles.name}>Courses</div>
				<Dialog>
					<DialogTrigger asChild>
						<Button className={ styles.button }>Add a new course</Button>
					</DialogTrigger>
					<DialogContent>
						<VisuallyHidden>
							<DialogHeader>
								<DialogTitle />
								<DialogDescription />
							</DialogHeader>
						</VisuallyHidden>
						<CreateCourseForm />
					</DialogContent>
				</Dialog>
			</div>
			<CourseList />
		</Wrapper>
	)
}


export default CoursesPage;