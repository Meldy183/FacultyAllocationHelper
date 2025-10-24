"use client"
import React from "react";
import { useGetCourseFiltersQuery } from "../api";
import { Accordion, AccordionContent, AccordionItem, AccordionTrigger } from "@/shared/ui/accordion";
import { Label } from "@/shared/ui/label";
import { Checkbox } from "@/shared/ui/checkbox";
import { useAppDispatch, useAppSelector } from "@/features/store/hooks";
import { filtersEnum, toggleFilters, toggleIsAllocated } from "../models";
import { FilterItem } from "@/shared/types";
import styles from "./styles.module.scss";

const filterGroups: Array<{ key: filtersEnum; title: string }> = [
    { key: filtersEnum.academic_year, title: "Academic Year" },
    { key: filtersEnum.semester, title: "Semester" },
    { key: filtersEnum.study_program, title: "Study Program" },
    { key: filtersEnum.institute, title: "Institute" },
]

export const CourseFilters: React.FC = () => {
    const dispatcher = useAppDispatch();
    const filters = useAppSelector(state => state.courseFilters);

    const changeAllocationFilter = () =>
        dispatcher(toggleIsAllocated());

    const isChecked = (groupName: filtersEnum, filterName: string) =>
        filters[groupName].some(f => f.name === filterName);

    const changeFilters = (name: filtersEnum, filterItem: FilterItem) =>
        dispatcher(toggleFilters({ name, items: [filterItem] }));

    const { data, error } = useGetCourseFiltersQuery({});

    if (error || !data) return <div>cant load filers</div>;

    return (

        <div className={ styles.sideBar }>
            <Label>
                <Checkbox checked={ filters.allocaion_not_finished } onCheckedChange={ changeAllocationFilter }/>
                <span className={ styles.text }>Allocation not finished</span>
            </Label>
            <div className={ styles.menu }>
                <Accordion type="multiple">
                    {
                        filterGroups.map(({ key, title }) => (
                            <AccordionItem className={ styles.accordionItem } key={ key } value={ key }>
                                <AccordionTrigger className={ styles.button }>{ title }</AccordionTrigger>
                                <AccordionContent>
                                    {
                                        data[key].map(filter => (
                                            <Label key={ filter.name }>
                                                <Checkbox
                                                    checked={ isChecked(key, filter.name) }
                                                    onCheckedChange={ () => changeFilters(key, {
                                                        name: filter.name,
                                                        value: filter.id
                                                    }) }/>
                                                <span className={ styles.text }>{ filter.name }</span>
                                            </Label>
                                        ))
                                    }
                                </AccordionContent>
                            </AccordionItem>)
                        )
                    }
                </Accordion>
            </div>
        </div>

    )
}