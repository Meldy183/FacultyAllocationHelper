/** @type {import('jest').Config} */
const config = {
  preset: 'ts-jest',
  testEnvironment: 'jsdom',

  globals: {
    'ts-jest': {
      useESM: false
    }
  },

  moduleFileExtensions: [
    'js', 'jsx', 'ts', 'tsx', 'json', 'node'
  ],

  moduleNameMapper: {
    '^.+\\.(css|scss|sass)$': 'identity-obj-proxy',
    '^.+\\.(png|jpg|jpeg|gif|svg)$': '<rootDir>/__mocks__/fileMock.js',
    '^@/(.*)$': '<rootDir>/$1',
    '^next/image$': '<rootDir>/__mocks__/next/image.js',
    '^next/link$': '<rootDir>/__mocks__/next/link.js',
    '^next/navigation$': '<rootDir>/__mocks__/next/navigation.js',

  },
  // setupFilesAfterEnv: ['@testing-library/jest-dom/extend-expect'],
  transform: {
    '^.+\\.[tj]sx?$': 'ts-jest',
  },

  testMatch: [
    '**/__tests__/**/*.?([mc])[jt]s?(x)',
    '**/?(*.)+(spec|test).[tj]s?(x)'
  ],

  transformIgnorePatterns: [
    '/node_modules/'
  ],

  coverageProvider: 'v8',
};

module.exports = config;
