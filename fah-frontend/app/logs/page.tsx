"use client";

import React from "react";
import Wrapper from "@/shared/ui/wrapper";
import styles from "./styles.module.scss";

const LogsPage: React.FC = () => {
	return <Wrapper>
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


export default LogsPage;