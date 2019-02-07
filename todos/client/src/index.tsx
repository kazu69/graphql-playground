import React from 'react';
import { render } from 'react-dom';
import ApolloClient from 'apollo-boost';
import { ApolloProvider } from 'react-apollo';

const gql = 'http://localhost:8080';
const client = new ApolloClient({
  uri: gql
});

const App = () => (
  <ApolloProvider client={client}>
    <div>
      <h2>My first Apollo app</h2>
    </div>
  </ApolloProvider>
);

render(<App />, document.getElementById('root'));
