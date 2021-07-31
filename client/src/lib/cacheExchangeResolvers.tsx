import { UpdateResolver } from '@urql/exchange-graphcache';

import {
  GetTodosDocument,
  GetTodosQuery,
  TodoAddedSubscription,
  TodoDeletedSubscription,
} from '~components/Home/home.generated';

export const todoAdded: UpdateResolver<TodoAddedSubscription> = (result, _args, cache) => {
  if (!result) return;
  if (!result.todo) return;
  cache.updateQuery<GetTodosQuery>({ query: GetTodosDocument }, (data) => ({
    ...data,
    todos: [...(data?.todos || []), result.todo],
  }));
};
export const todoDeleted: UpdateResolver<TodoDeletedSubscription> = (result, _args, cache) => {
  if (!result) return;
  if (!result.todoID) return;
  cache.updateQuery<GetTodosQuery>({ query: GetTodosDocument }, (data) => ({
    ...data,
    todos: data?.todos?.filter(({ id }) => id !== result.todoID) || [],
  }));
};
