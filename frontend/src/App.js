import React, { useState } from "react";
import axios from "axios";

function App() {
  const [file, setFile] = useState(null);
  const [fileQuery, setFileQuery] = useState(""); 
  const [chatQuery, setChatQuery] = useState("");
  const [response, setResponse] = useState("");

  const handleFileChange = (e) => {
    setFile(e.target.files[0]);
  };

  const handleUpload = async () => {
    const formData = new FormData();
    formData.append("file", file);
    formData.append("query", fileQuery);

    try {
      const res = await axios.post("http://localhost:8080/upload", formData, {
        headers: {
          "Content-Type": "multipart/form-data",
        },
      });
      setResponse(res.data.answer);  
    } catch (error) {
      console.error("Error uploading file:", error);
    }
  };

  const handleChat = async () => {
    try {
      const res = await axios.post("http://localhost:8080/chat", { query: chatQuery });
      setResponse(res.data.answer);
    } catch (error) {
      console.error("Error querying chat:", error);
    }
  };

  return (
    <div
      style={{
        maxWidth: "600px",
        margin: "0 auto",
        padding: "20px",
        textAlign: "center",
        fontFamily: "Arial, sans-serif",
        backgroundColor: "#1c1c1c",
        color: "#fff", 
        borderRadius: "8px",
      }}
    >
      <h1 style={{ color: "#0ff", marginBottom: "20px", textShadow: "0 0 10px #0ff" }}>
        Data Analysis Chatbot
      </h1>
      <div style={{ marginBottom: "20px" }}>
        <input
          type="file"
          onChange={handleFileChange}
          style={{
            padding: "10px",
            marginRight: "10px",
            border: "1px solid #0ff",
            borderRadius: "4px",
            backgroundColor: "#333",
            color: "#fff",
          }}
        />
        <input
          type="text"
          value={fileQuery}
          onChange={(e) => setFileQuery(e.target.value)}
          placeholder="Custom query for file analysis..."
          style={{
            padding: "10px",
            marginRight: "10px",
            border: "1px solid #0ff",
            borderRadius: "4px",
            width: "calc(100% - 170px)",
            backgroundColor: "#333",
            color: "#fff",
          }}
        />
        <button
          onClick={handleUpload}
          style={{
            padding: "10px 20px",
            backgroundColor: "#0ff",
            color: "black",
            border: "none",
            borderRadius: "4px",
            cursor: "pointer",
            boxShadow: "0 0 10px #0ff",
          }}
        >
          Upload and Analyze
        </button>
      </div>
      <div style={{ marginBottom: "20px" }}>
        <input
          type="text"
          value={chatQuery}
          onChange={(e) => setChatQuery(e.target.value)}
          placeholder="Ask a question..."
          style={{
            padding: "10px",
            marginRight: "10px",
            border: "1px solid #0ff",
            borderRadius: "4px",
            width: "calc(100% - 140px)",
            backgroundColor: "#333",
            color: "#fff",
          }}
        />
        <button
          onClick={handleChat}
          style={{
            padding: "10px 20px",
            backgroundColor: "#0ff",
            color: "black",
            border: "none",
            borderRadius: "4px",
            cursor: "pointer",
            boxShadow: "0 0 10px #0ff",
          }}
        >
          Chat
        </button>
      </div>
      <div
        style={{
          marginTop: "20px",
          padding: "10px",
          border: "1px solid #0ff",
          borderRadius: "4px",
          backgroundColor: "#333",
          color: "#fff",
        }}
      >
        <h2>Response</h2>
        <p>{response}</p>
      </div>
    </div>
  );
}

export default App;
