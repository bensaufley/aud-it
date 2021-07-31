// Generated file. Do not edit!

import * as Urql from '@urql/preact';
import gql from 'graphql-tag';

import type * as SchemaTypes from '~graphql/schema.generated';

export type Omit<T, K extends keyof T> = Pick<T, Exclude<keyof T, K>>;
export type GetTodosQueryVariables = SchemaTypes.Exact<{ [key: string]: never }>;

export type GetTodosQuery = {
  __typename?: 'Query';
  todos: Array<{
    __typename?: 'Todo';
    id: string;
    contents: string;
    status: SchemaTypes.TodoStatus;
  }>;
};

export type TodoAddedSubscriptionVariables = SchemaTypes.Exact<{ [key: string]: never }>;

export type TodoAddedSubscription = {
  __typename?: 'Subscription';
  todo: { __typename?: 'Todo'; id: string; contents: string; status: SchemaTypes.TodoStatus };
};

export type TodoDeletedSubscriptionVariables = SchemaTypes.Exact<{ [key: string]: never }>;

export type TodoDeletedSubscription = { __typename?: 'Subscription'; todoID: string };

export type AddTodoMutationVariables = SchemaTypes.Exact<{
  contents: SchemaTypes.Scalars['String'];
}>;

export type AddTodoMutation = {
  __typename?: 'Mutation';
  addTodo: SchemaTypes.Maybe<{ __typename?: 'Todo'; id: string }>;
};

export type DeleteTodoMutationVariables = SchemaTypes.Exact<{
  id: SchemaTypes.Scalars['ID'];
}>;

export type DeleteTodoMutation = { __typename?: 'Mutation'; deleteTodo: boolean };

export const GetTodosDocument = gql`
  query GetTodos {
    todos: getTodos {
      id
      contents
      status
    }
  }
`;

export function useGetTodosQuery(
  options: Omit<Urql.UseQueryArgs<GetTodosQueryVariables>, 'query'> = {},
) {
  return Urql.useQuery<GetTodosQuery>({ query: GetTodosDocument, ...options });
}
export const TodoAddedDocument = gql`
  subscription TodoAdded {
    todo: todoAdded {
      id
      contents
      status
    }
  }
`;

export function useTodoAddedSubscription<TData = TodoAddedSubscription>(
  options: Omit<Urql.UseSubscriptionArgs<TodoAddedSubscriptionVariables>, 'query'> = {},
  handler?: Urql.SubscriptionHandler<TodoAddedSubscription, TData>,
) {
  return Urql.useSubscription<TodoAddedSubscription, TData, TodoAddedSubscriptionVariables>(
    { query: TodoAddedDocument, ...options },
    handler,
  );
}
export const TodoDeletedDocument = gql`
  subscription TodoDeleted {
    todoID: todoDeleted
  }
`;

export function useTodoDeletedSubscription<TData = TodoDeletedSubscription>(
  options: Omit<Urql.UseSubscriptionArgs<TodoDeletedSubscriptionVariables>, 'query'> = {},
  handler?: Urql.SubscriptionHandler<TodoDeletedSubscription, TData>,
) {
  return Urql.useSubscription<TodoDeletedSubscription, TData, TodoDeletedSubscriptionVariables>(
    { query: TodoDeletedDocument, ...options },
    handler,
  );
}
export const AddTodoDocument = gql`
  mutation AddTodo($contents: String!) {
    addTodo(input: { contents: $contents }) {
      id
    }
  }
`;

export function useAddTodoMutation() {
  return Urql.useMutation<AddTodoMutation, AddTodoMutationVariables>(AddTodoDocument);
}
export const DeleteTodoDocument = gql`
  mutation DeleteTodo($id: ID!) {
    deleteTodo(id: $id)
  }
`;

export function useDeleteTodoMutation() {
  return Urql.useMutation<DeleteTodoMutation, DeleteTodoMutationVariables>(DeleteTodoDocument);
}
