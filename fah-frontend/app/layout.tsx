import type { Metadata } from "next";
import { CustomProvider } from "./providers";
// import { Geist, Geist_Mono, Montserrat } from "next/font/google";
import "./globals.css";

// const geistSans = Geist({
//   variable: "--font-geist-sans",
//   subsets: ["latin"],
// });
//
// const geistMono = Geist_Mono({
//   variable: "--font-geist-mono",
//   subsets: ["latin"],
// });
//
// const montserrat = Montserrat({
//   variable: "--font-montserrat",
//   subsets: ["latin"],
// })

export const metadata: Metadata = {
  title: "Faculty Allocation Helper",
  description: "Faculty Allocation Helper",
  icons: {
    icon: "/icons/favicon.ico"
  }
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
    {/*<body className={ `${ geistSans.variable } ${ geistMono.variable } ${ montserrat.variable } antialiased` }>*/}
      <body className={ `antialiased` }>
        <CustomProvider>
          { children }
        </CustomProvider>
      </body>
    </html>
);
}
