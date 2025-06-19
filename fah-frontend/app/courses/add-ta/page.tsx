import React from "react";
import styles from "./styles.module.scss";
import Image from "next/image";
import { Button } from "@/components/ui/button";
import crossSvg from "@/public/icons/svg/cross.svg";
import arrowSvg from "@/public/icons/svg/arrow.svg";

const AddTaPage: React.FC = () => {
	return <>
		<div className={ styles.menu }>
			<div className={ styles.top }>
				<Image className={ styles.image } src={ crossSvg } alt={ "" } />
			</div>
			<div className={ styles.content }>
				<div className={ styles.title }>Add a TA</div>
				<div className={ styles.fieldBlock }>
					<div className={ styles.fieldDescription }>Course name</div>
					<input type="text" placeholder={ "(preset name of the chosen course) " } className={ styles.input }/>
				</div>
				<div className={ styles.fieldBlock }>
					<div className={ styles.fieldDescription }>Choose a TA for the course</div>
					<div className={ `${ styles.input } ${ styles.dropDownList }` }>
						<span>Chosen TA</span>
						<Image className={ styles.image } src={ arrowSvg } alt={ "" } />
					</div>
				</div>
				<Button className={ styles.button }>Submit</Button>
			</div>
		</div>
	</>
}

export default AddTaPage;