import React from "react";
import { Link } from "react-router-dom";
import EnergiQLogo from "./EnergiQ.png";

function Home() {
  return (
    <div style={{ textAlign: "center", padding: "50px", fontFamily: "Arial, sans-serif" }}>
      <img 
        src={EnergiQLogo} 
        alt="EnergiQ" 
        style={{ width: "100%", height: "100%" }}
      />
      <h1>Welcome to the Chatbot App</h1>
      <p>Please login or register to proceed.</p>
      <div style={{ marginTop: "20px" }}>
        <Link to="/login" style={buttonStyle}>
          Login
        </Link>
        <Link to="/register" style={buttonStyle}>
          Register
        </Link>
      </div>
    </div>
  );
}

const buttonStyle = {
  display: "inline-block",
  margin: "0 10px",
  padding: "10px 20px",
  textDecoration: "none",
  color: "#fff",
  backgroundColor: "#007BFF",
  borderRadius: "5px",
  boxShadow: "0 3px 5px rgba(0, 0, 0, 0.2)",
};

export default Home;
