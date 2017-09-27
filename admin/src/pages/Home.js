import React, {Component} from 'react';
import {Row} from 'react-bootstrap';

class Home extends Component {
    render() {
        return (
            <div className="container">
                <Row>
                    <div className="col content-box-wrapper col-xs-12">
                        <h2 className="page-title">Home</h2>
                    </div>
                    <div className="col content-box-wrapper col-md-6">
                        <div className="content-box">
                            <h3 className="content-box-title">Home</h3>
                            <div className="content-box-body">
                                This is a body!
                            </div>
                        </div>
                    </div>
                    <div className="col content-box-wrapper col-md-6">
                        <div className="content-box">
                            <h3 className="content-box-title">Home</h3>
                            <div className="content-box-body">
                                This is a body!
                            </div>
                        </div>
                    </div>
                    <div className="col content-box-wrapper col-md-12">
                        <div className="content-box">
                            <h3 className="content-box-title">Home</h3>
                            <div className="content-box-body">
                                This is a body!
                            </div>
                        </div>
                    </div>
                </Row>
            </div>
        );
    }
}

export default Home;
