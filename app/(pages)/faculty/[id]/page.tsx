"use client";

import React from "react";
import Image from "next/image";
import userIcon from "@/public/icons/faculty/faculty-member/faculty-member-icon.svg"
import Wrapper from "@/shared/ui/wrapper";
import { useGetUserQuery } from "@/features/api/slises/profile";
import styles from "./styles.module.scss";
import { useParams } from "next/navigation";

export default function ProfileDashboard() {
	const params = useParams();

	const id = params.id as string;

	const { data, error, isLoading } = useGetUserQuery({ id });

	if (error) return <>smth went wrong (error)</>

	if (isLoading) return <>wating</>

	if (!data) return <>smth went wrong (no data)</>;

	return (
		<Wrapper>
			<div className={styles.container}>
				<div className={styles.card}>
					<div className={styles.header}>
						<div className={styles.userInfo}>
							<Image src={ userIcon } alt={ "user icon" } className={ styles.avatar } />
							<div>
								<h1 className={styles.name}>{ data.name_eng }</h1>
								<p className={styles.subName}>{ data.name_ru }</p>
							</div>
						</div>
					</div>

					{/* Profile Info */}
					<div className={styles.section}>
						<div className={styles.row}><strong>Position:</strong> { data.position_name }</div>
						<div className={styles.row}><strong>Institutes:</strong> { data.institute_names?.map(institute => <span key={ institute }>{ institute }</span>) }</div>
					</div>

					<div className={styles.section}>
						<h2 className={styles.title}>Personal Information</h2>
						<div className={styles.grid}>
							<div>Email: { data.email }</div>
							<div>Telegram alias: { data.alias }</div>
							<div>Student? { data.student_type }</div>
							<div>Responsible from FSRO: { data.fsro }</div>
							<div>Degree: { data.degree }</div>
							<div>Languages: { data.languages?.map(({ language_code }) => <span key={ language_code }>{ language_code }</span>) }</div>
						</div>
					</div>

					<div className={styles.section}>
						<h2 className={styles.title}>Employment</h2>
						<div className={styles.grid}>
							<div>Type of employment: { data.employnment_type }</div>
							<div>Start date: </div>
							<div>Hiring status: { data.hiring_status }</div>
							<div>End date: 00.00.0000</div>
							<div>Mode: { data.mode }</div>
						</div>
					</div>
				</div>

				{/* Workload */}
				<div className={styles.card}>
					<div className={styles.sectionCardWhite}>
						<h2 className={styles.name}>Workload</h2>
						<div className={styles.workloadGrid}>
							<div className={ styles.table }>
								<div className={ `${ styles.row } ${ styles.title}` }>
									<div className={ `${ styles.block } ${ styles.designationBlock }` }></div>
									<div className={ styles.block }>LEC</div>
									<div className={ styles.block }>TUT</div>
									<div className={ styles.block }>LAB</div>
									<div className={ styles.block }>ELECTIVE</div>
									<div className={ `${ styles.block } bg-[#40BA2180]` }>RATE</div>
								</div>
								<div className={ `${ styles.row }` }>
									<div className={ `${ styles.block } ${ styles.designationBlock }` }>T1</div>
									<div className={ styles.block }>{ data.workload_stats.t1?.lec_hours }</div>
									<div className={ styles.block }>{ data.workload_stats.t1?.tut_hours }</div>
									<div className={ styles.block }>{ data.workload_stats.t1?.lab_hours }</div>
									<div className={ styles.block }>{ data.workload_stats.t1?.elective_hours }</div>
									<div className={ styles.block }>{ data.workload_stats.t1?.rate }</div>
								</div>
								<div className={ `${ styles.row }` }>
									<div className={ `${ styles.block } ${ styles.designationBlock }` }>T2</div>
									<div className={ styles.block }>{ data.workload_stats.t2?.lec_hours }</div>
									<div className={ styles.block }>{ data.workload_stats.t2?.tut_hours }</div>
									<div className={ styles.block }>{ data.workload_stats.t2?.lab_hours }</div>
									<div className={ styles.block }>{ data.workload_stats.t2?.elective_hours }</div>
									<div className={ styles.block }>{ data.workload_stats.t2?.rate }</div>
								</div>
								<div className={ `${ styles.row }` }>
									<div className={ `${ styles.block } ${ styles.designationBlock }` }>T3</div>
									<div className={ styles.block }>{ data.workload_stats.t3?.lec_hours }</div>
									<div className={ styles.block }>{ data.workload_stats.t3?.tut_hours }</div>
									<div className={ styles.block }>{ data.workload_stats.t3?.lab_hours }</div>
									<div className={ styles.block }>{ data.workload_stats.t3?.elective_hours }</div>
									<div className={ styles.block }>{ data.workload_stats.t3?.rate }</div>
								</div>
								<div className={ `${ styles.row }` }>
									<div className={ `${ styles.block } ${ styles.designationBlock }` }>Total</div>
									<div className={ styles.block }>{ data.workload_stats.total.lec_hours }</div>
									<div className={ styles.block }>{ data.workload_stats.total.tut_hours }</div>
									<div className={ styles.block }>{ data.workload_stats.total.lab_hours }</div>
									<div className={ styles.block }>{ data.workload_stats.total.elective_hours }</div>
									<div className={ styles.block }>{ data.workload_stats.total.rate }</div>
								</div>
							</div>
							<div className={styles.metrics}>
								<div><span>Max load:</span><span>{ data.max_load }</span></div>
								<div><span>Frontal Hours:</span><span>{ data.frontal_hours }</span></div>
								<div><span>Extra activities:</span><span>{ data.extra_activities }</span></div>
							</div>
						</div>
					</div>

				</div>
				{/* Teaching Courses */}
				<div className={styles.card}>
					<h2 className={styles.name}>Teaching courses:</h2>

					<div className={ styles.items }>
						{/*<div className={ styles.field }><CourseField { ...courseMock } /></div>*/}
						{/*<div className={ styles.field }><CourseField { ...courseMock } /></div>*/}
						{/*<div className={ styles.field }><CourseField { ...courseMock } /></div>*/}
						{/*<div className={ styles.field }><CourseField { ...courseMock } /></div>*/}
					</div>
				</div>
			</div>
		</Wrapper>
	);
}
