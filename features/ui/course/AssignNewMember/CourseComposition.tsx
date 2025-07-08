import React, { useState } from "react";
import styles from "./styles.module.scss";
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "@/shared/ui/select";
import { Button } from "@/shared/ui/button";

interface Props {
  handleAllocateMember: (member: string) => void
  faculties: string[]
}

const CourseComposition: React.FC<Props> = ({ handleAllocateMember, faculties }) => {
  const handleChange = (value: string, member: string) => {
    if (value === "Allocate existing members") handleAllocateMember(member);
  }

  return <div className={styles.allocation}>
    <div className={styles.head}>
      <div className={styles.title}>Allocate member</div>
    </div>
    <div className={styles.commonContainer}>
      <div className={ styles.allocateMember }>
        <div className={styles.allocationContainer}>
          <div className={ styles.innerContainer }>
            {
              faculties.map((faculty) => (
                <div key={ faculty }>
                  <FacultyAllocationSelect faculty={ faculty } handleAllocateMember={ handleChange } />
                </div>
              ))
            }
            <Button className={ styles.submitBtn }>Submit</Button>
          </div>
        </div>
      </div>
    </div>
  </div>
}

interface FacultySelectProps {
  faculty: string
  handleAllocateMember: (value: string, member: string) => void
}

const FacultyAllocationSelect: React.FC<FacultySelectProps> = ({ faculty, handleAllocateMember }) => {
  const [assignedValue, setAssignedValue] = useState<string>("");

  const handleChange = (value: string) => {
    setAssignedValue(value);
    handleAllocateMember(value, faculty);
  }

  return <>
    <div className={ styles.fieldDescription }>{ faculty }</div>
    <Select
      value={ assignedValue }
      onValueChange={ (value) => handleChange(value) }
    >
      <SelectTrigger className={ styles.selectValue }>
        <SelectValue placeholder={ "Not assigned" }/>
      </SelectTrigger>
      <SelectContent>
        <SelectItem
          className={ "text-[#666666] p-4 bg-white hover:bg-[#ECF9E9] transition-colors duration-200" }
          value="Not assigned">Not assigned</SelectItem>
        <SelectItem
          className={ "text-[#666666] p-4 bg-white hover:bg-[#ECF9E9] transition-colors duration-200" }
          value="Allocate existing members">Allocate members</SelectItem>
      </SelectContent>
    </Select></>
}

export default CourseComposition;