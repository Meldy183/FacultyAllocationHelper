"use client"
import React from "react";
import Link from "next/link";
import styles from "./styles.module.scss";
import { usePathname } from "next/navigation";
import userIcon from "@/public/icons/faculty/faculty-member/faculty-member-icon.svg"
import Image from "next/image";


//переместить
interface Routes {
	path: string;
	name: string;
}
const routes: Routes[] = [
	{
		name: "Home",
		path: "/dashboard"
	},
	{
		name: "Courses",
		path: "/courses"
	},
	{
		name: "Faculty",
		path: "/faculty"
	},
	{
		name: "Logs",
		path: "/logs"
	}
	// {
	// 	name: "My profile",
	// 	path: "/profile"
	// }
]

const NavBar: React.FunctionComponent = () => {
	const pagePath = usePathname();

	const isActiveTab = (path: string) => pagePath.includes(path);

	return (
		<div className={ styles.container }>
			<div className={styles.navbar}>
				<ul className={styles.ul}>
					{
					routes.map(({ name, path }) => (
						<li key={name} className={`${styles.li} ${isActiveTab(path) && styles.active}`}>
						<Link href={path}>{name}</Link>
						</li>
					))
					}
				</ul>
				<div className={ styles.profileWrapper }>
					<span className={styles.profileName}>Name Surname</span>
					<Image src={userIcon} alt={"user icon"} className={styles.avatar} />
				</div>
			</div>
		</div>
	)
}

export default NavBar;