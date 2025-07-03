"use client";

import React, { useState } from "react";
import { Dialog, DialogContent, DialogDescription, DialogHeader, DialogTitle, DialogTrigger } from "@/shared/ui/dialog";
import { VisuallyHidden } from "@radix-ui/react-visually-hidden";
import { AnimatePresence, motion } from "framer-motion";
import styles from "./styles.module.scss";

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
	return <div className={ styles.allocateMember }>
		<button onClick={ () => handleCreateMember() }>click</button>
	</div>
}

const CreateNewMember: React.FC = () => {
	return <div className={ styles.createNewMember }></div>
}

export default AssignNewMember;