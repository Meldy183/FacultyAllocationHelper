import React from "react";
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

// Derive form types from Zod schema
type FormInputType = z.infer<typeof CreateMemberResolver>;

const CreateFacultyMenu: React.FC = () => {
  const form = useForm<FormInputType>({
    resolver: zodResolver(CreateMemberResolver),
    defaultValues: {
      name: "",
      email: "",
      alias: "",
      department: "",
      memberPosition: "",
    },
  });

  const submitHandler = (data: any) => {
    console.log("Submitted:", data);
    // TODO: handle API call and close dialog
  };

  return (
    <Dialog>
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

              {/* Name */}
              <FormField
                control={form.control}
                name="name"
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

              {/* E-mail */}
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

              {/* Alias */}
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

              {/* Department */}
              <FormField
                control={form.control}
                name="department"
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

              {/* Position Select */}
              <FormField
                control={form.control}
                name="memberPosition"
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
