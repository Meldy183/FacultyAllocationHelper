"use client";

import React from "react";
import Wrapper from "@/shared/ui/wrapper";
import { useParams } from "next/navigation";
import { FacultyMenu } from "@/entities/faculty";

export default function ProfileDashboard() {
	const params = useParams();

	const id = params.id as string;

	return (
		<Wrapper>
			<FacultyMenu id={ id } />
		</Wrapper>
	);
}
