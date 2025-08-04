"use client";

import React from "react";
import Wrapper from "@/shared/ui/wrapper";
import { LogRecord } from "./modules/LogRecors";
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


export default LogsPage;