import { FunctionComponent, h, JSX } from 'preact';
import { useCallback, useState } from 'preact/hooks';

import {
  useAddTodoMutation,
  useDeleteTodoMutation,
  useGetTodosQuery,
  useTodoAddedSubscription,
  useTodoDeletedSubscription,
} from '~components/Home/home.generated';
import { Todo } from '~graphql/schema.generated';

const ListItem = ({ id, contents, status }: Todo) => {
  const [_, deleteMutation] = useDeleteTodoMutation();
  const deleteTodo: JSX.MouseEventHandler<HTMLButtonElement> = useCallback(
    (e) => {
      e.preventDefault();
      deleteMutation({ id: id.toString() });
    },
    [deleteMutation, id],
  );

  return (
    <li id={id} title={id}>
      {contents} ({status}){' '}
      <button type="button" onClick={deleteTodo}>
        &times;
      </button>
    </li>
  );
};

const Home: FunctionComponent = () => {
  const [{ data, error, fetching }] = useGetTodosQuery();
  const [_, addTodo] = useAddTodoMutation();
  useTodoAddedSubscription();
  useTodoDeletedSubscription();

  const [newTodoContent, setNewTodoContent] = useState('');
  const handleChange: JSX.GenericEventHandler<HTMLInputElement> = useCallback(
    ({ currentTarget }) => {
      setNewTodoContent(currentTarget.value);
    },
    [setNewTodoContent],
  );

  const createTodo = useCallback(
    (e: Event) => {
      e.preventDefault();
      addTodo({ contents: newTodoContent });
      setNewTodoContent('');
    },
    [addTodo, newTodoContent],
  );

  return (
    <div>
      <h2>Home</h2>
      <form onSubmit={createTodo}>
        <input type="text" onChange={handleChange} name="contents" value={newTodoContent} />
        <button type="submit">Create</button>
      </form>
      {fetching && <p>Loading&hellip;</p>}
      {!fetching && error && <p>Error: {error.message}</p>}
      {!fetching && !error && (
        <ul>
          {data?.todos?.map(({ id, contents, status }) => (
            <ListItem id={id} contents={contents} status={status} />
          ))}
        </ul>
      )}
    </div>
  );
};

export default Home;
