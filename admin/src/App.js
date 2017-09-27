import React, { Component } from 'react';
import {Row} from 'react-bootstrap';
import {NavLink, Route} from 'react-router-dom';
import FontAwesome from 'react-fontawesome';
import Home from './pages/Home';
import Models from './pages/Models';
import Features from './pages/Features';
import logo from './logo.svg';
import './App.css';

class App extends Component {
  render() {
    return (
      <Row>
        <div className="col-sm-3 col-md-2 sidebar">
          <h1>Photon</h1>
          <ul className="nav nav-sidebar">
            <li><NavLink exact to="/"><FontAwesome name="home"/> Home</NavLink></li>
            <li><NavLink to="/models"><FontAwesome name="th-large"/> Models</NavLink></li>
            <li><NavLink to="/features"><FontAwesome name="database"/> Features</NavLink></li>
            <li><NavLink to="/cluster"><FontAwesome name="sitemap"/> Cluster</NavLink></li>
            <li><NavLink to="/settings"><FontAwesome name="cog"/> Settings</NavLink></li>
          </ul>
        </div>
        <div className="col-sm-9 col-sm-offset-3 col-md-10 col-md-offset-2 top-nav">
          <FontAwesome name="bars" className="sidebar-toggle"/>
          <ul className="pull-right nav">
            <li><span>Marek Galovic</span></li>
            <li><FontAwesome name="sign-out"/></li>
          </ul>
        </div>
        <div className="col-sm-9 col-sm-offset-3 col-md-10 col-md-offset-2 main">
          <Row>
            <Route exact path="/" component={Home}/>
            <Route path="/models" component={Models}/>
            <Route path="/features" component={Features}/>
          </Row>
        </div>
      </Row>
    );
  }
}

export default App;
