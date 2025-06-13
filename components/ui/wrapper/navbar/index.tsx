import React from "react";
import Link from "next/link";
import styles from "./styles.module.scss";

//переместить
interface Routes {
	path: string;
	name: string;
}
const routes: Routes[] = [
	{
		name: "Home",
		path: "/home"
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
		name: "My profile",
		path: "/profile"
	}
]

const NavBar: React.FunctionComponent = () => {
	return <div className={ styles.navbar }>
		<ul className={ styles.ul }>
			{
				routes.map( ({ name, path }) => (
					<li key={ name } className={ styles.li }>
						<Link href={ path }>{ name }</Link>
					</li>
				))
			}
		</ul>
	</div>
}

export default NavBar;