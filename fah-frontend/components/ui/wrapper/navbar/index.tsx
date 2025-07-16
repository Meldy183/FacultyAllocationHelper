"use client"
import React from "react";
import Link from "next/link";
import styles from "./styles.module.scss";
import { usePathname } from "next/navigation";

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
		name: "My facultyProfile",
		path: "/facultyProfile"
	}
]

const NavBar: React.FunctionComponent = () => {
	const pagePath = usePathname();

	const isActiveTab = (path: string) => pagePath.includes(path);

	return <div className={ styles.navbar }>
		<ul className={ styles.ul }>
			{
				routes.map( ({ name, path }) => (
					<li key={ name } className={ `${ styles.li } ${ isActiveTab(path) && styles.active }` }>
						<Link href={ path }>{ name }</Link>
					</li>
				))
			}
		</ul>
	</div>
}

export default NavBar;