"use client"
import styles from "./styles.module.scss";
import React from "react";
import { Accordion, AccordionContent, AccordionItem, AccordionTrigger } from "@/shared/ui/accordion";
import { useGetFiltersQuery } from "@/features/api/slises/profile";
import { FilterGroup, FilterItem } from "@/shared/types/apiTypes/filters";
import { Label } from "@/shared/ui/label";
import { Checkbox } from "@/shared/ui/checkbox";

const SideBarContent: React.FC = () => {
	const { data, error } = useGetFiltersQuery({});

	React.useEffect(() => {
		console.log(data)
	}, [data]);

	return (
		<>
			<div className={ styles.sideBar }>
				<div className={ styles.menu }>
					<Accordion type={ "multiple" }>
						{
							data?.map((filterGroup: FilterGroup) => (
								<AccordionItem className={ styles.accordionItem } value={ filterGroup.name } key={ filterGroup.name }>
									<AccordionTrigger className={ styles.button }>{ filterGroup.name }</AccordionTrigger>
									<AccordionContent>
										{
											filterGroup.items.map((filter: FilterItem) =>
												<Label key={ filter.name }>
													<Checkbox />
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