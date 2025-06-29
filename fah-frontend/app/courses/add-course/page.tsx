import React from "react";
import styles from "./styles.module.scss";
import Image from "next/image";
import crossSvg from "@/public/icons/svg/cross.svg";
import { Button } from "@/shared/ui/button";

const AddCoursePage: React.FC = () => {
	return <>
		<div className={ styles.menu }>
			<div className={ styles.top }>
				<Image className={ styles.image } src={ crossSvg } alt={ "" } />
			</div>
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
	</>
}

export default AddCoursePage;