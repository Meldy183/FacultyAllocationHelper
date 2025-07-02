"use client";

import React from "react";
import { Button } from "@/shared/ui/button";
import Wrapper from "@/shared/ui/wrapper";
import SideBar from "@/shared/ui/wrapper/sidebar";
import SideBarContent from "@/app/faculty/SideBarContent";
import styles from "./styles.module.scss";
import Link from "next/link";

const AssistantsPage: React.FC = () => {
	return <Wrapper>
			<div className={ styles.assistance }>
				<ul className={styles.list}>
					<li className={styles.header}>
						<div className={styles.colName}>Name, alias</div>
						<div className={styles.colInstitute}>Date, time</div>
						<div className={styles.colEmail}>Action</div>
					</li>
					<LogRecord/>
					<LogRecord/>					
					<LogRecord/>
					<LogRecord/>
					<LogRecord/>					
					<LogRecord/>
					<LogRecord/>
					<LogRecord/>
					<LogRecord/>
					<LogRecord/>
					</ul>
			</div>
	</Wrapper>
}


const LogRecord: React.FC = () => {
	return (
		<li className={styles.row}>
			<div className={styles.colName}>
                <h2>Name Surname</h2>
                <div>@alias</div>
			</div>
			<div className={styles.colInstitute}>
                <h4>02.07.2025</h4>
                <h4>17:35</h4>
            </div>
			<div className={styles.colEmail}>Did something</div>
		</li>
	)
}


export default AssistantsPage;