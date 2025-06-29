import React from 'react';
export default function Link({ children, href }) {
    return <a href={href}>{children}</a>;
}