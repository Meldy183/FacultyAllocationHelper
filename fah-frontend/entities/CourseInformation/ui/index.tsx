import React from "react";
import Link from "next/link";
import { Accordion, AccordionContent, AccordionItem, AccordionTrigger } from "@/shared/ui/accordion";
import type { CourseType } from "@/shared/types";
import styles from "./styles.module.scss";

export const CourseInformation: React.FC<CourseType> = (props) => {
    return <div className={ styles.container }>
        <div className={ styles.card }>
            <div className={ styles.header }>
                <div className={ styles.userInfo }>
                    <div>
                        <h1 className={ styles.name }>{ props.brief_name }</h1>
                        <p className={ styles.subName }>{ props.official_name }</p>
                    </div>
                </div>
            </div>
            <Accordion className={ styles.section } type="single" collapsible>
                <AccordionItem value="item-1">
                    <AccordionTrigger
                        className={ `${ styles.title } cursor-pointer` }
                    >Course Information</AccordionTrigger>
                    <AccordionContent className={ styles.grid }>
                        <div>Responsible institute: { props.responsible_institute_name }</div>
                        <div>Program: program code</div>
                        <div>Track: { props.track_names.map((track, i) => <span key={ i }>{ track }</span>) }</div>
                        <div>Mode: { props.mode }</div>
                        <div>Form: { props.form }</div>
                        <div>Semester: { props.semester_name }</div>
                    </AccordionContent>
                </AccordionItem>
            </Accordion>
            <div className={ styles.section }>
                <h2 className={ styles.title }>УП</h2>
                <div className={ styles.grid }>
                    <div>Lecture hours (per course): N</div>
                    <div>Practical class hours: N</div>
                </div>
            </div>
        </div>
        <div className={ styles.card }>
            <div className={ styles.sectionCardWhite }>
                <h2 className={ styles.name }>Instructors on this course</h2>
                <h3 className={ styles.subName }>Primary instructor</h3>
                <ul className={ styles.list }>
                    <TeacherAssistance/>
                </ul>
                <h3 className={ styles.subName }>Tutor instructor</h3>
                <ul className={ styles.list }>
                    <TeacherAssistance/>
                </ul>

                <h3 className={ styles.subName }>Teaching assistants</h3>
                <div className={ styles.assistance }>
                    <ul className={ styles.list }>
                        <li className={ styles.tableHeader }>
                            <div className={ styles.colName }>Name, alias</div>
                            <div className={ styles.colEmail }>Email</div>
                        </li>
                        <TeacherAssistance/>
                        <TeacherAssistance/>
                        <TeacherAssistance/>
                        <TeacherAssistance/>
                        <TeacherAssistance/>
                        <TeacherAssistance/>
                        <TeacherAssistance/>
                        <TeacherAssistance/>
                        <TeacherAssistance/>
                        <TeacherAssistance/>
                    </ul>
                </div>
            </div>
        </div>
    </div>
}

const TeacherAssistance: React.FC = () => {
    return <Link href={ "/faculty/0" }>
        <li className={ styles.row }>
            <div className={ styles.colName }>
                <h2>Name Surname</h2>
                <div>@alias</div>
            </div>
            <div className={ styles.colEmail }>n.surname@innopolis.university</div>
        </li>
    </Link>
}