// Этот файл конфигурации Jest в формате CommonJS
// Переименуйте текущий jest.config.js в jest.config.cjs

/** @type {import('jest').Config} */
const config = {
  preset: 'ts-jest',
  testEnvironment: 'jsdom',

  // Сборка TS-файлов через ts-jest (CommonJS)
  globals: {
    'ts-jest': {
      useESM: false
    }
  },

  moduleFileExtensions: [
    'js', 'jsx', 'ts', 'tsx', 'json', 'node'
  ],

  moduleNameMapper: {
    // Стили (CSS/SCSS модули)
    '^.+\\.(css|scss|sass)$': 'identity-obj-proxy',
    // Изображения и статика
    '^.+\\.(png|jpg|jpeg|gif|svg)$': '<rootDir>/__mocks__/fileMock.js',
    // Алиасы
    '^@/(.*)$': '<rootDir>/$1',
    // Моки Next.js компонентов
    '^next/image$': '<rootDir>/__mocks__/next/image.js',
    '^next/link$': '<rootDir>/__mocks__/next/link.js',
    '^next/navigation$': '<rootDir>/__mocks__/next/navigation.js'
  },

  transform: {
    '^.+\\.[tj]sx?$': 'ts-jest'
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
