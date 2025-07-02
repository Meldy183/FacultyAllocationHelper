"use client"
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
				<form action="0">
					<label htmlFor="">
						<input type="checkbox" id="mark"/>
						<span className={ styles.text }>Allocation not finished</span>
					</label>
				</form>
				<div className={ styles.menu }>
					<Accordion type="multiple" collapsible>
						<AccordionItem value="item-1">
							<AccordionTrigger className={ styles.button }>Year of study</AccordionTrigger>
							<AccordionContent>
								<form action="1">
									<label htmlFor="1">
										<input type="checkbox" id="1"/>
										<span className={ styles.text }>BS - Year 1</span>
									</label>
									<label htmlFor="2">
										<input type="checkbox" id={ "2" }/>
										<span className={ styles.text }>BS - Year 2</span>
									</label>
									<label htmlFor="3">
										<input type="checkbox" id={ "3" }/>
										<span className={ styles.text }>BS - Year 3</span>
									</label>
									<label htmlFor="4">
										<input type="checkbox" id={ "4" }/>
										<span className={ styles.text }>MS</span>
									</label>
									<label htmlFor="5">
										<input type="checkbox" id={ "5" }/>
										<span className={ styles.text }>PhD</span>
									</label>
								</form>
							</AccordionContent>
						</AccordionItem>
						<AccordionItem value="item-2">
							<AccordionTrigger className={ styles.button }>Semester</AccordionTrigger>
							<AccordionContent>
								<form action="2">
									<label htmlFor="Fall">
										<input type="checkbox" id="Fall"/>
										<span className={ styles.text }>Fall</span>
									</label>
									<label htmlFor="Spring">
										<input type="checkbox" id="Spring"/>
										<span className={ styles.text }>Spring</span>
									</label>
									<label htmlFor="Summer">
										<input type="checkbox" id="Summer"/>
										<span className={ styles.text }>Summer</span>
									</label>

								</form>
							</AccordionContent>
						</AccordionItem>
						<AccordionItem value="item-3">
							<AccordionTrigger className={ styles.button }>Study program</AccordionTrigger>
							<AccordionContent>
								<form action="3">
									<label htmlFor="ISE">
										<input type="checkbox" id="ISE"/>
										<span className={ styles.text }>ISE</span>
									</label>
									<label htmlFor="DSAI">
										<input type="checkbox" id={ "DSAI" }/>
										<span className={ styles.text }>DSAI</span>
									</label>
									<label htmlFor="MFAI">
										<input type="checkbox" id={ "MFAI" }/>
										<span className={ styles.text }>MFAI</span>
									</label>
									<label htmlFor="AI360">
										<input type="checkbox" id={ "AI360" }/>
										<span className={ styles.text }>AI360</span>
									</label>
									<label htmlFor="RO">
										<input type="checkbox" id={ "RO" }/>
										<span className={ styles.text }>RO</span>
									</label>
									<label htmlFor="SE">
										<input type="checkbox" id={ "SE" }/>
										<span className={ styles.text }>SE</span>
									</label>
									<label htmlFor="RO">
										<input type="checkbox" id={ "RO" }/>
										<span className={ styles.text }>RO</span>
									</label>
									<label htmlFor="DS">
										<input type="checkbox" id={ "DS" }/>
										<span className={ styles.text }>DS</span>
									</label>
									<label htmlFor="TE">
										<input type="checkbox" id={ "TE" }/>
										<span className={ styles.text }>TE</span>
									</label>
									<label htmlFor="SNE">
										<input type="checkbox" id={ "SNE" }/>
										<span className={ styles.text }>SNE</span>
									</label>
								</form>
							</AccordionContent>
						</AccordionItem>
						<AccordionItem value="item-4">
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
					</Accordion>
				</div>
			</div>
		</>
	)
}

export default SideBarContent;