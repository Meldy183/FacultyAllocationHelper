"use client";

import React, { useState } from "react";
import { Dialog, DialogContent, DialogDescription, DialogHeader, DialogTitle, DialogTrigger } from "@/shared/ui/dialog";
import { VisuallyHidden } from "@radix-ui/react-visually-hidden";
import { AnimatePresence, motion } from "framer-motion";
import AllocateExistingMember from "./AllocateMemberMenu";
import CourseComposition from "./CourseComposition";
import { facultyPositions } from "@/shared/configs/constants/api/paths";
import styles from "./styles.module.scss";

const AssignNewMember: React.FC = () => {
	const [changNow, setChangNow] = useState<string>("");

	const handleChangedFaculty = (member: string) => {
		console.log(member);
		setChangNow(member);
	}

	const foldMenu = () => {
		setChangNow("");
	}

	return <Dialog>
		<DialogTrigger asChild>
			<button className={ styles.addNewMember }>Assign new</button>
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
					faculties={ facultyPositions }
				/>
			</motion.div>
			<AnimatePresence>
				{
					facultyPositions.map((position) => (
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