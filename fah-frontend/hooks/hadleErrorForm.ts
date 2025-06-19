import { UseFormSetError, Path } from "react-hook-form";

type ServerErrorResponse = {
	status?: number,
	data?: {
		message?: string,
		errors?: Record<string, string>,
	}
}

export function handleErrorForm<T extends Record<string, any>>(
	error: unknown,
	setError: UseFormSetError<T>
) {
	const err = error as ServerErrorResponse;

	if (err?.status === 401) {
		setError("root", {
			message: "Invalid credentials"
		});
		return;
	}

	if (err?.status === 422 && err.data?.errors) {
		const errors = err.data.errors;

		for (const key in errors) {
			if (Object.prototype.hasOwnProperty.call(errors, key)) {
				const field = key as Path<T>;
				const message = errors[key];
				setError(field, { message });
			}
		}
		return;
	}

	setError("root", {
		message: err?.data?.message || "Something went wrong. Please try again."
	})
}