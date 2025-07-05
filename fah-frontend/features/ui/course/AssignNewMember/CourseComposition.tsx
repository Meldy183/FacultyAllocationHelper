import React, { useState } from "react";
import styles from "./styles.module.scss";
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "@/shared/ui/select";
import { Button } from "@/shared/ui/button";

interface Props {
  handleCreateMember: () => void
  handleAllocateMember: () => void
}

const CourseComposition: React.FC<Props> = ({ handleCreateMember, handleAllocateMember }) => {
  const [selectValue, setSelectValue] = useState<string>("Not assigned");

  const handleChange = (value: string) => {
    setSelectValue(value);
    if (value === "Create new faculty member") handleCreateMember();
    if (value === "Allocate existing members") handleAllocateMember();
  }

  return <div className={styles.allocation}>
    <div className={styles.head}>
      <div className={styles.title}>Allocate member</div>
    </div>
    <div className={styles.commonContainer}>
      <div className={ styles.allocateMember }>
        <div className={styles.allocationContainer}>
          <div className={styles.innerContainer}>
            <div className={styles.fieldDescription}>Primary instructor</div>
            <Select value={ selectValue } onValueChange={ handleChange }>
              <SelectTrigger className={ styles.selectValue }>
                <SelectValue placeholder={ "Not assigned" } />
              </SelectTrigger>
              <SelectContent>
                <SelectItem className={ "text-[#666666] p-4 bg-white hover:bg-[#ECF9E9] transition-colors duration-200" } value="Not assigned">Not assigned</SelectItem>
                <SelectItem className={ "text-[#666666] p-4 bg-white hover:bg-[#ECF9E9] transition-colors duration-200" } value="Not needed">Not needed</SelectItem>
                <SelectItem className={ "text-[#666666] p-4 bg-white hover:bg-[#ECF9E9] transition-colors duration-200" } value="Create new faculty member">Create new faculty member</SelectItem>
                <SelectItem className={ "text-[#666666] p-4 bg-white hover:bg-[#ECF9E9] transition-colors duration-200" } value="Allocate existing members">Allocate existing members</SelectItem>
              </SelectContent>
            </Select>
            <Button className={ styles.submitBtn }>Submit</Button>
          </div>
        </div>
      </div>
    </div>
  </div>
}

export default CourseComposition;