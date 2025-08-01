import { render, screen } from '@testing-library/react';
import StartPage from '@/app/(pages)/start/page';
import { routesAuth } from '@/shared/configs/routes';
import '@testing-library/jest-dom';

describe('StartPage Component', () => {
    test('renders without crashing', () => {
        render(<StartPage />);
        expect(screen.getByText('Faculty allocation helper')).toBeInTheDocument();
    });

    test('renders title and description', () => {
        render(<StartPage />);
        const titleElement = screen.getByText('Faculty allocation helper');
        const descriptionElement = screen.getByText(/To get access to the service, please, login through your Innopolis University e-mail./);
        expect(titleElement).toBeInTheDocument();
        expect(descriptionElement).toBeInTheDocument();
    });

    test('renders buttons based on routesAuth', () => {
        render(<StartPage />);
        routesAuth.forEach(({ routeName, routePath }) => {
            const buttonElement = screen.getByText(routeName);
            expect(buttonElement).toBeInTheDocument();
            expect(buttonElement.closest('a')).toHaveAttribute('href', routePath);
        });
    });

    test('applies correct styles', () => {
        render(<StartPage />);
        const mainElement = screen.getByText('Faculty allocation helper').closest('main');
        expect(mainElement).toHaveClass('main');
        const contentElement = screen.getByText('Faculty allocation helper').parentElement;
        expect(contentElement).toHaveClass('content');
        const titleElement = screen.getByText('Faculty allocation helper');
        expect(titleElement).toHaveClass('title');
        const descriptionElement = screen.getByText(/To get access to the service, please, login through your Innopolis University e-mail./);
        expect(descriptionElement).toHaveClass('description');
        // const buttonsElement = screen!.getByText(routesAuth[0]!.routeName)!.parentElement.parentElement;
        // expect(buttonsElement).toHaveClass('buttons');
    });
});