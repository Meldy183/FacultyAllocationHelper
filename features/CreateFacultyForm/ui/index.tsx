import React from "react";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { useCreateFacultyMutation } from "../api";
import { handleErrorForm } from "@/shared/hooks/hadleErrorForm";
import { CreateFacultyResolver, CreateFacultyType } from "@/features/CreateFacultyForm";
import { instituteList, roleList } from "@/shared/configs/constants/ui";
import { Form, FormControl, FormField, FormItem, FormLabel } from "@/shared/ui/form";
import { MultiSelect } from "@/shared/ui/MultiSelect";
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "@/shared/ui/select";
import { CustomField } from "@/shared/ui/CustomField";
import { Button } from "@/shared/ui/button";
import styles from "./styles.module.scss";

interface IProps {
    onSubmit: (response?: any) => void
}

export const CreateNewFacultyForm: React.FC<IProps> = ({ onSubmit }) => {
    const [createUser, { isLoading }] = useCreateFacultyMutation();

    const form = useForm<CreateFacultyType>({
        resolver: zodResolver(CreateFacultyResolver),
        defaultValues: {
            name_eng: "",
            email: "",
            alias: "",
            institute_id: [],
            position_id: 1
        },
    });

    const submitHandler = async (formData: CreateFacultyType) => {
        try {
            const { data, error } = await createUser(formData);
            console.log(data, error);
            if (error) throw error;
            console.log(formData);
            onSubmit(data);
        } catch (e) {
            handleErrorForm<CreateFacultyType>(e, form.setError);
        }
    };

    return <Form { ...form }>
        <form onSubmit={ form.handleSubmit(submitHandler) } className={ styles.menu }>
            <div className={ styles.content }>
                <div className={ styles.title }>Add a Faculty Member</div>
                <FormField
                    control={ form.control }
                    name="name_eng"
                    render={ ({ field, fieldState }) => (
                        <div className={ styles.fieldBlock }>
                            <CustomField
                                field={ field }
                                fieldName={ field.name }
                                title="Name"
                                customClassName={ styles.input }
                            />
                            { fieldState.error && <div className={ styles.error }>{ fieldState.error.message }</div> }
                        </div>
                    ) }
                />
                <FormField
                    control={ form.control }
                    name="email"
                    render={ ({ field, fieldState }) => (
                        <div className={ styles.fieldBlock }>
                            <CustomField
                                field={ field }
                                fieldName={ field.name }
                                title="E-mail"
                                type="email"
                                customClassName={ styles.input }
                            />
                            { fieldState.error && <div className={ styles.error }>{ fieldState.error.message }</div> }
                        </div>
                    ) }
                />
                <FormField
                    control={ form.control }
                    name="alias"
                    render={ ({ field, fieldState }) => (
                        <div className={ styles.fieldBlock }>
                            <CustomField
                                field={ field }
                                fieldName={ field.name }
                                title="Alias"
                                customClassName={ styles.input }
                            />
                            { fieldState.error && <div className={ styles.error }>{ fieldState.error.message }</div> }
                        </div>
                    ) }
                />
                <FormField
                    control={ form.control }
                    name="position_id"
                    render={ ({ field, fieldState }) => (
                        <div className={ styles.fieldBlock }>
                            <div className={ styles.fieldDescription }>Position</div>
                            <FormControl>
                                <Select
                                    value={ field.value?.toString() }
                                    onValueChange={ (value) => field.onChange(Number(value)) }
                                >
                                    <SelectTrigger className={ styles.select }>
                                        <SelectValue placeholder="Select a position"/>
                                    </SelectTrigger>
                                    <SelectContent>
                                        {
                                            roleList.map(({ value, label }) => (
                                                <SelectItem
                                                    key={ value }
                                                    value={ value.toString() }>{ label }</SelectItem>
                                            ))
                                        }
                                    </SelectContent>
                                </Select>
                            </FormControl>
                            { fieldState.error && <div className={ styles.error }>{ fieldState.error.message }</div> }
                        </div>
                    ) }
                />
                <FormField name={ "institute_id" } render={
                    ({ field, fieldState }) => (
                        <FormItem className={ styles.fieldBlock }>
                            <FormLabel>Institute</FormLabel>
                            <FormControl>
                                <MultiSelect
                                    className={ styles.multiSelect }
                                    options={ instituteList.map(({ value, label }) => ({
                                        value: value.toString(),
                                        label
                                    })) }
                                    defaultValue={ field.value }
                                    onValueChange={ (value) => field.onChange(value.map(_val => Number(_val))) }
                                />
                            </FormControl>
                            { fieldState.error &&
                                <div className={ styles.error }>error: { fieldState.error.message } </div> }
                        </FormItem>
                    )
                }/>
                { form.formState.errors.root && (
                    <p className="text-red-500 text-sm">{ form.formState.errors.root.message }</p>
                ) }
                <Button type="submit" className={ styles.button }>{ isLoading ? "Submitting..." : "Submit" }</Button>
            </div>
        </form>
    </Form>
}