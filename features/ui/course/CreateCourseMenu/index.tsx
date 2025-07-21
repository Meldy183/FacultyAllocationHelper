import React, { useState } from "react";
import styles from "./styles.module.scss";
import { Button } from "@/shared/ui/button";
import { Dialog, DialogContent, DialogDescription, DialogHeader, DialogTitle, DialogTrigger } from "@/shared/ui/dialog";
import { VisuallyHidden } from "@radix-ui/react-visually-hidden";
import { Form, FormField } from "@/shared/ui/form";
import CustomField from "@/shared/ui/CustomField";
import { useForm, useWatch } from "react-hook-form";
import { CreateCourseResolver, CreateCourseType } from "@/shared/types/resolvers/course";
import { zodResolver } from "@hookform/resolvers/zod";
import { Label } from "@/shared/ui/label";
import { Switch } from "@/shared/ui/switch";
import { Checkbox } from "@/shared/ui/checkbox";
import { useCreateNewMutation } from "@/features/api/slises/courses";
import { handleErrorForm } from "@/shared/hooks/hadleErrorForm";

interface Props {
	children: React.ReactNode;
}

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

// const AddCourseMenu: React.FC<Props> = ({ children }) => {
// 	const submitHandler = (formData: CreateCourseType) => {
// 		console.log(formData);
// 	}
//
// 	const form = useForm<CreateCourseType>({
// 		resolver: zodResolver(CreateCourseResolver),
// 		defaultValues: {
// 			brief_name: "",
// 			academic_year_id: 0,
// 			semester_id: 0,
// 			year: 0,
// 			program_ids: [],
// 			track_ids: [],
// 			responsible_institute_id: 0,
// 			groups_needed: 0,
// 			is_elective: false
// 		}
// 	});
//
// 	const podglyadun = useWatch(
// 		form
// 	);
//
// 	const handleClick = () => {
// 		console.log(podglyadun)
// 	}
//
// 	return <Dialog>
// 		<DialogTrigger asChild>{ children }</DialogTrigger>
// 		<DialogContent>
// 			<VisuallyHidden>
// 				<DialogHeader>
// 					<DialogTitle />
// 					<DialogDescription />
// 				</DialogHeader>
// 			</VisuallyHidden>
// 			<div className={ styles.menu }>
// 				<div className={ styles.content }>
// 					<div className={ styles.title }>Add a new course</div>
// 					<Form { ...form }>
// 						<form onSubmit={form.handleSubmit(submitHandler)}>
// 							<FormField
// 								control={ form.control }
// 								name="brief_name"
// 								render={({ field, fieldState }) => (
// 									<div className={styles.fieldBlock}>
// 										<CustomField
// 											field={field}
// 											fieldName={field.name}
// 											title="course name"
// 											customClassName={ styles.input }
// 										/>
// 										{fieldState.error && <div className={styles.error}>{fieldState.error.message}</div>}
// 									</div>
// 								)} />
// 							<FormField
// 								control={ form.control }
// 								name="academic_year_id"
// 								render={({ field, fieldState }) => (
// 									<div className={styles.fieldBlock}>
// 										<CustomField
// 											field={field}
// 											fieldName={field.name}
// 											title="academic year id"
// 											customClassName={ styles.input }
// 										/>
// 										{fieldState.error && <div className={styles.error}>{fieldState.error.message}</div>}
// 									</div>
// 								)} />
// 							<FormField
// 								control={ form.control }
// 								name="semester_id"
// 								render={({ field, fieldState }) => (
// 									<div className={styles.fieldBlock}>
// 										<CustomField
// 											field={field}
// 											fieldName={field.name}
// 											title="semestr id"
// 											customClassName={ styles.input }
// 										/>
// 										{fieldState.error && <div className={styles.error}>{fieldState.error.message}</div>}
// 									</div>
// 								)} />
// 							<FormField
// 								control={ form.control }
// 								name="year"
// 								render={({ field, fieldState }) => (
// 									<div className={styles.fieldBlock}>
// 										<CustomField
// 											field={field}
// 											fieldName={field.name}
// 											title="academic year"
// 											customClassName={ styles.input }
// 										/>
// 										{fieldState.error && <div className={styles.error}>{fieldState.error.message}</div>}
// 									</div>
// 								)} />
// 							<FormField
// 								control={ form.control }
// 								name="groups_needed"
// 								render={({ field, fieldState }) => (
// 									<div className={styles.fieldBlock}>
// 										<CustomField
// 											field={field}
// 											fieldName={field.name}
// 											title="groups needed on course"
// 											customClassName={ styles.input }
// 										/>
// 										{fieldState.error && <div className={styles.error}>{fieldState.error.message}</div>}
// 									</div>
// 								)} />
// 							<FormField
// 								control={ form.control }
// 								name="is_elective"
// 								render={({ field }) => (
// 									<Label className={ "font-semibold text-[12px] leading-[18px] text-[#666666] mt-[24px]" }>
// 										<div>is elective course</div>
// 										<Switch onCheckedChange={ field.onChange } />
// 									</Label>
// 								)}
// 							/>
// 						</form>
// 					</Form>
// 					<Button type={ "submit" } onClick={ handleClick } className={ styles.button }>Submit</Button>
// 				</div>
// 			</div>
// 		</DialogContent>
// 	</Dialog>
// }

