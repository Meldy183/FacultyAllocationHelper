import React from 'react';
import '@testing-library/jest-dom';
import { render, screen, fireEvent } from '@testing-library/react';
import ProfilePage from '@/app/profile/page';

describe('RegistrationPage', () => {
	beforeEach(() => {
		render(<ProfilePage />);
	});

	it('renders the cross icon', () => {
		const img = screen.getByRole('img', { hidden: true });
		expect(img).toBeInTheDocument();
	});

	it('renders the title', () => {
		expect(screen.getByText('Registration/Login')).toBeInTheDocument();
	});

	it('renders email and password inputs with correct placeholders and types', () => {
		const emailInput = screen.getByPlaceholderText('your email');
		expect(emailInput).toBeInTheDocument();
		expect(emailInput).toHaveAttribute('type', 'text');

		const passInput = screen.getByPlaceholderText('password');
		expect(passInput).toBeInTheDocument();
		expect(passInput).toHaveAttribute('type', 'password');
	});

	it('renders the submit button', () => {
		const button = screen.getByRole('button', { name: 'Submit' });
		expect(button).toBeInTheDocument();
	});

	it('allows typing into inputs and clicking submit', () => {
		const emailInput = screen.getByPlaceholderText('your email') as HTMLInputElement;
		const passInput = screen.getByPlaceholderText('password') as HTMLInputElement;
		const button = screen.getByRole('button', { name: 'Submit' });

		fireEvent.change(emailInput, { target: { value: 'test@example.com' } });
		fireEvent.change(passInput, { target: { value: 'secret' } });

		expect(emailInput.value).toBe('test@example.com');
		expect(passInput.value).toBe('secret');

		fireEvent.click(button);
	});
});
