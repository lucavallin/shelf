import React from 'react';
import { Col, Button, Form, FormGroup, Input, Alert } from 'reactstrap';
import { postHerd } from '../services/herd';

export default class HerdUpload extends React.Component {
  constructor(props) {
    super(props);
    this.initialState = { xml: '', error: null };
    this.state = this.initialState;

    this.handleChange = this.handleChange.bind(this);
    this.renderMessage = this.renderMessage.bind(this);
    this.uploadXml = this.uploadXml.bind(this);
  }

  handleChange(e) {
    const xml = e.target.value;
    this.setState(state => ({ ...state, ...{ xml } }));
  }

  renderMessage() {
    const { state } = this;
    if (state.error === null) return null;

    return (
      <Alert color={state.error ? 'danger' : 'success'}>
        {state.error
          ? 'An error has occurred while uploading the new herd.'
          : 'The herd XML has been uploaded successfully.'}
      </Alert>
    );
  }

  // Using redux, this could have been an action being dispatched
  uploadXml(e) {
    e.preventDefault();
    if (this.state.xml.length === 0) return;

    postHerd(this.state.xml)
      .then(() => {
        this.setState(state => ({ ...state, ...{ error: false } }));
      })
      .catch(() => {
        this.setState(state => ({ ...state, ...{ error: true } }));
      })
      .finally(() =>
        // Clear everything
        setTimeout(() => this.setState(this.initialState), 3000),
      );
  }

  render() {
    const { state } = this;

    return (
      <Col md={12}>
        <Col>
          <h2>Herd upload</h2>
        </Col>
        <Col>
          {this.renderMessage()}

          <Form>
            <FormGroup>
              <Input
                type="textarea"
                name="xml"
                id="upload-xml"
                placeholder="Insert your herd xml here..."
                value={state.xml}
                onChange={this.handleChange}
              />
            </FormGroup>
            <Button onClick={this.uploadXml} disabled={state.xml.length === 0}>
              Submit
            </Button>
          </Form>
        </Col>
      </Col>
    );
  }
}
