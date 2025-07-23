import React, { useState } from "react";
import { Tooltip, TooltipContent, TooltipProvider, TooltipTrigger } from "@/shared/ui/tooltip";
import { Dialog, DialogContent, DialogTitle, DialogTrigger } from "@/shared/ui/dialog";
import Image from "next/image";
import { VisuallyHidden } from "@radix-ui/react-visually-hidden";
import PersonDialogMenuContent from "@/entities/faculty/PersonDialogMenuContent";
import arrowRightIcon from "@/public/icons/svg/right-arrow.svg";
import checkMarkIcon from "@/public/icons/svg/check-mark.svg";
import crossIcon from "@/public/icons/svg/cross.svg";
import userIcon from "@/public/icons/faculty/faculty-member/faculty-member-icon.svg";
import styles from "./styles.module.scss";
import { CourseTeacher } from "@/shared/types/ui/courses";

type Props = CourseTeacher;

const TAElement: React.FC<Props> = (props) => {
	const [openDialog, setOpenDialog] = useState(false);
	return (
		<TooltipProvider>
			<Dialog open={openDialog} onOpenChange={setOpenDialog}>
				<Tooltip>
					<div>
						<div className={styles.facultyElement}>
							<div className={ styles.buttons }>
								<Image src={ checkMarkIcon } alt={ "approve" } />
								<Image src={ crossIcon } alt={ "dis-approve" } />
							</div>
							<TooltipTrigger>
	              <span className={ styles.menuTrigger }>
                { props.profile_data.name_eng }
              </span>
							</TooltipTrigger>
						</div>
					</div>
					<TooltipContent side="right" className={ styles.contextMenu }>
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
												<span>{ props.profile_data.name_eng }</span>
												<Image src={arrowRightIcon} alt="go to profile" />
											</div>
										</DialogTrigger>
										<div className={styles.tg}>{ props.profile_data.alias }</div>
									</div>
								</div>
								<div className={styles.email}>{ props.profile_data.email }</div>
								<div className={styles.workInformation}>
									<div className={styles.department}>
										<div className={styles.placeholder}>Institute:</div>
										<div className={styles.value}>{ props.profile_data.institute_names.map(institute => <span key={ institute }>{ institute }</span>) }</div>
									</div>
									<div className={styles.department}>
										<div className={styles.placeholder}>Position:</div>
										<div className={styles.value}>{ props.profile_data.position_name }</div>
									</div>
								</div>
								{/*<div className={ styles.workload }>*/}
								{/*	<div className={ styles.block }>*/}
								{/*		<div className={ styles.number }>0.2</div>*/}
								{/*		<div className={ styles.text }>Rate T1</div>*/}
								{/*	</div>*/}
								{/*	<div className={ styles.block }>*/}
								{/*		<div className={ styles.number }>0.2</div>*/}
								{/*		<div className={ styles.text }>Rate T2</div>*/}
								{/*	</div>*/}
								{/*	<div className={ styles.block }>*/}
								{/*		<div className={ styles.number }>0.2</div>*/}
								{/*		<div className={ styles.text }>Rate T3</div>*/}
								{/*	</div>*/}
								{/*</div>*/}
							</div>
							{/*<ul className={styles.labList}>*/}
							{/*	<li className={styles.lab}>AGLA I - 2 lab</li>*/}
							{/*	<li className={styles.lab}>MA I - 3 labs</li>*/}
							{/*	<li className={styles.lab}>AGLA I - 1 tut</li>*/}
							{/*	<li className={styles.lab}>Smt else - 1 lab</li>*/}
							{/*	<li className={styles.lab}>Smt else - 1 lab</li>*/}
							{/*</ul>*/}
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
					<PersonDialogMenuContent id={ props.profile_data.profile_id } />
				</DialogContent>
			</Dialog>
		</TooltipProvider>
	);
};

export default TAElement;