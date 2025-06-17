import styles from "./styles.module.scss";
import React, { ReactNode } from "react";
import NavBar from "./navbar";

interface Props {
	children: ReactNode;
	hasNavBar?: boolean
}

const Wrapper: React.FunctionComponent<Props> = ({ children, hasNavBar = true }: Props) => {
	return (
		<div className={ styles.wrapper }>
			<header>
				{ hasNavBar && <NavBar /> }
			</header>
			<main className={ styles.container }>
				{ children }
			</main>
			<footer />
		</div>
	)
}

export default Wrapper;