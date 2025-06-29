import React from "react";
import Image from "next/image";
import { Button } from "@/shared/ui/button";
import Wrapper from "@/shared/ui/wrapper";
import crossSvg from "@/public/icons/svg/cross.svg";
import styles from "./styles.module.scss";

const ProfilePage: React.FC = () => {
	return <Wrapper>
		<div className={ styles.menu }>
			<div className={ styles.top }>
				<Image className={ styles.image } src={ crossSvg } alt={ "close page" } />
			</div>
			<div className={ styles.content }>
				<div className={ styles.title }>Registration/Login</div>
				<div className={ styles.fieldBlock }>
					<div className={ styles.fieldDescription }>E-mail</div>
					<input type="text" placeholder={ "your email" } className={  styles.input }/>
				</div>
				<div className={ styles.fieldBlock }>
					<div className={ styles.fieldDescription }>Password</div>
					<input type="password" placeholder={ "password" } className={  styles.input }/>
				</div>
				<Button className={ styles.button }>Submit</Button>
			</div>
		</div>
	</Wrapper>
}

export default ProfilePage;