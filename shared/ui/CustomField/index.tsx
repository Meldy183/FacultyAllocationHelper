import React from "react";
import { FormControl, FormItem, FormLabel } from "@/shared/ui/form";
import { Input } from "@/shared/ui/input";
import styles from "./styles.module.scss";

interface fieldProps {
  field: any,
  fieldName: string,
  title: string
  type?: string
  customClassName?: string,
  error?: string
}

export const CustomField: React.FC<fieldProps> = ({ field, fieldName, title, type = "text", customClassName = "" }) => {
  const autocompleteValue = type === "password" ? "current-password" : "off";

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