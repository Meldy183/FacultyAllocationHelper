import { z } from "zod";

export const registerResolver = z.object({
	email: z.string().email({
		message: "Email is required",
	}),
	password: z.string().min(8, {
		message: "Password must be greater than 8 characters"
	}),
	passwordAgain: z.string().min(8, {
		message: "Password must be greater than 8 characters"
	})
});

export const authResolver = z.object({
	email: z.string().email({
		message: "Email is required",
	}),
	password: z.string().min(9, {
		message: "Password must be greater than 8 characters"
	})
})