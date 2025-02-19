"use client"
import { useState } from "react";
import axios from "axios";

export default function Login() {
  const [form, setForm] = useState({ email: "", password: "" });

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    try {
      const res = await axios.post("http://localhost:8080/login", form);
      alert("Login Successful!");
      localStorage.setItem("token", res.data.token);
    } catch (error: any) {
      alert(error.response.data.error);
    }
  };

  return (
    <div className="flex justify-center items-center h-screen bg-gray-100">
      <form className="bg-white p-6 rounded-lg shadow-md w-96" onSubmit={handleSubmit}>
        <h2 className="text-xl font-bold mb-4">Login</h2>
        <input type="email" placeholder="Email" className="border p-2 w-full mb-3"
          onChange={(e) => setForm({ ...form, email: e.target.value })} />
        <input type="password" placeholder="Password" className="border p-2 w-full mb-3"
          onChange={(e) => setForm({ ...form, password: e.target.value })} />
        <button type="submit" className="bg-green-500 text-white p-2 rounded w-full">Login</button>
      </form>
    </div>
  );
}

