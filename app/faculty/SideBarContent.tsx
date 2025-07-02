"use client"
import Image from "next/image";
import arrowSvg from "@/public/icons/svg/arrow.svg";
import styles from "./styles.module.scss";
import React from "react";
import { Accordion, AccordionContent, AccordionItem, AccordionTrigger } from "@/shared/ui/accordion";

const SideBarContent: React.FC = () => {
	React.useEffect(() => {
		console.log("start")
		return () => console.log("finish")
	}, [])

	return (
		<>
			<div className={ styles.sideBar }>
				<div className={ styles.menu }>
					<Accordion type={ "multiple" }>
						<AccordionItem className={ styles.accordionItem } value={ "item-1" }>
							<AccordionTrigger className={ styles.button }>Institute</AccordionTrigger>
							<AccordionContent>
								<form action="4">
									<label htmlFor="DS">
										<input type="checkbox" id="DS"/>
										<span className={ styles.text }>DS</span>
									</label>
									<label htmlFor="DS/Math">
										<input type="checkbox" id={ "DS/Math" }/>
										<span className={ styles.text }>DS/Math</span>
									</label>
									<label htmlFor="DS/SDE">
										<input type="checkbox" id={ "DS/SDE" }/>
										<span className={ styles.text }>DS/SDE</span>
									</label>
									<label htmlFor="GAMEDEV">
										<input type="checkbox" id={ "GAMEDEV" }/>
										<span className={ styles.text }>GAMEDEV</span>
									</label>
									<label htmlFor="RO">
										<input type="checkbox" id={ "RO" }/>
										<span className={ styles.text }>RO</span>
									</label>
									<label htmlFor="HUM">
										<input type="checkbox" id={ "HUM" }/>
										<span className={ styles.text }>HUM</span>
									</label>
									<label htmlFor="SDE">
										<input type="checkbox" id={ "SDE" }/>
										<span className={ styles.text }>SDE</span>
									</label>
									<label htmlFor="SNE">
										<input type="checkbox" id={ "SNE" }/>
										<span className={ styles.text }>SNE</span>
									</label>
								</form>
							</AccordionContent>
						</AccordionItem>
						<AccordionItem className={ styles.accordionItem } value={ "item-2" }>
							<AccordionTrigger className={ styles.button }>Position</AccordionTrigger>
							<AccordionContent>
								<form action="4">
									<label htmlFor="Prof">
										<input type="checkbox" id="Prof"/>
										<span className={ styles.text }>Professor</span>
									</label>
									<label htmlFor="Docent">
										<input type="checkbox" id={ "Docent" }/>
										<span className={ styles.text }>Docent</span>
									</label>
									<label htmlFor="Senior Instructor">
										<input type="checkbox" id={ "Senior Instructor" }/>
										<span className={ styles.text }>Senior Instructor</span>
									</label>
									<label htmlFor="Instructor">
										<input type="checkbox" id={ "Instructor" }/>
										<span className={ styles.text }>Instructor</span>
									</label>
									<label htmlFor="TA">
										<input type="checkbox" id={ "TA" }/>
										<span className={ styles.text }>TA</span>
									</label>
									<label htmlFor="TA Intern">
										<input type="checkbox" id={ "TA Intern" }/>
										<span className={ styles.text }>TA Intern</span>
									</label>
									<label htmlFor="Visiting">
										<input type="checkbox" id={ "Visiting" }/>
										<span className={ styles.text }>Visiting</span>
									</label>
								</form>
							</AccordionContent>
						</AccordionItem>
					</Accordion>
				</div>
			</div>
		</>
	)
}

export default SideBarContent;