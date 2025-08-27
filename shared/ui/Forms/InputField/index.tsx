import React from "react";
import { FormField } from "@/shared/ui/form";
import { Control } from "react-hook-form";
import { CustomField } from "@/shared/ui/CustomField";
import styles from "./styles.module.scss";

interface Props {
    name: string,
    label: string,
    control: Control,
    title: string,
    other_props?: any
}

export const InputField: React.FC<Props> = (props) => {
    return <FormField
        name={ props.name }
        control={ props.control }
        render={ ({ field, fieldState }) => (
            <CustomField
                field={ field }
                fieldName={ field.name }
                title={ props.label }
                customClassName={ styles.input }
                { ...props.other_props }
            />
        ) }
    />
}