"use client";

import React, { useState } from "react";
import { Dialog, DialogContent, DialogDescription, DialogHeader, DialogTitle, DialogTrigger } from "@/shared/ui/dialog";
import { VisuallyHidden } from "@radix-ui/react-visually-hidden";
import { AnimatePresence, motion } from "framer-motion";
import styles from "./styles.module.scss";
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "@/shared/ui/select";
import CreateNewMemberMenu from "@/app/courses/addCourseMenu/CreateNewMemberMenu";
import AllocateExistingMember from "@/app/courses/addCourseMenu/AllocateExistingMemberMenu";

const AssignNewMember: React.FC = () => {
	const [creatingMember, setCreatingMember] = useState<boolean>(false);
	const [allocatingMember, setAllocatingMember] = useState<boolean>(false);

	const handleCreateMember = () => {
		setCreatingMember(true);
		setAllocatingMember(false);
	}

	const handleFoldCreateMember = () => {
		setCreatingMember(false);
	}

	const handleAllocateMember = () => {
		setAllocatingMember(true);
		setCreatingMember(false);
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
				className={ styles.menu }>
				<AllocateMember
					handleCreateMember={ handleCreateMember }
					handleAllocateMember={ handleAllocateMember }
				/>
			</motion.div>
			<AnimatePresence>
				{
					creatingMember &&
					<motion.div
              key="right"
              initial={{ opacity: 0, x: 50 }}
              animate={{ opacity: 1, x: 0 }}
              exit={{ opacity: 0, x: 50 }}
              transition={{ duration: 0.3 }}
              layout
							className={ styles.menu }
					>
							<CreateNewMemberMenu position={ "position-name" } handleFoldMenu={ handleFoldCreateMember } />
					</motion.div>
				}
				{
					allocatingMember &&
					<motion.div
              key="right"
              initial={{ opacity: 0, x: 50 }}
              animate={{ opacity: 1, x: 0 }}
              exit={{ opacity: 0, x: 50 }}
              transition={{ duration: 0.3 }}
              layout
              className={ styles.menu }
					>
							<AllocateExistingMember />
					</motion.div>
				}
			</AnimatePresence>
		</DialogContent>
	</Dialog>
}

interface Props {
	handleCreateMember: () => void
	handleAllocateMember: () => void
}

const AllocateMember: React.FC<Props> = ({ handleCreateMember, handleAllocateMember }) => {
	const [selectValue, setSelectValue] = useState<string>("Not assigned");

	const handleChange = (value: string) => {
		setSelectValue(value);
		if (value === "Create new faculty member") handleCreateMember();
		if (value === "Allocate existing members") handleAllocateMember();
	}

	return <div className={ styles.allocateMember }>
		<div>primary instructor</div>
		<Select value={ selectValue } onValueChange={ handleChange }>
			<SelectTrigger className={ styles.selectValue }>
				<SelectValue placeholder={ "Not assigned" } />
			</SelectTrigger>
			<SelectContent>
				<SelectItem className={ "text-[#666666] p-4 bg-white hover:bg-[#ECF9E9] transition-colors duration-200" } value="Not assigned">Not assigned</SelectItem>
				<SelectItem className={ "text-[#666666] p-4 bg-white hover:bg-[#ECF9E9] transition-colors duration-200" } value="Not needed">not needed</SelectItem>
				<SelectItem className={ "text-[#666666] p-4 bg-white hover:bg-[#ECF9E9] transition-colors duration-200" } value="Create new faculty member">Create new faculty member</SelectItem>
				<SelectItem className={ "text-[#666666] p-4 bg-white hover:bg-[#ECF9E9] transition-colors duration-200" } value="Allocate existing members">Allocate existing members</SelectItem>
			</SelectContent>
		</Select>
	</div>
}

export default AssignNewMember;