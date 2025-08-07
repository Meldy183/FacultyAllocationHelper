import z from "zod";

export const CreateCourseResolver = z.object({
    brief_name: z.string(),
    academic_year_id: z.number(),
    semester_id: z.number(),
    year: z.number(),
    program_ids: z.array(z.number()),
    track_ids: z.array(z.number()),
    responsible_institute_id: z.number(),
    groups_needed: z.number(),
    is_elective: z.boolean()
})

export type CreateCourseType = z.infer<typeof CreateCourseResolver>;