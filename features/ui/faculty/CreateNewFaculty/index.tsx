import React, { useState } from "react";
import styles from "./styles.module.scss";
import { Button } from "@/shared/ui/button";
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "@/shared/ui/select";
import { Dialog, DialogContent, DialogDescription, DialogHeader, DialogTitle, DialogTrigger } from "@/shared/ui/dialog";
import { VisuallyHidden } from "@radix-ui/react-visually-hidden";
import { Form, FormField, FormControl } from "@/shared/ui/form";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { z } from "zod";
import { CreateMemberResolver } from "@/shared/types/resolvers/createMember";
import CustomField from "@/shared/ui/CustomField";
import { useCreateUserMutation } from "@/features/api/slises/courses/members";
import { handleErrorForm } from "@/shared/hooks/hadleErrorForm";

type FormInputType = z.infer<typeof CreateMemberResolver>;

const CreateFacultyMenu: React.FC = () => {
  const [isOpen, setIsOpen] = useState(false);
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
      console.log(data, error);
       if (error) throw error;
       setIsOpen(false);
    } catch (e) {
      handleErrorForm<FormInputType>(e, form.setError);
    }
  };

  return (
    <Dialog open={ isOpen } onOpenChange={ setIsOpen }>
      <DialogTrigger asChild>
        <Button className={styles.button}>Add a new faculty member</Button>
      </DialogTrigger>
      <DialogContent className={styles.dialogContent}>
        <VisuallyHidden>
          <DialogHeader>
            <DialogTitle>Create New Faculty Member</DialogTitle>
            <DialogDescription>Fill out the form to add a new faculty member</DialogDescription>
          </DialogHeader>
        </VisuallyHidden>
        <Form {...form}>
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
              <Button type="submit" className={styles.button}>Submit</Button>
            </div>
          </form>
        </Form>
      </DialogContent>
    </Dialog>
  );
};

export default CreateFacultyMenu;
