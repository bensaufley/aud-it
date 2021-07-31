// Generated file. Do not edit!

export type Maybe<T> = T | null;
export type Exact<T extends { [key: string]: unknown }> = { [K in keyof T]: T[K] };
export type MakeOptional<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]?: Maybe<T[SubKey]> };
export type MakeMaybe<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]: Maybe<T[SubKey]> };
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: string;
  String: string;
  Boolean: boolean;
  Int: number;
  Float: number;
};

export type Mutation = {
  __typename?: 'Mutation';
  addTodo: Maybe<Todo>;
  deleteTodo: Scalars['Boolean'];
};

export type MutationAddTodoArgs = {
  input: TodoInput;
};

export type MutationDeleteTodoArgs = {
  id: Scalars['ID'];
};

export type Query = {
  __typename?: 'Query';
  getTodos: Array<Todo>;
  getTodo: Maybe<Todo>;
};

export type QueryGetTodoArgs = {
  id: Scalars['ID'];
};

export type Subscription = {
  __typename?: 'Subscription';
  todoAdded: Todo;
  todoDeleted: Scalars['ID'];
};

export type Todo = {
  __typename?: 'Todo';
  id: Scalars['ID'];
  contents: Scalars['String'];
  status: TodoStatus;
};

export type TodoInput = {
  contents: Scalars['String'];
  status: Maybe<TodoStatus>;
};

export enum TodoStatus {
  Unstarted = 'UNSTARTED',
  InProgress = 'IN_PROGRESS',
  Abandoned = 'ABANDONED',
  Complete = 'COMPLETE',
  Deleted = 'DELETED',
}
