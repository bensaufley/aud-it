/* eslint-disable import/no-duplicates */

declare namespace NodeJS {
  export interface ProcessEnv {
    NODE_ENV: 'test' | 'development' | 'production';
  }
}

declare module '*.modules.css' {
  const cssModules: { [key: string]: string };
  export default cssModules;
}

declare module '*.css' {}
