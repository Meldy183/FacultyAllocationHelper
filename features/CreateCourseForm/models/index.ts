import z from "zod";

export const CreateCourseResolver = z.object({
    brief_name: z
        .string(),
    academic_year_id: z.number().min(1, {
        message: "choose the academic year"
    }),
    semester_id: z.number().min(1),
    year: z
        .string()
        .min(1, {
            message: "Choose the year"
        })
        .regex(/^[0-9]*$/, {
            message: "Please enter a valid year."
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
    groups_needed: z
        .string()
        .min(1, {
            message: "Choose the number of groups"
        })
        .regex(/^[0-9]*$/, {
            message: "Please enter a valid number of groups."
        }),
    is_elective: z.boolean()
})

export const ACADEMIC_YEARS = [
    { id: 1, name: "BS1" },
    { id: 2, name: "BS2" },
    { id: 3, name: "BS3" },
    { id: 4, name: "BS4" },
    { id: 5, name: "MS1" },
    { id: 6, name: "MS2" },
    { id: 7, name: "PhD1" },
    { id: 8, name: "PhD2" },
];

export const SEMESTERS = [
    { id: 1, name: "T1" },
    { id: 2, name: "T2" },
    { id: 3, name: "T3" },
];

export const RESPONSIBLE_INSTITUTES = [
    { id: 1, name: "DS" },
    { id: 2, name: "DS/Math" },
    { id: 3, name: "DS/SDE" },
    { id: 4, name: "GAMEDEV" },
    { id: 5, name: "HUM" },
    { id: 6, name: "RO" },
    { id: 7, name: "SDE" },
    { id: 8, name: "SNE" },
];

export const PROGRAMS = [
    { id: 1, name: "AI360" },
    { id: 2, name: "МОИИ" },
    { id: 3, name: "BS RO" },
    { id: 4, name: "AIDE" },
    { id: 5, name: "SE" },
    { id: 6, name: "SNE" },
    { id: 7, name: "ROCV" },
    { id: 8, name: "MSRO" },
    { id: 9, name: "TE" },
    { id: 10, name: "УРКИ" },
    { id: 11, name: "КБ" },
    { id: 12, name: "УнОД" },
    { id: 13, name: "УЦП" },
    { id: 14, name: "DS" },
    { id: 15, name: "R" },
    { id: 16, name: "ITE" },
    { id: 17, name: "ИиВТ" },
    { id: 18, name: "DSAI" },
    { id: 19, name: "CSE" },
];

export const TRACKS = [
    { id: 1, name: "AAI" },
    { id: 2, name: "AAIR" },
    { id: 3, name: "CS" },
    { id: 4, name: "CSDS" },
    { id: 5, name: "DS" },
    { id: 6, name: "GD" },
    { id: 7, name: "ITE" },
    { id: 8, name: "R" },
    { id: 9, name: "SD" },
    { id: 10, name: "SE" },
    { id: 11, name: "SNE" },
];

export type CreateCourseType = z.infer<typeof CreateCourseResolver>;