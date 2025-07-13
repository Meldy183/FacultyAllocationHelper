import { z } from "zod";

export const CreateMemberResolver = z.object({
  name_eng: z.string(),
  email: z.string().email({}),
  alias: z.string(),
  institute_id: z.number(),
  position_id: z.number().min(1, {
    message: "please choose one of provided variants"
  }).max(5, {
    message: "please choose one of provided variants"
  }),
  is_repr: z.boolean(),
});

export type CreateMemberType = z.infer<typeof CreateMemberResolver>;