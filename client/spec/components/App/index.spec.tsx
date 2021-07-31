import { render } from '@testing-library/preact';
import { h } from 'preact';

import App from '~components/App';

describe('~components/App', () => {
  it('works', () => {
    const { getByText } = render(<App />);

    getByText('Todo App');
  });
});
