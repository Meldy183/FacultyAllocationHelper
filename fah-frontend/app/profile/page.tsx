"use client";

import React from "react";
import Image from "next/image";
import { Button } from "@/components/ui/button";
import Wrapper from "@/components/ui/wrapper";
import crossSvg from "@/public/icons/svg/cross.svg";
import { useLogoutMutation } from "@/features/api/slises/authSlice";
import { useRouter } from 'next/navigation';
import styles from "./styles.module.scss";

const RegistrationPage: React.FC = () => {
	const [logout] = useLogoutMutation();
	const router = useRouter();

	const handleClick = async () => {
		const result = await logout();
		router.push("/start");
	}
 
	return <Wrapper>
	<Button onClick={ () => handleClick() } className={ styles.button }>log out</Button>
	</Wrapper>
}

export default RegistrationPage