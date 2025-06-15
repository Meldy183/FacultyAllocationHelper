"use client"
import React from "react";
import Image from "next/image";
import { useForm, SubmitHandler } from "react-hook-form";
import { Form, FormControl, FormField, FormItem, FormLabel } from "@/components/ui/form";
import { Input } from "@/components/ui/input";
import { z } from "zod";
import { zodResolver } from "@hookform/resolvers/zod";

import { registerResolver } from "@/types/resolvers/auth";
import { Button } from "@/components/ui/button";
import Wrapper from "@/components/ui/wrapper";
import crossSvg from "@/public/icons/svg/cross.svg";
import styles from "./styles.module.scss";
import Link from "next/link";

type RegisterInput = z.infer<typeof registerResolver>

const RegistrationPage: React.FC = () => {
	const form = useForm<RegisterInput>({
		resolver: zodResolver(registerResolver),
		defaultValues: {
			email: "",
			password: ""
		}
	});

	const submitHandler: SubmitHandler<RegisterInput> = (data) => {
		console.log(data);
	}

	return <Wrapper>
		<div className={ styles.menu }>
			<div className={ styles.top }>
				<Button className={ styles.button } variant={ "strictWhite" }>
					<Link href={ "/start" }>go back</Link>
				</Button>
			</div>
			<Form { ...form }>
				<form onSubmit={ form.handleSubmit(submitHandler) } className={ styles.content }>
					<div className={ styles.title }>Login</div>
					<div className={ styles.fieldBlock }>
						<FormField name={ "email" }
						           render={ ({ field }) =>
							           <CustomField fieldName={ "email" } type={ "text" } title={ "E-mail" } field={ field }/> }
						/>
						<FormField name={ "password" }
						           render={ ({ field }) =>
							           <CustomField fieldName={ "password" } type={ "password" } title={ "password" } field={ field }/> }
						/>
					</div>
					<Button className={ styles.button }>Submit</Button>
				</form>
			</Form>
		</div>
	</Wrapper>
}

interface fieldProps {
	field: any,
	fieldName: string,
	title: string
	type?: string
	customClassName?: string
}

const CustomField: React.FC<fieldProps> = ({ field, fieldName, title, type = "text", customClassName = "" }) => {
	const autocompleteValue = fieldName === "password" ? "current-password" : "off";

	return (
		<>
			<FormItem>
				<FormLabel><div className={ styles.fieldDescription }>{ title }</div></FormLabel>
				<FormControl>
					<Input
						name={ fieldName }
						className={ `${ styles.input } ${ customClassName }` }
						type={ type }
						placeholder={ title }
						autoComplete={ autocompleteValue }
						{ ...field }
					/>
				</FormControl>
			</FormItem>
		</>
	)
}

export default RegistrationPage