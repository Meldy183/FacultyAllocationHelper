import React from "react";
import styles from "./styles.module.scss";
import Image from "next/image";
import crossSvg from "@/public/icons/svg/cross.svg";
import { Button } from "@/components/ui/button";
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "@/components/ui/select";
import Link from "next/link";

const AddTaPage: React.FC = () => {
	return <>
		<div className={ styles.menu }>
			<div className={ styles.top }>
				<Link href={ "/dashboard" }><Image className={ styles.image } src={ crossSvg } alt={ "" } /></Link>
			</div>
			<div className={ styles.content }>
				<div className={ styles.title }>Add a TA</div>
				<div className={ styles.fieldBlock }>
					<div className={ styles.fieldDescription }>Name</div>
					<input type="text" placeholder={ "Enter the name" } className={ styles.input }/>
				</div>
				<div className={ styles.fieldBlock }>
					<div className={ styles.fieldDescription }>E-mail</div>
					<input type="text" placeholder={ "Enter TA’s email" } className={ styles.input }/>
				</div>
				<div className={ styles.fieldBlock }>
					<div className={ styles.fieldDescription }>Alias</div>
					<input type="text" placeholder={ "Enter the member’s alias" } className={ styles.input }/>
				</div>
				<div className={ styles.fieldBlock }>
					<div className={ styles.fieldDescription }>Phone number</div>
					<input type="text" placeholder={ "Enter the member’s phone number" } className={ styles.input }/>
				</div>
				<div className={ styles.fieldBlock }>
					<div className={ styles.fieldDescription }>Department</div>
					<input type="text" placeholder={ "Enter the member’s department" } className={ styles.input }/>
				</div>
				<CustomSelect />
				<Link href={ "/faculty" }><Button className={ styles.button }>Submit</Button></Link>
			</div>
		</div>
	</>
}

const CustomSelect: React.FC = () => {
	return (
		<Select>
			<SelectTrigger className={ styles.select }>
				<SelectValue placeholder="Enter the member’s position" />
			</SelectTrigger>
			<SelectContent>
				<SelectItem value="Primary instructor">Primary instructor</SelectItem>
				<SelectItem value="Tutor instructor">Tutor instructor</SelectItem>
				<SelectItem value="teacher assistant">teacher assistant</SelectItem>
			</SelectContent>
		</Select>
	)
}

export default AddTaPage;