import { render, screen } from '@testing-library/react';
import { Button } from '@/shared/ui/button';
import '@testing-library/jest-dom';

test('renders Button component', () => {
	render(<Button onClick={() => {}}>Click Me</Button>);
	const buttonElement = screen.getByText('Click Me');
	expect(buttonElement).toBeInTheDocument();
});