import React, { useState } from "react";
import { Tooltip, TooltipContent, TooltipProvider, TooltipTrigger } from "@/shared/ui/tooltip";
import { Dialog, DialogContent, DialogTitle, DialogTrigger } from "@/shared/ui/dialog";
import styles from "@/shared/ui/CourseField/styles.module.scss";
import Image from "next/image";
import userIcon from "@/public/icons/faculty/faculty-member/faculty-member-icon.svg";
import arrowRightIcon from "@/public/icons/svg/right-arrow.svg";
import { VisuallyHidden } from "@radix-ui/react-visually-hidden";
import PersonDialogMenuContent from "../../../../entities/faculty/PersonDialogMenuContent";

interface Faculty {
	name: string;
	surname: string;
	department: string[];
	role: string;
	workload: number;
}

const TAElement: React.FC<Faculty> = (faculty) => {
	const [openDialog, setOpenDialog] = useState(false);
	return (
		<TooltipProvider>
			<Dialog open={openDialog} onOpenChange={setOpenDialog}>
				<Tooltip>
					<TooltipTrigger asChild>
						<div className={styles.facultyElement}>
              <span className={styles.menuTrigger}>
                {faculty.name[0] + ". " + faculty.surname}
              </span>
						</div>
					</TooltipTrigger>
					<TooltipContent side="right" className={styles.contextMenu}>
						<div className={styles.menu}>
							<div className={styles.header}>
								<div className={styles.head}>
									<Image
										src={userIcon}
										alt="user icon"
										className={styles.userImage}
									/>
									<div className={styles.information}>
										<DialogTrigger className={ "cursor-pointer" } asChild>
											<div
												title="Go to profile"
												className={styles.name}
												onClick={() => setOpenDialog(true)}
											>
												<span>Name Surname</span>
												<Image src={arrowRightIcon} alt="go to profile" />
											</div>
										</DialogTrigger>
										<div className={styles.tg}>@alias</div>
									</div>
								</div>
								<div className={styles.email}>n.surname@innopolis.university</div>
								<div className={styles.workInformation}>
									<div className={styles.department}>
										<div className={styles.placeholder}>Department:</div>
										<div className={styles.value}>______________</div>
									</div>
									<div className={styles.department}>
										<div className={styles.placeholder}>Position:</div>
										<div className={styles.value}>______________</div>
									</div>
									<div className={styles.workload}>
										<div className={styles.number}>0.2</div>
										<div className={styles.text}>Workload</div>
									</div>
								</div>
							</div>
							<ul className={styles.labList}>
								<li className={styles.lab}>AGLA I - 2 lab</li>
								<li className={styles.lab}>MA I - 3 labs</li>
								<li className={styles.lab}>AGLA I - 1 tut</li>
								<li className={styles.lab}>Smt else - 1 lab</li>
								<li className={styles.lab}>Smt else - 1 lab</li>
							</ul>
						</div>
					</TooltipContent>
				</Tooltip>
				<VisuallyHidden>
					<DialogTitle></DialogTitle>
					<DialogContent />
				</VisuallyHidden>
				<DialogContent className={styles.dialogMenu}>
					<VisuallyHidden>
						<div className={styles.header}>
							<h2>Меню</h2>
						</div>
					</VisuallyHidden>
					<PersonDialogMenuContent />
				</DialogContent>
			</Dialog>
		</TooltipProvider>
	);
};

export default TAElement;