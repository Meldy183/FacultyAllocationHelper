"use client"
import styles from "./styles.module.scss";
import React from "react";
import { Accordion, AccordionContent, AccordionItem, AccordionTrigger } from "@/shared/ui/accordion";
import { useGetCourseFiltersQuery } from "@/features/api/slises/filters";
import { Label } from "@/shared/ui/label";
import { Checkbox } from "@/shared/ui/checkbox";
import { useAppDispatch, useAppSelector } from "@/features/store/hooks";
import { toggleFilters, toggleIsAllocated } from "@/features/store/slices/filters/course";
import { FilterItem } from "@/shared/types/api/filters";

const SideBarContent: React.FC = () => {
	const dispatcher = useAppDispatch();
	const filters = useAppSelector(state => state.courseFilters.filters);

	const changeAllocationFilter = () => {
		dispatcher(toggleIsAllocated());
	}

	const isChecked = (groupName: "academic_year" | "semester" | "study_program" | "institute", filterName: string) => {
		return filters[groupName].some(filter => filter.name === filterName);
	}

	const changeFilters = (name: "academic_year" | "semester" | "study_program" | "institute", filterItem: FilterItem) => {
		dispatcher(toggleFilters({
			name,
			items: [filterItem]
		}));
	}

	const { data, error } = useGetCourseFiltersQuery({});

	React.useEffect(() => {
		console.log(data)
	}, [data])

	if (error || !data) return <div>cant load filers</div>;

	return (
		<>
			<div className={ styles.sideBar }>
				<Label>
					<Checkbox checked={ filters.allocaion_not_finished } onCheckedChange={ changeAllocationFilter } />
					<span className={ styles.text }>Allocation not finished</span>
				</Label>
				<div className={ styles.menu }>
					<Accordion type="multiple">
						<AccordionItem className={ styles.accordionItem } value="item-1">
							<AccordionTrigger className={ styles.button }>academic year</AccordionTrigger>
							<AccordionContent>
								{
									data.filters.academic_year.map(filter => (
										<Label key={ filter.name }>
											<Checkbox checked={ isChecked("academic_year", filter.name) } onCheckedChange={ () => changeFilters("academic_year", { name: filter.name, value: filter.id }) } />
											<span className={ styles.text }>{ filter.name }</span>
										</Label>
									))
								}
							</AccordionContent>
						</AccordionItem>
						<AccordionItem className={ styles.accordionItem } value="item-2">
							<AccordionTrigger className={ styles.button }>semester</AccordionTrigger>
							<AccordionContent>
								{
									data.filters.academic_year.map(filter => (
										<Label key={ filter.name }>
											<Checkbox checked={ isChecked("semester", filter.name) } onCheckedChange={ () => changeFilters("semester", { name: filter.name, value: filter.id }) } />
											<span className={ styles.text }>{ filter.name }</span>
										</Label>
									))
								}
							</AccordionContent>
						</AccordionItem>
						<AccordionItem className={ styles.accordionItem } value="item-3">
							<AccordionTrigger className={ styles.button }>study program</AccordionTrigger>
							<AccordionContent>
								{
									data.filters.academic_year.map(filter => (
										<Label key={ filter.name }>
											<Checkbox checked={ isChecked("study_program", filter.name) } onCheckedChange={ () => changeFilters("study_program", { name: filter.name, value: filter.id }) } />
											<span className={ styles.text }>{ filter.name }</span>
										</Label>
									))
								}
							</AccordionContent>
						</AccordionItem>
						<AccordionItem className={ styles.accordionItem } value="item-4">
							<AccordionTrigger className={ styles.button }>institute</AccordionTrigger>
							<AccordionContent>
								{
									data.filters.academic_year.map(filter => (
										<Label key={ filter.name }>
											<Checkbox checked={ isChecked("institute", filter.name) } onCheckedChange={ () => changeFilters("institute", { name: filter.name, value: filter.id }) } />
											<span className={ styles.text }>{ filter.name }</span>
										</Label>
									))
								}
							</AccordionContent>
						</AccordionItem>
					</Accordion>
				</div>
			</div>
		</>
	)
}

export default SideBarContent;