"use client";
import { redirect } from "next/navigation";

// вынести все роуты в отдельный файл с перечислениями откуда->куда
export default function Home() {
  redirect("/home");
}
