"use client"
import React from "react";
import Link from "next/link";
import { Button } from "@/components/ui/button";
import { routesAuth } from "@/configs/routes";
import styles from "./styles.module.scss";
import { GetServerSideProps } from "next";

const StartPage: React.FC = () => {
	return <main className={ styles.main }>
		<div className={ styles.content }>
			<h3 className={ styles.title }>Faculty allocation helper</h3>
			<div className={ styles.description }>To get access to the service, please, login through your Innopolis
				University e-mail.
			</div>
			<div className={ styles.buttons }>
				{ routesAuth.map(({ routeName, routePath }) =>
					<Link key={ routePath } href={ routePath }>
						<Button variant={ "strictWhite" } className={ styles.button }>
							{ routeName }
						</Button>
          </Link>)
				}
			</div>
		</div>
	</main>;
};

export default StartPage;