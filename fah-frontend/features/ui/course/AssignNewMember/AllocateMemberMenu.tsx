import React from "react";
import styles from "@/features/ui/course/AssignNewMember/styles.module.scss";
import { Button } from "@/shared/ui/button";
import { Input } from "@/shared/ui/input"
import { Accordion, AccordionContent, AccordionItem, AccordionTrigger } from "@/shared/ui/accordion";
import CreateNewFacultyForm from "@/entities/faculty/CreateFacultyForm/ui";

interface IProps {
	position: string
	foldMenu: () => void
}

const AllocateMemberMenu: React.FC<IProps> = ({ position, foldMenu }) => {
  return (
	  <div className={ styles.allocateExisting }>
		  <div className={ styles.head }>
			  <div className={ styles.title }>{ position }</div>
		  </div>

		  <div className={ styles.searchBarContainer }>
			  <Input type="text" className={ styles.searchBar } placeholder="Search..."/>
		  </div>

		  <Accordion type={ "single" } collapsible>
			  <AccordionItem value={ "create-user" }>
				  <AccordionTrigger>
					  <Button variant={ "secondary" } className={ `${ styles.searchButton } w-max` } asChild><span>Create new member</span></Button>
				  </AccordionTrigger>
				  <AccordionContent>
					  <CreateNewFacultyForm onSubmit={ (data) => { console.log(data) } } />
				  </AccordionContent>
			  </AccordionItem>
			  <AccordionItem value={ "add-filters" }>
				  <AccordionTrigger>
					  <Button className={ styles.searchButton } asChild><span>add searching filters</span></Button>
				  </AccordionTrigger>
				  <AccordionContent>
					  something filters
				  </AccordionContent>
			  </AccordionItem>
		  </Accordion>

		  <div className={ styles.membersContainer }>
			  <ul className={ styles.column1 }>
				  <li>
					  <div>Member Name</div>
					  <div>@alias</div>
				  </li>
				  <li>
					  <div>Member Name</div>
					  <div>@alias</div>
				  </li>
				  <li>
					  <div>Member Name</div>
					  <div>@alias</div>
				  </li>
				  <li>
					  <div>Member Name</div>
					  <div>@alias</div>
				  </li>
				  <li>
					  <div>Member Name</div>
					  <div>@alias</div>
				  </li>
				  <li>
					  <div>Member Name</div>
					  <div>@alias</div>
				  </li>
			  </ul>
			  <ul className={ styles.column2 }>
				  <li>
					  <Button className={ styles.searchButton }>Allocate</Button>
				  </li>
				  <li>
					  <Button className={ styles.searchButton }>Allocate</Button>
				  </li>
				  <li>
					  <Button className={ styles.searchButton }>Allocate</Button>
				  </li>
				  <li>
					  <Button className={ styles.searchButton }>Allocate</Button>
				  </li>
				  <li>
					  <Button className={ styles.searchButton }>Allocate</Button>
				  </li>
				  <li>
					  <Button className={ styles.searchButton }>Allocate</Button>
				  </li>
			  </ul>
		  </div>

		  <Button onClick={ foldMenu }>fold menu</Button>
	  </div>
  )
}

export default AllocateMemberMenu;