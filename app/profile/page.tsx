import React from "react";
import CourseField from "@/shared/ui/CourseField";
import Image from "next/image";
// import Image from "next/image";
// import { Button } from "@/shared/ui/button";
// import Wrapper from "@/shared/ui/wrapper";
// import crossSvg from "@/public/icons/svg/cross.svg";
// import styles from "./styles.module.scss";
import styles from "./ProfileDashboard.module.scss";
import userIcon from "@/public/icons/faculty/faculty-member/faculty-member-icon.svg"
import { Button } from "@/shared/ui/button";
import Link from "next/link";
import Wrapper from "@/shared/ui/wrapper";


const faculty = {
	name: "Name",
	surname: "Surname",
	department: [],
	role: "TA",
	workload: 0.2
}

const courseMock = {
	courseName: "CourseName 1",
	PI: faculty,
	faculties: [faculty, faculty, faculty, faculty, faculty, faculty, faculty, faculty, faculty],
}

export default function ProfileDashboard() {
  return (
	<Wrapper>
    <div className={styles.container}>
      <div className={styles.card}>
        {/* Header */}
			<div className={styles.header}>
			<div className={styles.userInfo}>
				<Image src={ userIcon } alt={ "user icon" } className={ styles.avatar } />
				<div>
				<h1 className={styles.name}>Name Surname</h1>
				<p className={styles.subName}>Фамилия Имя Отчество</p>
				</div>
			</div>
			<Button><Link href={ "/" }>Edit profile</Link></Button>

			</div>

			{/* Profile Info */}
			<div className={styles.section}>
			<div className={styles.row}><strong>Position:</strong> Professor</div>
			<div className={styles.row}><strong>Institute:</strong> Институт разработки ПО и программной инженерии</div>
			</div>

			<div className={styles.section}>
			<h2 className={styles.title}>Personal Information</h2>
			<div className={styles.grid}>
				<div>Email: n.surname@innopolis.university</div>
				<div>Telegram alias: @alias</div>
				<div>Student? StudentType</div>
				<div>Responsible from FSRO: Employee</div>
				<div>Degree: With/Without</div>
				<div>Languages: Eng/Ru/Eng, Ru</div>
			</div>
			</div>

			<div className={styles.section}>
			<h2 className={styles.title}>Employment</h2>
			<div className={styles.grid}>
				<div>Type of employment: Combination of positions</div>
				<div>Start date: 00.00.0000</div>
				<div>Hiring status: Status</div>
				<div>End date: 00.00.0000</div>
				<div>Mode: Remote</div>
				
			</div>
			</div>
		</div>

        {/* Workload */}
		<div className={styles.card}>
		<div className={styles.sectionCardWhite}>
			<h2 className={styles.title}>Workload</h2>
          <div className={styles.workloadGrid}>
            <table className={styles.table}>
              <thead>
                <tr>
                  <th></th>
                  <th>LEC</th>
                  <th>TUT</th>
                  <th>LAB</th>
                  <th>ELECTIVE</th>
                  <th>RATE</th>
                </tr>
              </thead>
              <tbody>
                <tr>
                  <td>T1</td>
                  <td>10</td>
                  <td>5</td>
                  <td>8</td>
                  <td>3</td>
                  <td>0.25</td>
                </tr>
                <tr>
                  <td>T2</td>
                  <td>12</td>
                  <td>6</td>
                  <td>9</td>
                  <td>2</td>
                  <td>0.30</td>
                </tr>
                <tr>
                  <td>T3</td>
                  <td>8</td>
                  <td>4</td>
                  <td>6</td>
                  <td>1</td>
                  <td>0.20</td>
                </tr>
                <tr className={styles.totalRow}>
                  <td>Total</td>
                  <td>30</td>
                  <td>15</td>
                  <td>23</td>
                  <td>6</td>
                  <td>0.75</td>
                </tr>
              </tbody>
            </table>
            <div className={styles.metrics}>
              <div><span>Workload:</span><span className={styles.highlight}>0.75</span></div>
              <div><span>Max load:</span><span>N</span></div>
              <div><span>Frontal Hours:</span><span>N</span></div>
              <div><span>Extra activities:</span><span>N</span></div>
            </div>
          </div>
        </div>

		</div>
        {/* Teaching Courses */}
		<div className={styles.card}>
			<h2 className={styles.title}>Teaching courses:</h2>
			
				<div className={ styles.field }><CourseField { ...courseMock } /></div>
				<div className={ styles.field }><CourseField { ...courseMock } /></div>
				<div className={ styles.field }><CourseField { ...courseMock } /></div>
			</div>
    </div>
	</Wrapper>
  );
}
