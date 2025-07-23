"use client";

import { Provider } from "react-redux";
import { store } from "@/features/store";
import { ReactNode } from "react";

export function CustomProvider({ children }: { children: ReactNode }) {
	return <Provider store={ store }>{children}</Provider>;
}