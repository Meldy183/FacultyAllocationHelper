import React from "react";
import styles from "./styles.module.scss";
import { Button } from "@/shared/ui/button";
import { Dialog, DialogContent, DialogDescription, DialogHeader, DialogTitle, DialogTrigger } from "@/shared/ui/dialog";
import { VisuallyHidden } from "@radix-ui/react-visually-hidden";

interface Props {
	children: React.ReactNode;
}

const AddCourseMenu: React.FC<Props> = ({ children }) => {
	return <Dialog>
		<DialogTrigger asChild>{ children }</DialogTrigger>
		<DialogContent>
			<VisuallyHidden>
				<DialogHeader>
					<DialogTitle />
					<DialogDescription />
				</DialogHeader>
			</VisuallyHidden>
			<div className={ styles.menu }>
				<div className={ styles.content }>
					<div className={ styles.title }>Add a new course</div>
					<div className={ styles.fieldBlock }>
						<div className={ styles.fieldDescription }>Name</div>
						<input type="text" placeholder={ "Enter course name" } className={ styles.input }/>
					</div>
					<div className={ styles.fieldBlock }>
						<div className={ styles.fieldDescription }>Track</div>
						<input type="text" placeholder={ "Enter course track" } className={ styles.input }/>
					</div>
					<div className={ styles.fieldBlock }>
						<div className={ styles.fieldDescription }>Semester</div>
						<input type="text" placeholder={ "Enter year of study + semester" } className={ styles.input }/>
					</div>
					<Button className={ styles.button }>Submit</Button>
				</div>
			</div>
		</DialogContent>
	</Dialog>
}

export default AddCourseMenu;