// @flow
import * as React from 'react';
import { connect } from 'react-redux';
import compose from 'recompose/compose';
import resetJoinTokenMutation from 'shared/graphql/mutations/channel/resetChannelJoinToken';
import { addToastWithTimeout } from 'src/actions/toasts';
import { OutlineButton } from 'src/components/button';
import type { Dispatch } from 'redux';

type Props = {
  id: string,
  settings: {
    tokenJoinEnabled: boolean,
  },
  resetChannelJoinToken: Function,
  dispatch: Dispatch<Object>,
};

type State = {
  isLoading: boolean,
};

class ResetJoinToken extends React.Component<Props, State> {
  state = { isLoading: false };

  reset = () => {
    this.setState({ isLoading: true });
    return this.props
      .resetChannelJoinToken({ id: this.props.id })
      .then(() => {
        this.setState({
          isLoading: false,
        });
        return this.props.dispatch(
          addToastWithTimeout('success', 'Link reset!')
        );
      })
      .catch(err => {
        this.setState({
          isLoading: false,
        });
        return this.props.dispatch(addToastWithTimeout('error', err.message));
      });
  };

  render() {
    const { isLoading } = this.state;

    return (
      <div
        style={{
          display: 'flex',
          justifyContent: 'flex-end',
          marginTop: '16px',
        }}
      >
        <OutlineButton
          loading={isLoading}
          onClick={this.reset}
          data-cy="refresh-join-link-token"
        >
          {isLoading ? 'Resetting' : 'Reset this link'}
        </OutlineButton>
      </div>
    );
  }
}

export default compose(
  connect(),
  resetJoinTokenMutation
)(ResetJoinToken);
