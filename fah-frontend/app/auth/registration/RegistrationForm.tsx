"use client";
import React from "react";
import { SubmitHandler, useForm } from "react-hook-form";
import { useRegisterMutation } from "@/features/api/slises/authSlice";
import { Form, FormField } from "@/components/ui/form";
import { Button } from "@/components/ui/button";
import { registerResolver } from "@/types/resolvers/auth";
import CustomField from "@/app/auth/registration/CustomField";
import { z } from "zod";
import { zodResolver } from "@hookform/resolvers/zod";
import styles from "./styles.module.scss";
import { handleErrorForm } from "@/hooks/hadleErrorForm";
import { useRouter } from "next/navigation";
import { dashboardRoute } from "@/configs/routes";

type RegisterInput = z.infer<typeof registerResolver>

const RegistrationForm: React.FC = () => {
	const router = useRouter();

	const [register] = useRegisterMutation();

	const form = useForm<RegisterInput>({
		resolver: zodResolver(registerResolver),
		defaultValues: {
			email: "",
			password: "",
			passwordAgain: ""
		}
	});

	const submitHandler: SubmitHandler<RegisterInput> = async (formData) => {
		try {
			if (formData.password !== formData.passwordAgain) {
				form.setError("root", { message: "passwords must be same" });
				return;
			}

			const { data, error } = await register(formData).unwrap();
			if (error) throw error;
			router.push(dashboardRoute.routePath)
		} catch (e: Response) {
			handleErrorForm<RegisterInput>(e, form.setError);
		}
	}

	return <Form { ...form }>
		<form onSubmit={ form.handleSubmit(submitHandler) } className={ styles.content }>
			<div className={ styles.title }>Registration</div>
			<div className={ styles.fieldBlock }>
				<FormField name={ "email" }
				           render={ ({ field }) =>
					           (
											 <>
												 <CustomField fieldName={ "email" } type={ "text" } title={ "E-mail" } field={ field }/>
												 form.formState.errors.email && (
												 <p className="text-red-500 text-sm mt-1">
													 {form.formState.errors.email.message}
												 </p>
												 )}
											 </>
					           ) }
				/>
				<FormField name={ "password" }
				           render={ ({ field }) =>
					           (
											 <>
												 <CustomField fieldName={ "password" } type={ "password" } title={ "password" } field={ field }/>
												 {form.formState.errors.password && (
													 <p className="text-red-500 text-sm mt-1">
														 {form.formState.errors.password.message}
													 </p>
												 )}
											 </>
					           ) }
				/>
				<FormField name={ "passwordAgain" }
				           render={ ({ field }) =>
					           <CustomField fieldName={ "passwordAgain" } type={ "password" } title={ "write you password again" } field={ field }/> }
				/>
			</div>
			<Button className={ styles.button }>Submit</Button>
			{form.formState.errors.root && (
				<p className="text-red-500 text-sm">{form.formState.errors.root.message}</p>
			)}
		</form>
	</Form>
}

export default RegistrationForm;