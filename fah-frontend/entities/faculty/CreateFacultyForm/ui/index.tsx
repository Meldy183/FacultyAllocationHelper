import React from "react";
import { Form, FormControl, FormField } from "@/shared/ui/form";
import CustomField from "@/shared/ui/CustomField";
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "@/shared/ui/select";
import { Button } from "@/shared/ui/button";
import { handleErrorForm } from "@/shared/hooks/hadleErrorForm";
import { useForm } from "react-hook-form";
import styles from "@/features/ui/faculty/CreateNewFaculty/styles.module.scss";
import { zodResolver } from "@hookform/resolvers/zod";
import { CreateMemberResolver } from "@/shared/types/resolvers/createMember";
import { z } from "zod";
import { useCreateUserMutation } from "@/features/api/slises/profile";

type FormInputType = z.infer<typeof CreateMemberResolver>;

interface IProps {
  onSubmit: (response?: any) => void
}

const CreateNewFacultyForm: React.FC<IProps> = ({ onSubmit }) => {
  const [createUser] = useCreateUserMutation();

  const form = useForm<FormInputType>({
    resolver: zodResolver(CreateMemberResolver),
    defaultValues: {
      nameEng: "",
      email: "",
      alias: "",
      institute: "",
      position: "",
    },
  });

  const submitHandler = async (formData: FormInputType) => {
    try {
      const { data, error } = await createUser(formData);
      if (error) throw error;
      console.log(data)
      onSubmit(data);
    } catch (e) {
      handleErrorForm<FormInputType>(e, form.setError);
    }
  };

  return <Form {...form}>
    <form onSubmit={form.handleSubmit(submitHandler)} className={styles.menu}>
      <div className={styles.content}>
        <div className={styles.title}>Add a Faculty Member</div>
        <FormField
          control={form.control}
          name="nameEng"
          render={({ field, fieldState }) => (
            <div className={styles.fieldBlock}>
              <CustomField
                field={field}
                fieldName={field.name}
                title="Name"
                customClassName={styles.input}
              />
              {fieldState.error && <div className={styles.error}>{fieldState.error.message}</div>}
            </div>
          )}
        />
        <FormField
          control={form.control}
          name="email"
          render={({ field, fieldState }) => (
            <div className={styles.fieldBlock}>
              <CustomField
                field={field}
                fieldName={field.name}
                title="E-mail"
                type="email"
                customClassName={styles.input}
              />
              {fieldState.error && <div className={styles.error}>{fieldState.error.message}</div>}
            </div>
          )}
        />
        <FormField
          control={form.control}
          name="alias"
          render={({ field, fieldState }) => (
            <div className={styles.fieldBlock}>
              <CustomField
                field={field}
                fieldName={field.name}
                title="Alias"
                customClassName={styles.input}
              />
              {fieldState.error && <div className={styles.error}>{fieldState.error.message}</div>}
            </div>
          )}
        />
        <FormField
          control={form.control}
          name="institute"
          render={({ field, fieldState }) => (
            <div className={styles.fieldBlock}>
              <CustomField
                field={field}
                fieldName={field.name}
                title="Department"
                customClassName={styles.input}
              />
              {fieldState.error && <div className={styles.error}>{fieldState.error.message}</div>}
            </div>
          )}
        />
        <FormField
          control={form.control}
          name="position"
          render={({ field, fieldState }) => (
            <div className={styles.fieldBlock}>
              <div className={styles.fieldDescription}>Position</div>
              <FormControl>
                <Select
                  value={field.value}
                  onValueChange={field.onChange}
                >
                  <SelectTrigger className={styles.select}>
                    <SelectValue placeholder="Select a position" />
                  </SelectTrigger>
                  <SelectContent>
                    <SelectItem value="Primary instructor">Primary instructor</SelectItem>
                    <SelectItem value="Tutor instructor">Tutor instructor</SelectItem>
                    <SelectItem value="Teaching assistant">Teaching assistant</SelectItem>
                  </SelectContent>
                </Select>
              </FormControl>
              {fieldState.error && <div className={styles.error}>{fieldState.error.message}</div>}
            </div>
          )}
        />
        {form.formState.errors.root && (
          <p className="text-red-500 text-sm">{ form.formState.errors.root.message }</p>
        )}
        <Button type="submit" className={styles.button}>Submit</Button>
      </div>
    </form>
  </Form>
}

export default CreateNewFacultyForm;