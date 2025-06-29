import React from "react";
import styles from "./styles.module.scss";
import Link from "next/link";
import Image from "next/image";
import crossIcon from "@/public/icons/svg/cross.svg";
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "@/shared/ui/select";

const ManageCourseDialogMenu: React.FC = () => {
	return <div className={ styles.container }>
		<div className={ styles.content }>
			<form>
				<div className={ styles.element }>
					<div className={ styles.title }>Course name</div>
					<input type="text" placeholder={ "(preset name of the chosen course) " } className={ styles.input }/>
				</div>
				<div className={ styles.element }>
					<div className={ styles.title }>Primary Instructor</div>
					<CustomSelect/>
				</div>
				<div className={ styles.element }>
					<div className={ styles.title }>Tutor Instructor</div>
					<CustomSelect/>
				</div>
				<div className={ styles.element }>
					<div className={ styles.title }>Teaching assistant</div>
					<CustomSelect/>
				</div>
				<div className={ styles.element }>
					<div className={ styles.title }>Teaching assistant</div>
					<CustomSelect/>
				</div>
				<div className={ styles.element }>
					<div className={ styles.title }>Teaching assistant</div>
					<CustomSelect/>
				</div>
			</form>
		</div>
	</div>
}

const CustomSelect: React.FC = () => {
	return (
		<Select>
			<SelectTrigger className={ styles.select }>
				<SelectValue placeholder="Blank if not chosen / current instructor" />
			</SelectTrigger>
			<SelectContent>
				<SelectItem value="Add a new faculty member">Add a new faculty member</SelectItem>
				<SelectItem value="Not assigned">Not assigned</SelectItem>
				<SelectItem value="Not needed">Not needed</SelectItem>
			</SelectContent>
		</Select>
	)
}

export default ManageCourseDialogMenu;