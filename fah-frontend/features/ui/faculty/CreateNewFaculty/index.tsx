import React, { useState } from "react";
import styles from "./styles.module.scss";
import { Button } from "@/shared/ui/button";
import { Dialog, DialogContent, DialogDescription, DialogHeader, DialogTitle, DialogTrigger } from "@/shared/ui/dialog";
import { VisuallyHidden } from "@radix-ui/react-visually-hidden";
import CreateNewFacultyForm from "@/entities/faculty/CreateFacultyForm/ui";

const CreateFacultyMenu: React.FC = () => {
  const [isOpen, setIsOpen] = useState<boolean>(false);

  const onSubmit = () => {
    setIsOpen(false);
  }

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
        <CreateNewFacultyForm onSubmit={onSubmit} />
      </DialogContent>
    </Dialog>
  );
};

export default CreateFacultyMenu;
