import React from "react";
import styles from "@/features/ui/course/AssignNewMember/styles.module.scss";
import { Button } from "@/shared/ui/button";
import { Input } from "@/shared/ui/input"


export default function SearchBar() {
  return (
	<div className={styles.allocateExisting}>
		<div className={styles.head}>
			<div className={ styles.title }>Existing faculty</div>
		</div>
		<div className={styles.searchBarContainer}>
			<Input type="text" className={styles.searchBar} placeholder="Search..." />
			<Button className={styles.searchButton}>Search</Button>
		</div>
		<div className={styles.membersContainer}>
			<ul className={styles.column1}>
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
			<ul className={styles.column2}>
				<li>
					<Button className={styles.searchButton}>Allocate</Button>
				</li>
				<li>
					<Button className={styles.searchButton}>Allocate</Button>
				</li>
				<li>
					<Button className={styles.searchButton}>Allocate</Button>
				</li>
				<li>
					<Button className={styles.searchButton}>Allocate</Button>
				</li>
				<li>
					<Button className={styles.searchButton}>Allocate</Button>
				</li>
				<li>
					<Button className={styles.searchButton}>Allocate</Button>
				</li>
			</ul>
		</div>
	</div>
  )
}


const AllocateExistingMember: React.FC = () => {
	return <div className={ styles.createNewMember }>
		Вот тут надо будет сверстать поиск TA (добавить верстку фильтров и поисковую строку)
	</div>
}


// export default AllocateExistingMember;