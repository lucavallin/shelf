import React from 'react';
import { Container, Row } from 'reactstrap';
import Root from './Root';

// adding prop types validation for children would be better
export default ({ children }) => (
  <Root>
    <Container style={{ padding: '30px 0' }}>
      <Row>{children}</Row>
    </Container>
  </Root>
);