const AddCourseMenu: React.FC<Props> = ({ children }) => {
	const [isOpen, setIsOpen] = useState<boolean>(false);
	const [createCourse, { isLoading }] = useCreateNewMutation();

	const submitHandler = async (formData: CreateCourseType) => {
		try {
			const { data, error } = await createCourse(formData);
			if (error) throw error;
		} catch (e) {
			handleErrorForm(e, form.setError);
		}

		console.log(formData);
		setIsOpen(false);
		form.reset();
	};

	const form = useForm<CreateCourseType>({
		resolver: zodResolver(CreateCourseResolver),
		defaultValues: {
			brief_name: "",
			academic_year_id: undefined, // Changed to undefined for select default
			semester_id: undefined, // Changed to undefined for select default
			year: 0,
			program_ids: [],
			track_ids: [],
			responsible_institute_id: undefined, // Changed to undefined for select default
			groups_needed: 0,
			is_elective: false,
		},
	});

	return (
		<Dialog open={isOpen} onOpenChange={setIsOpen}>
			<DialogTrigger asChild>{children}</DialogTrigger>
			<DialogContent>
				<VisuallyHidden>
					<DialogHeader>
						<DialogTitle />
						<DialogDescription />
					</DialogHeader>
				</VisuallyHidden>
				<div className={styles.menu}>
					<div className={styles.content}>
						<div className={styles.title}>Add a new course</div>
						<Form {...form}>
							<form onSubmit={form.handleSubmit(submitHandler)}>
								<FormField
									control={form.control}
									name="brief_name"
									render={({ field, fieldState }) => (
										<div className={styles.fieldBlock}>
											<CustomField
												field={field}
												fieldName={field.name}
												title="Course Name"
												customClassName={styles.input}
											/>
											{fieldState.error && (
												<div className={styles.error}>
													{fieldState.error.message}
												</div>
											)}
										</div>
									)}
								/>

								<FormField
									control={form.control}
									name="academic_year_id"
									render={({ field, fieldState }) => (
										<div className={styles.fieldBlock}>
											<Label className={styles.fieldDescription}>
												Academic Year
											</Label>
											<select
												{...field}
												className={styles.input} // Apply your input styles to the select
												onChange={(e) => field.onChange(Number(e.target.value))}
												value={field.value || ""} // Ensure value is controlled
											>
												<option value="">Select an academic year</option>
												{ACADEMIC_YEARS.map((year) => (
													<option key={year.id} value={year.id}>
														{year.name}
													</option>
												))}
											</select>
											{fieldState.error && (
												<div className={styles.error}>
													{fieldState.error.message}
												</div>
											)}
										</div>
									)}
								/>

								<FormField
									control={form.control}
									name="semester_id"
									render={({ field, fieldState }) => (
										<div className={styles.fieldBlock}>
											<Label className={styles.fieldDescription}>Semester</Label>
											<select
												{...field}
												className={styles.input}
												onChange={(e) => field.onChange(Number(e.target.value))}
												value={field.value || ""}
											>
												<option value="">Select a semester</option>
												{SEMESTERS.map((semester) => (
													<option key={semester.id} value={semester.id}>
														{semester.name}
													</option>
												))}
											</select>
											{fieldState.error && (
												<div className={styles.error}>
													{fieldState.error.message}
												</div>
											)}
										</div>
									)}
								/>

								<FormField
									control={form.control}
									name="year"
									render={({ field, fieldState }) => (
										<div className={styles.fieldBlock}>
											<CustomField
												field={{
													...field,
													onChange: (e: React.ChangeEvent<HTMLInputElement>) =>
														field.onChange(Number(e.target.value)),
												}} // Ensure year is number
												fieldName={field.name}
												title="Year"
												type="number" // Set type to number
												customClassName={styles.input}
											/>
											{fieldState.error && (
												<div className={styles.error}>
													{fieldState.error.message}
												</div>
											)}
										</div>
									)}
								/>

								<FormField
									control={form.control}
									name="responsible_institute_id"
									render={({ field, fieldState }) => (
										<div className={styles.fieldBlock}>
											<Label className={styles.fieldDescription}>
												Responsible Institute
											</Label>
											<select
												{...field}
												className={styles.input}
												onChange={(e) => field.onChange(Number(e.target.value))}
												value={field.value || ""}
											>
												<option value="">Select an institute</option>
												{RESPONSIBLE_INSTITUTES.map((institute) => (
													<option key={institute.id} value={institute.id}>
														{institute.name}
													</option>
												))}
											</select>
											{fieldState.error && (
												<div className={styles.error}>
													{fieldState.error.message}
												</div>
											)}
										</div>
									)}
								/>

								<FormField
									control={form.control}
									name="groups_needed"
									render={({ field, fieldState }) => (
										<div className={styles.fieldBlock}>
											<CustomField
												field={{
													...field,
													onChange: (e: React.ChangeEvent<HTMLInputElement>) =>
														field.onChange(Number(e.target.value)),
												}} // Ensure groups_needed is number
												fieldName={field.name}
												title="Groups Needed on Course"
												type="number" // Set type to number
												customClassName={styles.input}
											/>
											{fieldState.error && (
												<div className={styles.error}>
													{fieldState.error.message}
												</div>
											)}
										</div>
									)}
								/>

								<FormField
									control={form.control}
									name="program_ids"
									render={({ field }) => (
										<div className={styles.fieldBlock}>
											<Label className={styles.fieldDescription}>Programs</Label>
											<div className={styles.checkboxGroup}>
												{PROGRAMS.map((program) => (
													<div
														key={program.id}
														className={styles.checkboxItem}
													>
														<Checkbox
															id={`program-${program.id}`}
															checked={field.value.includes(program.id)}
															onCheckedChange={(checked) => {
																return checked
																	? field.onChange([...field.value, program.id])
																	: field.onChange(
																		field.value.filter(
																			(value) => value !== program.id
																		)
																	);
															}}
														/>
														<label
															htmlFor={`program-${program.id}`}
															className="text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70"
														>
															{program.name}
														</label>
													</div>
												))}
											</div>
											{form.formState.errors.program_ids && (
												<div className={styles.error}>
													{form.formState.errors.program_ids.message}
												</div>
											)}
										</div>
									)}
								/>

								<FormField
									control={form.control}
									name="track_ids"
									render={({ field }) => (
										<div className={styles.fieldBlock}>
											<Label className={styles.fieldDescription}>Tracks</Label>
											<div className={styles.checkboxGroup}>
												{TRACKS.map((track) => (
													<div key={track.id} className={styles.checkboxItem}>
														<Checkbox
															id={`track-${track.id}`}
															checked={field.value.includes(track.id)}
															onCheckedChange={(checked) => {
																return checked
																	? field.onChange([...field.value, track.id])
																	: field.onChange(
																		field.value.filter(
																			(value) => value !== track.id
																		)
																	);
															}}
														/>
														<label
															htmlFor={`track-${track.id}`}
															className="text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70"
														>
															{track.name}
														</label>
													</div>
												))}
											</div>
											{form.formState.errors.track_ids && (
												<div className={styles.error}>
													{form.formState.errors.track_ids.message}
												</div>
											)}
										</div>
									)}
								/>

								<FormField
									control={form.control}
									name="is_elective"
									render={({ field }) => (
										<Label
											className={
												"font-semibold text-[12px] leading-[18px] text-[#666666] mt-[24px]"
											}
										>
											<div>Is Elective Course</div>
											<Switch
												checked={field.value}
												onCheckedChange={field.onChange}
											/>
										</Label>
									)}
								/>
								<Button type={"submit"} className={styles.button} onClick={ () => { console.log("clicked") } }>Submit</Button>
							</form>
						</Form>
					</div>
				</div>
			</DialogContent>
		</Dialog>
	);
};

export default AddCourseMenu;