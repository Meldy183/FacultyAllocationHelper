import React, { useEffect } from "react";
import { useLazyGetAllCoursesQuery } from "../api";
import { transformFilters } from "../lib";
import { useAppSelector } from "@/features/store/hooks";
import { useDebounce } from "@/shared/hooks/useDebounce";
import CourseField from "@/shared/ui/CourseField";
import styles from "./styles.module.scss";

export const CourseList: React.FC = () => {
    const filters = useAppSelector(state => state.courseFilters);
    const [getCourses, { data, error }] = useLazyGetAllCoursesQuery();

    const debouncedFilters = useDebounce(filters);

    useEffect(() => {
        getCourses(transformFilters(debouncedFilters));
    }, [debouncedFilters, getCourses]);

    return (
        <div className={ styles.container }>
            <div className={ styles.courses }>
                {
                    error && <span>smth went wrong. cant load courses</span>
                }
                {
                    data?.courses.map((course) => (
                        <div key={ course.course_id } className={ styles.field }><CourseField { ...course } /></div>
                    ))
                }
            </div>
        </div>
    )
}