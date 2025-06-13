"use client";
import React from "react";
import styles from "./styles.module.scss";

interface Props {
	children: React.ReactNode;
	hiddenText: string;
}

const SideBar: React.FC<Props> = ({ children, hiddenText }) => {
	const [show, setShow] = React.useState(true);

	return <div className={ styles.sideBar }>
		<div onClick={ () => setShow(!show) } className={ styles.toggleButton }>
			<p></p>
			<p></p>
			<p></p>
		</div>
		<div>
			{
				show ? children  : hiddenText
			}
		</div>
	</div>
}

export default SideBar;