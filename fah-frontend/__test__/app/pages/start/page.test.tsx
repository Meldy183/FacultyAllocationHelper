//@ts-ignore
import("@testing-library/jest-dom");
import { render, screen } from "@testing-library/react";
import StartPage from "@/app/start/page";

describe('Page', () => {
    it('renders a heading', () => {
        render(<StartPage />)

        const heading = screen.getByRole('main')

        expect(heading).toBeInTheDocument()
    })
})