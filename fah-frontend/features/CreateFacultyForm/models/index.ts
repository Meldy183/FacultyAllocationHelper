import { z } from "zod";

export const CreateFacultyResolver = z.object({
    name_eng: z.string(),
    email: z.string().email({}),
    alias: z.string(),
    institute_id: z.array(z.number()).min(1, {
        message: 'you must choose the institutes',
    }),
    position_id: z.number().min(1, {
        message: "please choose one of provided variants"
    }).max(5, {
        message: "please choose one of provided variants"
    }),
});

export type CreateFacultyType = z.infer<typeof CreateFacultyResolver>;