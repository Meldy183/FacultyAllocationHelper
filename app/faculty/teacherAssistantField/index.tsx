import React from "react";
import Link from "next/link";
import styles from "../styles.module.scss";

const TeacherAssistance: React.FC = () => {
  return <Link href={ "/faculty/faculty-member" }>
    <li className={styles.row}>
      <div className={styles.colName}>
        <h2>Name Surname</h2>
        <div>@alias</div>
      </div>
      <div className={styles.colEmail}>n.surname@innopolis.university</div>
      <div className={styles.colInstitute}>Institute</div>
      <div className={styles.colPosition}>Position</div>
    </li>
  </Link>
}

export default TeacherAssistance;