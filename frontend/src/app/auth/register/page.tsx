"use client";
import { useState } from "react";
import axios from "axios";
import SignupFormDemo from "@/components/signup-form-demo";

export default function Register() {
  const [form, setForm] = useState({ username: "", email: "", password: "" });

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    try {
      const res = await axios.post("http://localhost:8080/register", form);
      alert(res.data.message);
    } catch (error: any) {
      alert(error.response.data.error);
    }
  };

  return (
    <div className="flex justify-start items-start h-screen bg-gray-100 px-20">
      <div className="max-w-md w-full"> 
        <SignupFormDemo />
      </div>
    </div>
  );
}
