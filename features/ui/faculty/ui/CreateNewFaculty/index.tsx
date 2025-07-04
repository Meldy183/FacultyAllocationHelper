import React from "react";
import styles from "./styles.module.scss";
import Link from "next/link";
import Image from "next/image";
import crossSvg from "@/public/icons/svg/cross.svg";
import { Button } from "@/shared/ui/button";
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "@/shared/ui/select";
import { Dialog, DialogContent, DialogDescription, DialogHeader, DialogTitle, DialogTrigger } from "@/shared/ui/dialog";
import { VisuallyHidden } from "@radix-ui/react-visually-hidden";

const CreateFacultyMenu: React.FC = () => {
  return <Dialog>
    <DialogTrigger asChild>
      <Button className={ styles.button }>Add a new	faculty member</Button>
    </DialogTrigger>
    <DialogContent className={ styles.dialogContent }>
      <VisuallyHidden>
        <DialogHeader>
          <DialogTitle />
          <DialogDescription />
        </DialogHeader>
      </VisuallyHidden>
      <div className={ styles.menu }>
        <div className={ styles.content }>
          <div className={ styles.title }>Add a TA</div>
          <div className={ styles.fieldBlock }>
            <div className={ styles.fieldDescription }>Name</div>
            <input type="text" placeholder={ "Enter the name" } className={ styles.input }/>
          </div>
          <div className={ styles.fieldBlock }>
            <div className={ styles.fieldDescription }>E-mail</div>
            <input type="text" placeholder={ "Enter TA’s email" } className={ styles.input }/>
          </div>
          <div className={ styles.fieldBlock }>
            <div className={ styles.fieldDescription }>Alias</div>
            <input type="text" placeholder={ "Enter the member’s alias" } className={ styles.input }/>
          </div>
          <div className={ styles.fieldBlock }>
            <div className={ styles.fieldDescription }>Phone number</div>
            <input type="text" placeholder={ "Enter the member’s phone number" } className={ styles.input }/>
          </div>
          <div className={ styles.fieldBlock }>
            <div className={ styles.fieldDescription }>Department</div>
            <input type="text" placeholder={ "Enter the member’s department" } className={ styles.input }/>
          </div>
          <CustomSelect />
          <Link href={ "/faculty" }><Button className={ styles.button }>Submit</Button></Link>
        </div>
      </div>
    </DialogContent>
  </Dialog>
}

const CustomSelect: React.FC = () => {
  return (
    <Select>
      <SelectTrigger className={ styles.select }>
        <SelectValue placeholder="Enter the member’s position" />
      </SelectTrigger>
      <SelectContent>
        <SelectItem value="Primary instructor">Primary instructor</SelectItem>
        <SelectItem value="Tutor instructor">Tutor instructor</SelectItem>
        <SelectItem value="teacher assistant">teacher assistant</SelectItem>
      </SelectContent>
    </Select>
  )
}

export default CreateFacultyMenu;