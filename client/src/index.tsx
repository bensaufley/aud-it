/* eslint-disable import/first */
/* istanbul ignore if */
if (process.env.NODE_ENV === 'development') {
  /* eslint-disable global-require */
  require('preact/debug');
  require('preact/devtools');
  /* eslint-enable global-require */
}

import 'core-js/stable';
import 'whatwg-fetch';

import { h, render } from 'preact';

import App from '~components/App';

render(<App />, document.getElementById('root')!);

/* istanbul ignore if */
if (module.hot) {
  module.hot.accept();
}
