import React from 'react';
import Head from '../components/Head';
import Footer from '../components/Footer';
import Navbar from '../components/Navbar';
import 'bootstrap/dist/css/bootstrap.min.css';

// adding prop types validation for children would be better
export default ({ children }) => (
  <div>
    <Head />
    <Navbar />
    {children}
    <Footer />
  </div>
);
