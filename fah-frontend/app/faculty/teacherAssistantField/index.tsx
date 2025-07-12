import React from "react";
import Link from "next/link";
import { UserDataInterface } from "@/shared/types/apiTypes/members";
import styles from "../styles.module.scss";

const TeacherAssistance: React.FC<UserDataInterface> = (props: UserDataInterface) => {
  return <Link href={ `/faculty/${ props.profile_id }` }>
    <li className={styles.row}>
      <div className={styles.colName}>
        <h2>{ props.nameEng }</h2>
        <div>{ props.alias }</div>
      </div>
      <div className={styles.colEmail}>{ props.email }</div>
      <div className={styles.colInstitute}>{ props.institute }</div>
      <div className={styles.colPosition}>{ props.position }</div>
    </li>
  </Link>
}

export default TeacherAssistance;