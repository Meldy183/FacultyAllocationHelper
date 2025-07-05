"use client";

import React, { useState } from "react";
import { Dialog, DialogContent, DialogDescription, DialogHeader, DialogTitle, DialogTrigger } from "@/shared/ui/dialog";
import { VisuallyHidden } from "@radix-ui/react-visually-hidden";
import { AnimatePresence, motion } from "framer-motion";
import CreateNewMemberMenu from "./CreateNewMemberMenu";
import AllocateExistingMember from "./AllocateExistingMemberMenu";
import styles from "./styles.module.scss";
import CourseComposition from "./CourseComposition";

const AssignNewMember: React.FC = () => {
	const [creatingMember, setCreatingMember] = useState<boolean>(false);
	const [allocatingMember, setAllocatingMember] = useState<boolean>(false);

	const [] = useState<string>();

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
				<CourseComposition
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
							<CreateNewMemberMenu position={ "Position-name" } handleFoldMenu={ handleFoldCreateMember } />
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

export default AssignNewMember;