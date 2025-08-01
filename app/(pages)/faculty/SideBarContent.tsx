"use client"
import styles from "./styles.module.scss";
import React from "react";
import { Accordion, AccordionContent, AccordionItem, AccordionTrigger } from "@/shared/ui/accordion";
import { FilterGroup, FilterItem } from "@/shared/types/api/filters";
import { Label } from "@/shared/ui/label";
import { Checkbox } from "@/shared/ui/checkbox";
import { useAppDispatch, useAppSelector } from "@/features/store/hooks";
import { toggleFilter } from "@/features/store/slices/filters/faculty";
import { useGetProfileFiltersQuery } from "@/features/api/slises/filters";

const SideBarContent: React.FC = () => {
	const filters = useAppSelector(state => state.facultyFilters.filters);
	const dispatcher = useAppDispatch();

	const { data, isError } = useGetProfileFiltersQuery({});

	const toggleFilters = (filterGroupName: string, filter: FilterItem) => {
		dispatcher(toggleFilter({
			name: filterGroupName,
			items: [filter]
		}))
	}

	const isChecked = (filterGroupName: string, filter: FilterItem): boolean => {
		return filters.some(
				(filterGroup) =>
					filterGroup.name === filterGroupName &&
					filterGroup.items.some((i) => i.name === filter.name)
			)
	}

	return (
		<>
			<div className={ styles.sideBar }>
				<div className={ styles.menu }>
					<Accordion type={ "multiple" }>
						{
							isError ? <>could not load filters</> :
							data?.map((filterGroup: FilterGroup) => (
								<AccordionItem className={ styles.accordionItem } value={ filterGroup.name } key={ filterGroup.name }>
									<AccordionTrigger className={ styles.button }>{ filterGroup.name }</AccordionTrigger>
									<AccordionContent>
										{
											filterGroup.items.map((filter: FilterItem) =>
												<Label key={ filter.name }>
													<Checkbox checked={ isChecked(filterGroup.name, filter) } onCheckedChange={ () => toggleFilters(filterGroup.name, filter) } />
													<span className={ styles.text }>{ filter.name }</span>
												</Label>
											)
										}
									</AccordionContent>
								</AccordionItem>
							))
						}
					</Accordion>
				</div>
			</div>
		</>
	)
}

export default SideBarContent;