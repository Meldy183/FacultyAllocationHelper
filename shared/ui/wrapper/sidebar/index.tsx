"use client";
import React from "react";
import { motion, AnimatePresence } from "framer-motion";
import clsx from "clsx";
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

export default function SideBar({ children, hiddenText }: Props) {
	const [isCollapsed, setIsCollapsed] = React.useState(false);
	const toggle = () => setIsCollapsed((prev) => !prev);

	return (
		<motion.div
			className={styles.sideBar}
			variants={sidebarVariants}
			animate={isCollapsed ? "closed" : "open"}
			transition={{ duration: 0.3 }}
		>
			<div className={`${ styles.header } ${ isCollapsed ? "" : styles.open }`}>
				<button
					className={clsx(styles.toggleButton, {
						[styles.closed]: isCollapsed,
						[styles.open]: !isCollapsed,
					})}
					onClick={toggle}
					aria-label={isCollapsed ? "Open sidebar" : "Close sidebar"}
				>
					<span />
					<span />
					<span />
				</button>
			</div>

			<div className={styles.contentWrapper}>
				<AnimatePresence initial={false} mode="wait">
					{!isCollapsed ? (
						<motion.div
							key="content"
							className={styles.content}
							variants={contentVariants}
							initial="hidden"
							animate="visible"
							exit="hidden"
							transition={{ duration: 0.2 }}
						>
							{children}
						</motion.div>
					) : (
						<motion.div
							key="hiddenText"
							className={styles.hiddenText}
							variants={contentVariants}
							initial="hidden"
							animate="visible"
							exit="hidden"
							transition={{ duration: 0.2 }}
						>
							<span>{ hiddenText }</span>
						</motion.div>
					)}
				</AnimatePresence>
			</div>
		</motion.div>
	);
}
