"use client";

import React, { useState } from "react";
import { Dialog, DialogContent, DialogDescription, DialogHeader, DialogTitle, DialogTrigger } from "@/shared/ui/dialog";
import { VisuallyHidden } from "@radix-ui/react-visually-hidden";
import { AnimatePresence, motion } from "framer-motion";
import styles from "./styles.module.scss";
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "@/shared/ui/select";
import { Input } from "@/shared/ui/input";

const AssignNewMember: React.FC = () => {
	const [createMember, setCreateMember] = useState<boolean>(false);

	const handleCreateMember = () => {
		setCreateMember(true);
	}

	return <Dialog>
		<DialogTrigger>
			<div className={ styles.addNewMember }>Assign new</div>
		</DialogTrigger>
		<DialogContent className={ styles.dialogContent }>
			<VisuallyHidden>
				<DialogHeader />
				<DialogTitle />
				<DialogDescription />
			</VisuallyHidden>
			<motion.div
				layout
				transition={{ type: "spring", stiffness: 300, damping: 30 }}
				className={ styles.menu }><AllocateMember handleCreateMember={ handleCreateMember } />
			</motion.div>
			<AnimatePresence>
				{
					createMember &&
					<motion.div
              key="right"
              initial={{ opacity: 0, x: 50 }}
              animate={{ opacity: 1, x: 0 }}
              exit={{ opacity: 0, x: 50 }}
              transition={{ duration: 0.3 }}
              layout
							className={ styles.menu }>
								<CreateNewMember/>
					</motion.div>
				}
			</AnimatePresence>
		</DialogContent>
	</Dialog>
}

interface Props {
	handleCreateMember: () => void
}

const AllocateMember: React.FC<Props> = ({ handleCreateMember }) => {
	const [selectValue, setSelectValue] = useState<string>("Not assigned");

	const handleChange = (value: string) => {
		setSelectValue(value);
		if (value === "Create new faculty member") handleCreateMember();
	}

	return <div className={ styles.allocateMember }>
		<div>primary instructor</div>
		<Select value={ selectValue } onValueChange={ handleChange }>
			<SelectTrigger>
				<SelectValue placeholder={ "Not assigned" } />
			</SelectTrigger>
			<SelectContent>
				<SelectItem value="Not assigned">Not assigned</SelectItem>
				<SelectItem value="Not needed">not needed</SelectItem>
				<SelectItem value="Create new faculty member">Create new faculty member</SelectItem>
			</SelectContent>
		</Select>
	</div>
}

const CreateNewMember: React.FC = () => {
	return <div className={ styles.createNewMember }>
		<div className={ styles.memberField }>
			<div className={ styles.title }>Name</div>
			<Input placeholder={ "Enter the member’s name" }/>
		</div>
		<div className={ styles.memberField }>
			<div className={ styles.title }>E-mail</div>
			<Input placeholder={ "Enter member’s email" }/>
		</div>
		<div className={ styles.memberField }>
			<div className={ styles.title }>Alias</div>
			<Input placeholder={ "Enter the member’s alias" }/>
		</div>
		<div className={ styles.memberField }>
			<div className={ styles.title }>Institute</div>
			<Input placeholder={ "Enter the member’s institute" }/>
		</div>
		<div className={ styles.memberField }>
			<div className={ styles.title }>Position</div>
			<Select>
				<SelectTrigger>
					<SelectValue placeholder={ "Enter the member’s position" } />
				</SelectTrigger>
				<SelectContent>
					<SelectItem value="TA">TA</SelectItem>
					<SelectItem value="PI">PI</SelectItem>
					<SelectItem value="TI">TI</SelectItem>
				</SelectContent>
			</Select>
		</div>
	</div>
}

export default AssignNewMember;