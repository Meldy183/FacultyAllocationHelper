import z from "zod";

export const CreateCourseResolver = z.object({
    brief_name: z
        .string(),
    academic_year_id: z.number().min(1, {
        message: "choose the academic year"
    }),
    semester_id: z.number().min(1),
    year: z.number().min(1, {
        message: "Choose the year"
    }),
    program_ids: z.array(z.number()).min(1, {
        message: "choose the program"
    }),
    track_ids: z.array(z.number()).min(1,{
        message: "choose the track"
    }),
    responsible_institute_id: z.number().min(1, {
        message: "choose the institute"
    }),
    groups_needed: z.number().min(1, {
        message: "Choose the number of groups"
    }),
    is_elective: z.boolean()
});

export type CreateCourseType = z.infer<typeof CreateCourseResolver>;