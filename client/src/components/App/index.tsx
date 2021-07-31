import '~components/App/index.css';

import type urqlDevtools from '@urql/devtools';
import { cacheExchange } from '@urql/exchange-graphcache';
import {
  Client,
  createClient,
  dedupExchange,
  fetchExchange,
  Provider as UrqlProvider,
  subscriptionExchange,
} from '@urql/preact';
import { FunctionComponent, h } from 'preact';
import { SubscriptionClient } from 'subscriptions-transport-ws';

import styles from '~components/App/styles.modules.css';
import Home from '~components/Home';
import { todoAdded, todoDeleted } from '~lib/cacheExchangeResolvers';

let exchanges = [
  dedupExchange,
  cacheExchange({
    updates: {
      Subscription: {
        todoAdded,
        todoDeleted,
      },
    },
  }),
  fetchExchange,
];
/* istanbul ignore if */
if (process.env.NODE_ENV === 'development') {
  /* eslint-disable global-require, @typescript-eslint/no-var-requires */
  const { devtoolsExchange }: typeof urqlDevtools = require('@urql/devtools');
  exchanges = [devtoolsExchange, ...exchanges];
  /* eslint-enable global-require, @typescript-eslint/no-var-requires */
}

const App: FunctionComponent<{ client?: Client; subscriptionClient?: SubscriptionClient }> = ({
  subscriptionClient = new SubscriptionClient(
    `ws://${window.location.hostname}${
      window.location.port ? `:${window.location.port}` : ''
    }/graphql`,
    { reconnect: true, timeout: 10_000 },
  ),
  client = createClient({
    url: '/graphql',
    exchanges: [
      ...exchanges,
      subscriptionExchange({
        forwardSubscription(operation) {
          return subscriptionClient.request(operation);
        },
      }),
    ],
  }),
}) => (
  <UrqlProvider value={client}>
    <header class={styles.siteHeader}>
      <h1>Todo App</h1>
    </header>
    <nav />
    <main>
      <Home />
    </main>
  </UrqlProvider>
);

export default App;
