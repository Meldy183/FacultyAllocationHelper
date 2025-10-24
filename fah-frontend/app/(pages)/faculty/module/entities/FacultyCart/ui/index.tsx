import React from "react";
import Link from "next/link";
import { GetSimpleUserDataInterface } from "@/shared/types/ui/faculties";
import styles from "./styles.module.scss";

export const TeacherAssistance: React.FC<GetSimpleUserDataInterface> = (props) => {
    return <Link prefetch={ false } href={ `/faculty/${ props.profile_id }` }>
        <li className={ styles.row }>
            <div className={ styles.colName }>
                <h2>{ props.name_eng }</h2>
                <div>{ props.alias }</div>
            </div>
            <div className={ styles.colEmail }>{ props.email }</div>
            <div className={ styles.colInstitute }>
                { props.institute_names?.map((name) => <span key={ name }> { name } </span>) }
            </div>
            <div className={ styles.colPosition }>{ props.position_name }</div>
        </li>
    </Link>
}