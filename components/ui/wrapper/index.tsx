import styles from "./styles.module.scss";
import React, { ReactNode } from "react";
import NavBar from "./navbar";

interface Props {
	children: ReactNode;
}

const Wrapper: React.FunctionComponent<Props> = ({ children }: Props) => {
	return (
		<div className={ styles.wrapper }>
			<header>
				<NavBar />
			</header>
			<main className={ styles.container }>
				{ children }
			</main>
			<footer />
		</div>
	)
}

export default Wrapper;