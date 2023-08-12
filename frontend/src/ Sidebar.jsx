import React, { useEffect, useState } from 'react';
import { NavLink } from 'react-router-dom';
import logo from '../../assets/img/logo.png'; 
import { Version } from '../wailsjs/go/app/app';

function Sidebar() {
  const [appVersion, setVersion] = useState('');


  useEffect(()=>{
    Version().then((result)=> {
      console.log(result)
      setVersion(result)
    })
  });


  return (
    <div className="sidebar">
      <div className="sidebar-logo">
        <img src={logo} alt="Logo" className="logo-image" />
      </div>
        <NavLink to="/" className="sidebar-link" activeClassName="active" exact='true'>
          <span>On-Flight Generation</span>
        </NavLink>
        <NavLink to="/passwords" className="sidebar-link" activeClassName="active">
          <span>Passwords</span>
        </NavLink>
      <div className="version">Version {appVersion}</div>
    </div>
  );
}

export default Sidebar;
