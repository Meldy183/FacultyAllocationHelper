import React from "react";
import styles from "@/features/ui/course/AssignNewMember/styles.module.scss";
import { Input } from "@/shared/ui/input";
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "@/shared/ui/select";
import { Button } from "@/shared/ui/button";

interface IProps {
	position: string;
	handleFoldMenu: () => void;
}

const CreateNewMemberMenu: React.FC<IProps> = ({ position, handleFoldMenu }) => {
	return <div className={ styles.createNewMember }>
		<div className={styles.head}>
			<div className={ styles.title }>Create new member and allocate them in course on position: { position }</div>
		</div>
		<div className={styles.commonContainer}> 

		<div className={ styles.memberField }>
			<div className={ styles.fieldDescription }>Name</div>
			<Input placeholder={ "Enter the member’s name" }/>
		</div>
		<div className={ styles.memberField }>
			<div className={ styles.fieldDescription }>E-mail</div>
			<Input placeholder={ "Enter member’s email" }/>
		</div>
		<div className={ styles.memberField }>
			<div className={ styles.fieldDescription }>Alias</div>
			<Input placeholder={ "Enter the member’s alias" }/>
		</div>
		<div className={ styles.memberField }>
			<div className={ styles.fieldDescription }>Institute</div>
			<Input placeholder={ "Enter the member’s institute" }/>
		</div>
		<div className={ styles.memberField }>
			<div className={ styles.fieldDescription }>Position</div>
			<Select>
				<SelectTrigger className={ styles.selectValue }>
					<SelectValue placeholder={ "Enter the member’s position" } />
				</SelectTrigger>
				<SelectContent>
					<SelectItem value="TA">TA</SelectItem>
					<SelectItem value="PI">PI</SelectItem>
					<SelectItem value="TI">TI</SelectItem>
				</SelectContent>
			</Select>
		</div>
		<div className={ styles.memberField }>
			<Button onClick={ () => handleFoldMenu() } className={ styles.foldBtn }>Fold menu</Button>
		</div>
		<div className={ styles.memberField }>
			<Button onClick={ () => handleFoldMenu() } className={ styles.submitBtn }>Confirm</Button>
		</div>
		</div>
	</div>
}

export default CreateNewMemberMenu;