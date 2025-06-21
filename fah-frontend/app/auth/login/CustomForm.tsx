import React from "react";
import { SubmitHandler, useForm } from "react-hook-form";
import { z } from "zod";
import { zodResolver } from "@hookform/resolvers/zod";
import { authResolver } from "@/types/resolvers/auth";
import { Form, FormField } from "@/components/ui/form";
import CustomField from "@/app/auth/login/CustomField";
import { Button } from "@/components/ui/button";
import styles from "./styles.module.scss";
import { useLoginMutation } from "@/features/api/slises/authSlice";
import { handleErrorForm } from "@/hooks/hadleErrorForm";
import { useRouter } from "next/navigation";
import { dashboardRoute } from "@/configs/routes";

type LoginInput = z.infer<typeof authResolver>;

const CustomForm: React.FC = () => {
	const router = useRouter();

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
			const response = await login(formData);
			router.push(dashboardRoute.routePath);
			console.log(response);
		} catch (e: Response) {
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
			</form>
		</Form>
	)
}

export default CustomForm;