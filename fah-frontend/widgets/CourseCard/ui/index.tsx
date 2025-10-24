"use client"

import React from "react";
import TAElement from "@/features/ui/course/FacultyMemberBlock";
import AssignNewMember from "@/features/ui/course/AssignNewMember";
import { CourseInformation } from "@/entities/CourseInformation";
import { CourseType } from "@/shared/types/ui/courses";
import {
    Dialog,
    DialogContent, DialogDescription, DialogHeader,
    DialogTitle,
    DialogTrigger,
} from "@/shared/ui/dialog";
import { VisuallyHidden } from "@radix-ui/react-visually-hidden";
import styles from "./styles.module.scss";

type Props = CourseType & {}

export const CourseCard: React.FC<Props> = (props) => {
    return <div className={ styles.card }>
        <div className={ styles.header }>
            <div className={ styles.information }>
                <div className={ styles.courseName }>
                    <Dialog>
                        <DialogTrigger className={ "cursor-pointer" }>
                            <div className={ styles.name }>{ props.brief_name }</div>
                        </DialogTrigger>
                        <DialogContent className={ styles.dialogMenu }>
                            <VisuallyHidden>
                                <DialogHeader>
                                    <DialogTitle/>
                                    <DialogDescription/>
                                </DialogHeader>
                            </VisuallyHidden>
                            <CourseInformation { ...props } />
                        </DialogContent>
                    </Dialog>
                    {/*<Dialog>*/ }
                    {/*    <DialogTrigger>*/ }
                    {/*        <div className={ `${ styles.icon } cursor-pointer` }>*/ }
                    {/*            <Image src={ settingsIcon } alt={ "settings" } className={ styles.icon }/>*/ }
                    {/*        </div>*/ }
                    {/*    </DialogTrigger>*/ }
                    {/*    <DialogContent className={ styles.dialogMenu }>*/ }
                    {/*        <VisuallyHidden>*/ }
                    {/*            <DialogHeader>*/ }
                    {/*                <DialogTitle/>*/ }
                    {/*                <DialogDescription/>*/ }
                    {/*            </DialogHeader>*/ }
                    {/*        </VisuallyHidden>*/ }
                    {/*        <ManageCourseDialogMenu />*/ }
                    {/*    </DialogContent>*/ }
                    {/*</Dialog>*/ }
                </div>
                <ul>
                    {/*<li>Study year: { props.study_year }</li>*/ }
                    <li>Semester: { props.semester_name }</li>
                    <li>Study program: {
                        props.study_program_names?.map(program => <span key={ program }>{ program } </span>)
                    }</li>
                    <li>Institute: { props.responsible_institute_name }</li>
                </ul>
            </div>
            <div className={ styles.instructors }>
                <div className={ styles.instructor }>
                    <div className={ styles.title }>PI</div>
                    {
                        props.pi?.allocation_status === "assigned" ? <TAElement { ...props.pi } /> :
                            <span>not allocated</span>
                    }
                </div>
                <div className={ styles.instructor }>
                    <div className={ styles.title }>Tutor</div>
                    {
                        props.ti?.allocation_status === "assigned" ? <TAElement { ...props.ti } /> :
                            <span>not allocated</span>
                    }
                </div>
            </div>
        </div>
        <div className={ styles.body }>
            <div className={ styles.title }>Teaching assistants</div>
            <div className={ styles.assistance }>
                <AssignNewMember/>
                { props.tas?.map(ta => <TAElement { ...ta } key={ ta.profile_data.profile_id }/>) }
            </div>
        </div>
    </div>
}