import React from "react";
import Image from "next/image";
import { StaticImport } from "next/dist/shared/lib/get-img-props";
import bookIcon from "@/public/icons/main-page/book.svg";
import clockIcon from "@/public/icons/main-page/clock.svg";
import userIcon from "@/public/icons/main-page/user.svg";
import styles from "./styles.module.scss";

type elementType = {
	icon: StaticImport,
	title: string,
	description: string,
};

const elementsInformation: elementType[] = [
	{
		icon: bookIcon,
		title: "View courses",
		description: "See information about studying courses of Innopolis University."
	},
	{
		icon: clockIcon,
		title: "Upload Data",
		description: "Use Excel table to upload data about courses and TAs."
	},
	{
		icon: userIcon,
		title: "View TAs",
		description: "See information about teaching assistants of Innopolis University."
	}
]

const Home: React.FC = () => {
	return <div className={ styles.mainScreen }>
		<h1 className={ styles.title }>Welcome!</h1>
		<div className={ styles.previewElements }>
			{
				elementsInformation.map((element, i) => <Element key={ i } icon={element.icon} title={element.title} description={element.description} />)
			}
		</div>
	</div>
}

const Element: React.FC<elementType> = ({ icon, title, description }) => {
	return <div className={ styles.previewElement }>
		<div className={ styles.previewImage }>
			<Image className={ styles.img } src={icon} alt={""} />
		</div>
		<h4 className={ styles.previewTitle }>{ title }</h4>
		<div className={ styles.previewDescription }>{ description }</div>
	</div>
}

export default Home;