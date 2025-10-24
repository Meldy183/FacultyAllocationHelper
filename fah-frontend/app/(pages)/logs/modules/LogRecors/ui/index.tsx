import React from "react";
import styles from "./styles.module.scss";

export const LogRecord: React.FC = () => {
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