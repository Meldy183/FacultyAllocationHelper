import { z } from "zod";

export const CreateMemberResolver = z.object({
  name: z.string(),
  email: z.string().email({}),
  alias: z.string(),
  department: z.string(),
  memberPosition: z.string(),
})