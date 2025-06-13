import React from "react";
import styles from "./styles.module.scss";
import Image from "next/image";
import crossSvg from "@/public/icons/svg/cross.svg";
import { Button } from "@/components/ui/button";

const AddTaPage: React.FC = () => {
	return <>
		<div className={ styles.menu }>
			<div className={ styles.top }>
				<Image className={ styles.image } src={ crossSvg } alt={ "" } />
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
					<div className={ styles.fieldDescription }>Department</div>
					<input type="text" placeholder={ "Enter TA’s department" } className={ styles.input }/>
				</div>
				<Button className={ styles.button }>Submit</Button>
			</div>
		</div>
	</>
}

export default AddTaPage;