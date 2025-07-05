import { z } from "zod";

export const CreateMemberResolver = z.object({
  nameEng: z.string(),
  email: z.string().email({}),
  alias: z.string(),
  institute: z.string(),
  position: z.string(),
})