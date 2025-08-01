import React from "react";
import { Button } from "@/shared/ui/button";
import Wrapper from "@/shared/ui/wrapper";
import styles from "./styles.module.scss";
import Link from "next/link";
import LoginForm from "@/app/(pages)/auth/login/LoginForm";

const AuthPage: React.FC = () => {
	return <Wrapper hasNavBar={ false }>
		<div className={ styles.menu }>
			<div className={ styles.top }>
				<Button className={ styles.button } variant={ "strictWhite" }>
					<Link href={ "/start" }>go back</Link>
				</Button>
			</div>
			<LoginForm />
		</div>
	</Wrapper>
}

export default AuthPage;