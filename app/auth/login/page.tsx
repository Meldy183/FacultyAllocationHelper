import React from "react";
import { Button } from "@/components/ui/button";
import Wrapper from "@/components/ui/wrapper";
import styles from "./styles.module.scss";
import Link from "next/link";
import CustomForm from "@/app/auth/login/CustomForm";

const AuthPage: React.FC = () => {
	return <Wrapper hasNavBar={ false }>
		<div className={ styles.menu }>
			<div className={ styles.top }>
				<Button className={ styles.button } variant={ "strictWhite" }>
					<Link href={ "/start" }>go back</Link>
				</Button>
			</div>
			<CustomForm />
		</div>
	</Wrapper>
}

export default AuthPage;