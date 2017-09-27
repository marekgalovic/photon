import React, {Component} from 'react';
import {Row, Table, Button} from 'react-bootstrap';

class Models extends Component {
    render() {
        return (
            <div className="container">
                <Row>
                    <div className="col content-box-wrapper col-xs-12">
                        <h2 className="page-title">Models</h2>
                    </div>
                    <div className="col content-box-wrapper col-xs-12">
                        <div className="content-box">
                            <div className="content-box-body">
                                <Table responsive>
                                    <thead>
                                        <tr>
                                            <th>Model name</th>
                                            <th>Owner</th>
                                            <th>Versions</th>
                                            <th className="text-right actions">
                                                <Button bsStyle="success" bsSize="xsmall">New Model</Button>
                                            </th>
                                        </tr>
                                    </thead>
                                    <tbody>
                                        <tr>    
                                            <td>Iris</td>
                                            <td>foo@bar.com</td>
                                            <td>5</td>
                                            <td className="text-right actions">
                                                <Button bsStyle="primary" bsSize="xsmall">Show</Button>
                                                <Button bsStyle="danger" bsSize="xsmall">Delete</Button>
                                            </td>
                                        </tr>
                                    </tbody>
                                </Table>
                            </div>
                        </div>
                    </div>
                </Row>
            </div>
        );
    }
}

export default Models;
