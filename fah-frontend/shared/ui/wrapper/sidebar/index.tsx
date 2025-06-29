"use client";
import React from "react";
import { motion } from "framer-motion"; // Removed AnimatePresence since we won't unmount
import styles from "./styles.module.scss";

interface Props {
	children: React.ReactNode;
	hiddenText: string;
}

const sidebarVariants = {
	open: { width: "230px" },
	closed: { width: "50px" },
};

const contentVariants = {
	visible: { opacity: 1, x: 0 },
	hidden: { opacity: 0, x: -20 },
};

const SideBar: React.FC<Props> = ({ children, hiddenText }) => {
	const [visible, setVisible] = React.useState(false);

	return (
		<motion.div
			className={styles.sideBar}
			variants={sidebarVariants}
			animate={visible ? "closed" : "open"}
			transition={{ duration: 0.3 }}
		>
			<div className={styles.header}>
				<div
					onClick={() => setVisible(!visible)}
					className={`${styles.toggleButton} ${visible ? styles.closed : styles.open}`}
				>
					<p></p>
					<p></p>
					<p></p>
				</div>
			</div>
			<div className={ `${ styles.contentWrapper } ${ visible ? styles.active : "" }` }>
				<motion.div
					className={ `${ styles.hiddenText } ${ visible ? styles.active : "" }` }
					variants={contentVariants}
					initial="hidden"
					animate={visible ? "visible" : "hidden"}
					transition={{ duration: 0.2 }}
				>
					{hiddenText}
				</motion.div>
				<motion.div
					className={ `${ styles.content }` }
					variants={contentVariants}
					initial="visible"
					animate={visible ? "hidden" : "visible"}
					transition={{ duration: 0.2 }}
				>
					{children}
				</motion.div>
			</div>
		</motion.div>
	);
};

export default SideBar;