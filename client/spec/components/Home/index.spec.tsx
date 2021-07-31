import { render } from '@testing-library/preact';
import { h } from 'preact';

import Home from '~components/Home';

describe('~components/Home', () => {
  it('works', () => {
    const { getByText } = render(<Home />);

    getByText('Home');
  });
});
