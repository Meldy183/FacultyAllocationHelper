"use client"
import React from "react";
import { SubmitHandler, useForm } from "react-hook-form";
import { z } from "zod";
import { zodResolver } from "@hookform/resolvers/zod";
import { authResolver } from "@/shared/types/resolvers/auth";
import { Form, FormField } from "@/shared/ui/form";
import CustomField from "@/app/auth/login/CustomField";
import { Button } from "@/shared/ui/button";
import styles from "./styles.module.scss";
import { useLoginMutation } from "@/features/api/slises/authSlice";
import { handleErrorForm } from "@/shared/hooks/hadleErrorForm";
import { useRouter } from "next/navigation";
import { dashboardRoute } from "@/shared/configs/routes";
import { API_PATH } from "@/shared/configs/constants";

type LoginInput = z.infer<typeof authResolver>;

const LoginForm: React.FC = () => {
	const router = useRouter();

	console.log(process.env.NEXT_PUBLIC_BASE_API);
	console.log(API_PATH);

	const [login] = useLoginMutation();

	const form = useForm<LoginInput>({
		resolver: zodResolver(authResolver),
		defaultValues: {
			email: "",
			password: ""
		}
	});

	const submitHandler: SubmitHandler<LoginInput> = async (formData) => {
		try {
			const { data, error } = await login(formData);
			console.log(data, error);
			if (error) throw error;
			router.push(dashboardRoute.routePath);
		} catch (e: any) {
			handleErrorForm<LoginInput>(e, form.setError);
		}
	}

	return (
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
				{form.formState.errors.root && (
					<p className="text-red-500 text-sm">{ form.formState.errors.root.message }</p>
				)}
			</form>
		</Form>
	)
}

export default LoginForm;