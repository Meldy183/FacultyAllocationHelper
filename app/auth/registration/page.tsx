import React from "react";
import { Button } from "@/components/ui/button";
import Wrapper from "@/components/ui/wrapper";
import RegistrationForm from "@/app/auth/registration/RegistrationForm";
import Link from "next/link";
import styles from "./styles.module.scss";

const RegistrationPage: React.FC = () => {
	return <Wrapper hasNavBar={ false }>
		<div className={ styles.menu }>
			<div className={ styles.top }>
				<Button className={ styles.button } variant={ "strictWhite" }>
					<Link href={ "/start" }>go back</Link>
				</Button>
			</div>
			<RegistrationForm />
		</div>
	</Wrapper>
}

export default RegistrationPage