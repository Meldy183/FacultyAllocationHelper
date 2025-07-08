"use client";

import React, { useState } from "react";
import { Dialog, DialogContent, DialogDescription, DialogHeader, DialogTitle, DialogTrigger } from "@/shared/ui/dialog";
import { VisuallyHidden } from "@radix-ui/react-visually-hidden";
import { AnimatePresence, motion } from "framer-motion";
import AllocateExistingMember from "./AllocateMemberMenu";
import styles from "./styles.module.scss";
import CourseComposition from "./CourseComposition";

const positions = [
	"Primary instructor",
	"Tutor instructor",
	"Teacher assistant"
]

const AssignNewMember: React.FC = () => {
	const [changNow, setChangNow] = useState<string>("");

	const handleChangedFaculty = (member: string) => {
		setChangNow(member);
	}

	const foldMenu = () => {
		setChangNow("");
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
					handleAllocateMember={ handleChangedFaculty }
					faculties={ positions }
				/>
			</motion.div>
			<AnimatePresence>
				{
					positions.map((position) => (
						changNow == position &&
            <motion.div
                key="right"
                initial={ { opacity: 0, x: 50 } }
                animate={ { opacity: 1, x: 0 } }
                exit={ { opacity: 0, x: 50 } }
                transition={ { duration: 0.3 } }
                layout
                className={ styles.menu }
            >
			          <AllocateExistingMember
					          foldMenu={ foldMenu }
					          position={ position }
			          />
            </motion.div>
					))
				}
			</AnimatePresence>
		</DialogContent>
	</Dialog>
}

export default AssignNewMember;